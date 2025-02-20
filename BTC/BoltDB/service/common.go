package service

import (
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"github.com/boltdb/bolt"
	"log"
	"math/big"
	"time"
)

const dbName = "blockchain.db" // 数据库名称
const bkName = "blocks"        // 桶名称
const targetBit = 2

func DeserializeBlock(blockBytes []byte) *Block {
	var block Block
	decoder := gob.NewDecoder(bytes.NewReader(blockBytes))
	if err := decoder.Decode(&block); err != nil {
		log.Panicf("deserialize the block to byte failed %v \n", err)
	}
	return &block
}

func Blockchain_GenesisBlock() *Blockchain {
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
