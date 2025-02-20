package service

import (
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"fmt"
	"github.com/boltdb/bolt"
	"log"
	"math/big"
	"os"
	"time"
)

const dbName = "blockchain.db" // 数据库名称
const bkName = "blocks"        // 桶名称
const targetBit = 2

func IsValidArgs() {
	if len(os.Args) < 2 {
		PrintUsage() // 打印用法
		os.Exit(1)   // 退出程序
	}
}

// 展示用法
func PrintUsage() {
	fmt.Println("Usage:")
	fmt.Printf("\tcreateblockchain -- 创建区块链.\n")
	fmt.Printf("\taddblock -data DATA -- 交易数据\n")
	fmt.Printf("\tprintchain -- 输出区块链的信息\n")
}

func DeserializeBlock(blockBytes []byte) *Block {
	var block Block
	decoder := gob.NewDecoder(bytes.NewReader(blockBytes))
	if err := decoder.Decode(&block); err != nil {
		log.Panicf("deserialize the block to byte failed %v \n", err)
	}
	return &block
}

func dbExists() bool {
	if _, err := os.Stat(dbName); os.IsNotExist(err) {
		return false
	}
	return true
}

func Blockchain_GenesisBlock() *Blockchain {
	if dbExists() {
		fmt.Println("区块已经存在。。。")
		db, err := bolt.Open(dbName, 0600, nil)
		if err != nil {
			log.Panicf("open the Dbfailed! %v\n", err)
		}
		var blockchain *Blockchain
		err = db.View(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte(bkName))
			tip := b.Get([]byte("l")) // 这里使用的是 read-only事务的 Get 方法，从l中读取最后一块区块的编码，我们挖下一新块时会作为参数用到。
			//block := DeserializeBlock(b.Get(tip))
			blockchain = &Blockchain{
				Tip: tip,
				Db:  db,
			}
			return nil
		})
		if err != nil {
			log.Panicf("get the block from db failed! %v\n", err)
		}
		return blockchain
	} else {
		db, err := bolt.Open(dbName, 0600, nil)
		if err != nil {
			log.Panicf("open the Db failed! %v\n", err)
		}

		var (
			tip          []byte // 存储链中最后一个区块的哈希值
			genesisBlock *Block
		)
		err = db.Update(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte(bkName))
			if b == nil {
				b, err = tx.CreateBucket([]byte(bkName))
				if err != nil {
					log.Panicf("create the bucket [%s] failed! %v\n", bkName, err)
				}
			}
			genesisBlock = NewGenesisBlock()
			// 存储创世区块
			err = b.Put(genesisBlock.Hash, genesisBlock.Serialize())
			if err != nil {
				log.Panicf("put the data of genesisBlock to Db failed! %v\n", err)
			}
			// 存储最新区块链哈希
			err = b.Put([]byte("l"), genesisBlock.Hash) // 在挖出新块，将其序列化存储到数据库后，把最新的区块hash值更新到 l 值中
			if err != nil {
				log.Panicf("put the hash of latest block to Db failed! %v\n", err)
			}
			tip = genesisBlock.Hash
			return nil
		})
		if err != nil {
			log.Panicf("update the data of genesis block failed! %v\n", err)
		}
		return &Blockchain{
			Blocks: []*Block{
				genesisBlock,
			},
			Tip: tip,
			Db:  db,
		}
	}
}

func NewGenesisBlock() *Block {
	return NewBlock(0, "first block", []byte{})
}

func NewBlock(index int64, data string, prevBlockHash []byte) *Block {
	block := &Block{
		Index:         index,
		Timestamp:     time.Now().Unix(),
		Data:          []byte(data),
		PrevBlockHash: prevBlockHash,
		Hash:          []byte{},
		Nonce:         0,
	}
	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()
	block.Hash = hash[:]
	block.Nonce = nonce
	return block
}

func NewProofOfWork(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-targetBit))
	pow := &ProofOfWork{
		block:  b,
		target: target,
	}
	return pow
}

func IntToHex(data int64) []byte {
	buffer := new(bytes.Buffer) // 新建一个buffer
	err := binary.Write(buffer, binary.BigEndian, data)
	if nil != err {
		log.Panicf("int to []byte failed! %v\n", err)
	}
	return buffer.Bytes()
}
