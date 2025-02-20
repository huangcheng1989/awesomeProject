package service

import (
	"bytes"
	"encoding/binary"
	"log"
	"math/big"
	"time"
)

const targetBit = 2

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
