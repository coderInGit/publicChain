package BLC

type Blockchain struct {
	Blocks []*Block
}

func (blc *Blockchain) AddBlockToBlockchain(data string, height int64, preHash []byte) {
	newBlock := NewBlock(data, height, preHash)
	blc.Blocks = append(blc.Blocks, newBlock)
}
func CreateBlockchainWithGenesisBlock() *Blockchain {
	genesisBlock := CreategenesisBlock("Genesis Data。。。。")
	return &Blockchain{[]*Block{genesisBlock}}
}
