package main

import (
	"fmt"
	"publicChain/part4-Basic-Prototype/BLC"
)

func main() {
	createBlockchain := BLC.CreateBlockchainWithGenesisBlock()
	fmt.Println(createBlockchain)
}
