package service

import (
	"fmt"
	"github.com/boltdb/bolt"
	"log"
	"math/big"
)

type Blockchain struct {
	Blocks []*Block
	Tip    []byte   // 存储最新区块的哈希值
	Db     *bolt.DB // 数据库实例
}

// 创建迭代器对象
func (bc *Blockchain) Iterator() *BlockchainIterator {
	return &BlockchainIterator{
		DB:          bc.Db,
		CurrentHash: bc.Tip,
	}
}

func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := NewBlock(prevBlock.Index+1, data, prevBlock.Hash)

	var (
		tip []byte // 存储链中最后一个区块的哈希值
		err error
	)
	err = bc.Db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bkName))
		if b == nil {
			b, err = tx.CreateBucket([]byte(bkName))
			if err != nil {
				log.Panicf("create the bucket [%s] failed! %v\n", bkName, err)
			}
		}
		err = b.Put(newBlock.Hash, newBlock.Serialize())
		if err != nil {
			log.Panicf("put the data of genesisBlock to Db failed! %v\n", err)
		}
		err = b.Put([]byte("l"), newBlock.Hash) // 存储最新块的哈希
		if err != nil {
			log.Panicf("put the hash of latest block to Db failed! %v\n", err)
		}
		tip = newBlock.Hash
		return nil
	})
	if err != nil {
		log.Panicf("update the data of genesis block failed! %v\n", err)
	}

	bc.Blocks = append(bc.Blocks, newBlock)
	bc.Tip = tip
}

func (bc *Blockchain) PrintChain() {
	fmt.Println("——————————————打印区块链———————————————————————")
	var (
		curBlock *Block
		curHash  = bc.Tip
	)
	for {
		bc.Db.View(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte(bkName))
			if b != nil {
				blockBytes := b.Get(curHash)
				curBlock = DeserializeBlock(blockBytes)

				fmt.Printf("\tHeight : %d\n", curBlock.Index)
				fmt.Printf("\tTimestamp : %d\n", curBlock.Timestamp)
				fmt.Printf("\tPrevBlockHash : %x\n", curBlock.PrevBlockHash)
				fmt.Printf("\tHash : %x\n", curBlock.Hash)
				fmt.Printf("\tData : %s\n", string(curBlock.Data))
				fmt.Printf("\tNonce : %d\n", curBlock.Nonce)
			}
			return nil
		})

		// 判断是否到达创世区块
		var hashInt big.Int
		hashInt.SetBytes(curBlock.PrevBlockHash)
		if big.NewInt(0).Cmp(&hashInt) == 0 {
			break // 跳出循环
		}
		curHash = curBlock.PrevBlockHash
	}
}
