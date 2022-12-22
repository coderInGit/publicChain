package main

import (
	"fmt"
	"publicChain/part4-Basic-Prototype/BLC"
)

func main() {
	blockchain := BLC.CreateBlockchainWithGenesisBlock()
	fmt.Println(blockchain)

	blockchain.AddBlockToBlockchain("Send 100RMb", blockchain.Blocks[len(blockchain.Blocks)-1].Height+1,
		blockchain.Blocks[len(blockchain.Blocks)-1].Hash)
}
