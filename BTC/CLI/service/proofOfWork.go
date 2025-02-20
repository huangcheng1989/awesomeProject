package service

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math/big"
)

type ProofOfWork struct {
	block  *Block
	target *big.Int
}

func (pow *ProofOfWork) Run() (int64, []byte) {
	var (
		hashInt big.Int
		hash    [32]byte
		nonce   int64 = 0
	)
	fmt.Printf("Mining the block containing \"%s\"\n", pow.block.Data)
	for {
		dataBytes := pow.prepareData(nonce) //获取准备的数据
		hash = sha256.Sum256(dataBytes)     //对数据进行Hash
		hashInt.SetBytes(hash[:])
		fmt.Printf("hash: \r%x", hash)
		if pow.target.Cmp(&hashInt) == 1 { //对比hash值
			break
		}
		nonce++ //充当计数器，同时在循环结束后也是符合要求的值
	}
	fmt.Printf("\n碰撞次数: %d\n", nonce)
	return nonce, hash[:]
}

func (pow *ProofOfWork) prepareData(nonce int64) []byte {
	data := bytes.Join(
		[][]byte{
			pow.block.PrevBlockHash,
			pow.block.Data,
			IntToHex(pow.block.Index),
			IntToHex(pow.block.Timestamp),
			IntToHex(int64(targetBit)),
			IntToHex(nonce),
		},
		[]byte{},
	)
	return data
}
