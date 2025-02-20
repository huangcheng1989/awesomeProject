package service

import (
	"flag"
	"log"
	"os"
)

type CLI struct {
	BC *Blockchain
}

// 添加区块
func (cli *CLI) addBlock(data string) {
	cli.BC.AddBlock(data)
}

// 输出区块链信息
func (cli *CLI) printChain() {
	cli.BC.PrintChain()
}

// 创建区块链
func (cli *CLI) createBlockchainWithGenesis() {
	CreateBlockChainWithGenesisBlock()
}

func (cli *CLI) Run() {
	// 1. 检测参数数量
	IsValidArgs()
	// 2. 新建命令
	addBlockCmd := flag.NewFlagSet("addblock", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("printchain", flag.ExitOnError)
	createBlCWithGenesisCmd := flag.NewFlagSet("createblockchain", flag.ExitOnError)
	// 3. 获取命令行参数
	flagAddBlockArg := addBlockCmd.String("data", "send 100 BTC to everyone", "交易数据")
	switch os.Args[1] {
	case "addblock":
		err := addBlockCmd.Parse(os.Args[2:])
		if nil != err {
			log.Panicf("parse cmd of add block failed! %v\n", err)
		}
	case "printchain":
		err := printChainCmd.Parse(os.Args[2:])
		if nil != err {
			log.Panicf("parse cmd of printchain failed! %v\n", err)
		}
	case "createblockchain":
		err := createBlCWithGenesisCmd.Parse(os.Args[2:])
		if nil != err {
			log.Panicf("parse cmd of create block chain failed! %v\n", err)
		}
	default:
		PrintUsage()
		os.Exit(1)
	}
	// 添加区块命令
	if addBlockCmd.Parsed() {
		if *flagAddBlockArg == "" {
			PrintUsage()
			os.Exit(1)
		}
		cli.addBlock(*flagAddBlockArg)
	}

	// 输出区块链信息命令
	if printChainCmd.Parsed() {
		cli.printChain()
	}
	// 创建区块链
	if createBlCWithGenesisCmd.Parsed() {
		cli.createBlockchainWithGenesis()
	}
}
