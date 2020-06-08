package core

import (
	"AQChain/db"
	"AQChain/models"
)

type Chain struct {
	blockChain []models.Block
}

func NewChain() Chain {
	var blockChain []models.Block
	lastBlock := db.ReadLastBlock()
	if lastBlock == nil {
		genesisBlock := db.GenesisBlock()
		blockChain = append(blockChain, genesisBlock)
	} else {
		blockChain = db.ReadBlocks(0, lastBlock.Number)
	}

	return Chain{blockChain}
}

func (c Chain) Len() int64 {
	return int64(len(c.blockChain))
}
