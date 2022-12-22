package BLC

import "math/big"

const targetBit = 16

type ProofOfWork struct {
	Block  *Block
	target *big.Int
}

func (proofOfWork *ProofOfWork) Run() ([]byte, int64) {
	return nil, 0
}
func NewProofOfWork(block *Block) *ProofOfWork {
	return &ProofOfWork{block, 5}
}
