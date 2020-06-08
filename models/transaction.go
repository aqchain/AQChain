package models

import (
	msg "AQChain/models/pb"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"strconv"
	"time"
)

var (
	encodeTransactionMissFromError      = errors.New("EncodeTransaction miss From")
	encodeTransactionMissFileIDError    = errors.New("EncodeTransaction miss FileID")
	encodeTransactionMissTimestampError = errors.New("EncodeTransaction miss Timestamp")
)

//交易记录
type Transaction struct {
	Hash        string  `orm:"pk;column(hash)" json:"hash"`              // 交易hash
	From        string  `orm:"column(from)" json:"from"`                 // 发起者
	To          string  `orm:"column(to)" json:"to"`                     // 接收者
	Value       float64 `orm:"column(value)" json:"value"`               // 交易价格
	FileID      string  `orm:"column(file_id)" json:"file_id"`           // 交易的文件ID
	Timestamp   string  `orm:"column(timestamp)" json:"timestamp"`       // 交易时间
	Type        int32   `orm:"column(type)" json:"type"`                 // 0表示知识产权、1表示交易
	Status      int32   `orm:"column(status)" json:"status"`             // 0表示为确认的交易、1表示已经确认了交易、2表示完成转账操作的交易
	BlockNumber int64   `orm:"column(block_number)" json:"block_number"` // block
}

func NewTransaction(from, to, fileID string, value float64, t, status int32) *Transaction {
	tx := &Transaction{
		From:      from,
		To:        to,
		Value:     value,
		FileID:    fileID,
		Timestamp: time.Now().Format("2006-01-02 15:04:05"),
		Type:      t,
		Status:    status,
	}
	tx.Hash, _ = tx.EncodeTransaction()
	return tx
}

func (tx *Transaction) EncodeTransaction() (string, error) {
	h := sha256.New()

	if len(tx.From) == 0 {
		return "", encodeTransactionMissFromError
	}
	h.Write([]byte(tx.From))

	h.Write([]byte(tx.To))

	if len(tx.FileID) == 0 {
		return "", encodeTransactionMissFileIDError
	}
	h.Write([]byte(tx.FileID))

	if len(tx.Timestamp) == 0 {
		return "", encodeTransactionMissTimestampError
	}
	h.Write([]byte(tx.Timestamp))

	h.Write([]byte(strconv.Itoa(int(tx.Type))))

	h.Write([]byte(strconv.Itoa(int(tx.Status))))

	sum := h.Sum(nil)
	return hex.EncodeToString(sum), nil
}

func (tx *Transaction) ToMessage() *msg.Transaction {
	return &msg.Transaction{
		Hash:      tx.Hash,
		From:      tx.From,
		To:        tx.To,
		Value:     tx.Value,
		FileID:    tx.FileID,
		Timestamp: tx.Timestamp,
		Type:      tx.Type,
		Status:    tx.Status,
	}
}

type Transactions struct {
	Txs []*Transaction `orm:"-"`
}

func (txs Transactions) Len() int {
	return len(txs.Txs)
}

func (txs Transactions) Swap(i, j int) {
	txs.Txs[i], txs.Txs[j] = txs.Txs[j], txs.Txs[i]
}

type SortByHash struct{ Transactions }

func (p SortByHash) Less(i, j int) bool {
	return p.Transactions.Txs[i].Timestamp > p.Transactions.Txs[j].Timestamp
}
