package models

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
)

// 记录交易
type Transfer struct {
	ID         string  `orm:"pk;column(id)" json:"id"`                // 交易请求ID hash
	From       string  `orm:"column(creator)" json:"from"`            // 对应Transaction.From
	To         string  `orm:"column(owner)" json:"to"`                // 对应Transaction.To
	Value      float64 `orm:"column(value)" json:"value"`             // 对应Transaction.Value
	FileID     string  `orm:"column(file_id)" json:"file_id"`         // 文件ID
	CreateTx   string  `orm:"column(create_tx)" json:"create_tx"`     // 创建请求的Transaction.Hash
	ConfirmTx  string  `orm:"column(confirm_tx)" json:"confirm_tx"`   // 确认请求的Transaction.Hash
	CompleteTx string  `orm:"column(complete_tx)" json:"complete_tx"` // 完成请求的Transaction.Hash
	Status     int32   `orm:"column(status)" json:"status"`           // 0表示为确认的交易、1表示已经确认了交易、2表示完成转账操作的交易
}

// 对Creator CreateTime FileID 进行hash作为文件ID
func (req *Transfer) EncodeTransfer() (string, error) {
	h := sha256.New()
	if len(req.From) == 0 {
		return "", errors.New("EncodeTransfer miss From")
	}
	h.Write([]byte(req.From))
	if len(req.To) == 0 {
		return "", errors.New("EncodeTransfer miss To")
	}
	h.Write([]byte(req.To))
	if len(req.FileID) == 0 {
		return "", errors.New("EncodeTransfer miss FileID")
	}
	h.Write([]byte(req.FileID))
	sum := h.Sum(nil)
	return hex.EncodeToString(sum), nil
}

func TransferFromTx(tx *Transaction) *Transfer {
	if tx.Type != 1 {
		return nil
	}
	req := &Transfer{
		From:   tx.From,
		To:     tx.To,
		Value:  tx.Value,
		FileID: tx.FileID,
	}
	switch tx.Status {
	case 0:
		req.CreateTx = tx.Hash
	case 1:
		req.ConfirmTx = tx.Hash
	case 2:
		req.CompleteTx = tx.Hash
	}
	req.ID, _ = req.EncodeTransfer()
	return req
}
