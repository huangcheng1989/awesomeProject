package main

import (
	"awesomeProject/BTC/blockchain/service"
	"fmt"
)

func main() {
	bc := service.NewBlockchain()
	bc.AddBlock("Joy send 1 BTC to Jay")
	bc.AddBlock("Jack sent 2 BTC to Jay")

	for _, block := range bc.Blocks {
		fmt.Printf("Index : %d\n", block.Index)
		fmt.Printf("TimeStamp: %d\n", block.TimeStamp)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("PrevHash: %x\n", block.PrevBlockHash)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println("_____________________________")
	}
}
