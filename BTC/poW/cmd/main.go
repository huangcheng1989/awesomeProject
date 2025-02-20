package main

import (
	"awesomeProject/BTC/poW/service"
	"fmt"
)

func main() {
	bc := service.NewBlockchain()
	fmt.Printf("blockChain : %v\n", bc)
	bc.AddBlock("Aimi send 100 BTC	to Bob")
	bc.AddBlock("Aimi send 100 BTC	to Jay")
	bc.AddBlock("Aimi send 100 BTC	to Clown")
	length := len(bc.Blocks)
	fmt.Printf("length of blocks : %d\n", length)
	for i := 0; i < length; i++ {
		pow := service.NewProofOfWork(bc.Blocks[i])
		if pow.Validate() {
			fmt.Println("—————————————————————————————————————————————————————")
			fmt.Printf(" Block: %d\n", bc.Blocks[i].Index)
			fmt.Printf("Data: %s\n", bc.Blocks[i].Data)
			fmt.Printf("TimeStamp: %d\n", bc.Blocks[i].TimeStamp)
			fmt.Printf("Hash: %x\n", bc.Blocks[i].Hash)
			fmt.Printf("PrevHash: %x\n", bc.Blocks[i].PrevBlockHash)
			fmt.Printf("Nonce: %d\n", bc.Blocks[i].Nonce)
		} else {
			fmt.Println("illegal block")
		}
	}
}
