package main

import (
	"fmt"
	"publicChain/part2-Basic-Prototype/BLC"
)

func main() {
	block := BLC.CreategenesisBlock("Genes Block 。。。")
	fmt.Println(block)
}
