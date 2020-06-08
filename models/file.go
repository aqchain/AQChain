package models

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
)

// 文件记录
type File struct {
	ID string `orm:"pk;column(ID)" json:"id"` // 加密hash 这个应当是数据保全产生的ID

	// 记录上传者
	Creator    string `orm:"column(creator)" json:"creator"`         // 创建者
	CreateTime string `orm:"column(create_time)" json:"create_time"` // 时间

	// 记录交易
	Owner          string  `orm:"column(owner)" json:"owner"`                     // 当前拥有者
	Value          float64 `orm:"column(value)" json:"value"`                     // 价格
	TransferRecord string  `orm:"column(transfer_record)" json:"transfer_record"` // 所有权转移记录

	ContentHash string `orm:"column(content_hash)" json:"content_hash"` // 文件内容Hash
}

// 对Creator CreateTime FileID 进行hash作为文件ID
func (file *File) EncodeFile() (string, error) {
	h := sha256.New()
	if len(file.Creator) == 0 {
		return "", errors.New("encodeFile miss Creator")
	}
	h.Write([]byte(file.Creator))
	if len(file.CreateTime) == 0 {
		return "", errors.New("encodeFile miss CreateTime")
	}
	h.Write([]byte(file.CreateTime))
	if len(file.ContentHash) == 0 {
		return "", errors.New("encodeFile miss ContentHash")
	}
	h.Write([]byte(file.ContentHash))
	sum := h.Sum(nil)
	return hex.EncodeToString(sum), nil
}
