package main

import "awesomeProject/BTC/BoltDB/service"

func main() {
	blockChain := service.Blockchain_GenesisBlock()
	defer blockChain.Db.Close()
	//blockChain.AddBlock("Send 100 BTC to Jay")
	//blockChain.AddBlock("Send 50 BTC to Clown")
	//blockChain.AddBlock("Send 20 BTC to Bob")
	//blockChain.PrintChain()
	service.PrintUsage()
	cli := service.CLI{
		BC: blockChain,
	}
	cli.Run()
}
