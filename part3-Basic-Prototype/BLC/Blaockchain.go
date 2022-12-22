package BLC

type Blockchain struct {
	Blocks []*Block
}

func CreateBlockchainWithGenesisBlock() *Blockchain {
	genesisBlock := CreategenesisBlock("Genesis Data。。。。")
	return &Blockchain{[]*Block{genesisBlock}}
}
