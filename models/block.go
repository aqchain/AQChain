package models

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"time"
)

type Block struct {
	Number     int64  `orm:"pk;column(number)" json:"number"`
	Hash       string `orm:"column(hash)" json:"hash"`
	Creator    string `orm:"column(creator)" json:"creator"`
	PrevHash   string `orm:"column(prev_hash)" json:"prev_hash"`
	Timestamp  string `orm:"column(timestamp)" json:"timestamp"`
	MerkleRoot string `orm:"column(merkle_root)" json:"merkle_root"`

	Transactions `orm:"-"`
}

func NewBlock(prevBlock Block, txs []*Transaction, creator string) Block {
	block := Block{
		Number:       prevBlock.Number + 1,
		Creator:      creator,
		PrevHash:     prevBlock.Hash,
		Timestamp:    time.Now().Format("2006-01-02 15:04:05"),
		Transactions: Transactions{txs},
	}
	// 生成MerkelRoot 暂时没有这个方法

	// 生成区块Hash
	block.Hash, _ = block.EncodeBlock()
	return block
}

func (b *Block) EncodeBlock() (string, error) {
	h := sha256.New()
	if len(b.Creator) == 0 {
		return "", errors.New("EncodeBlock miss Creator")
	}
	h.Write([]byte(b.Creator))
	if len(b.PrevHash) == 0 {
		return "", errors.New("EncodeBlock miss PrevHash")
	}
	h.Write([]byte(b.PrevHash))
	/*if len(b.MerkleRoot) == 0{
		return "",errors.New("EncodeBlock miss MerkleRoot")
	}*/
	h.Write([]byte(b.MerkleRoot))

	sum := h.Sum(nil)
	return hex.EncodeToString(sum), nil
}
