package service

import "time"

func NewBlockchain() *Blockchain {
	return &Blockchain{
		Blocks: []*Block{
			NewGenesisBlock(),
		},
	}
}

func NewGenesisBlock() *Block {
	return NewBlock(0, "first block", []byte{})
}

func NewBlock(index int64, data string, prevBlockHash []byte) *Block {
	block := &Block{
		Index:         index,
		TimeStamp:     time.Now().Unix(),
		Data:          []byte(data),
		PrevBlockHash: prevBlockHash,
		Hash:          []byte{},
	}
	block.setHash() // 设置当前区块 Hash
	return block
}
