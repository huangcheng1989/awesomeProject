package service

import (
	"github.com/boltdb/bolt"
	"log"
)

// 区块链迭代器结构
type BlockchainIterator struct {
	DB          *bolt.DB // 数据库
	CurrentHash []byte   // 当前区块的哈希值
}

func (bcit *BlockchainIterator) Next() *Block {
	var block *Block
	err := bcit.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bkName))
		if nil != b {
			// 获取指定哈希的区块数据
			currentBlockBytes := b.Get(bcit.CurrentHash)
			block = DeserializeBlock(currentBlockBytes)
			// 更新迭代器中当前区块的哈希值
			bcit.CurrentHash = block.PrevBlockHash
		}
		return nil
	})
	if nil != err {
		log.Panicf("iterator the db of blockchain failed! %v\n", err)
	}
	return block
}
