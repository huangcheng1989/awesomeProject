package service

import (
	"bytes"
	"encoding/gob"
	"log"
)

type Block struct {
	Index         int64
	Timestamp     int64
	Data          []byte
	PrevBlockHash []byte
	Hash          []byte
	Nonce         int64
}

func (b *Block) Serialize() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)
	if err := encoder.Encode(b); err != nil {
		log.Panicf("serialize the block to byte failed %v \n", err)
	}
	return result.Bytes()
}
