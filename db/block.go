package db

import (
	"AQChain/models"
	"fmt"
	"github.com/astaxie/beego/orm"
	"log"
)

func ReadBlockByNumber(number int64) *models.Block {
	block := &models.Block{
		Number: number,
	}
	err := orm.NewOrm().Read(block)
	if err != nil {
		if err != orm.ErrNoRows {
			log.Println(err)
		}
		return nil
	}
	return block
}

func ReadBlockByHash(hash string) *models.Block {
	block := &models.Block{
		Hash: hash,
	}
	err := orm.NewOrm().Read(block)
	if err != nil {
		if err != orm.ErrNoRows {
			log.Println(err)
		}
		return nil
	}
	return block
}

func ReadLastBlock() *models.Block {
	block := &models.Block{}
	err := orm.NewOrm().Raw("select * from block where number = (select max(number) from block) ").QueryRow(block)
	if err != nil {
		if err != orm.ErrNoRows {
			log.Println(err)
		}
		return nil
	}
	return block
}

func ReadBlocks(from int64, len int64) []models.Block {
	var blocks []models.Block
	_, err := orm.NewOrm().QueryTable(&models.Block{}).Limit(len, from).All(&blocks)
	if err != nil {
		log.Println(err)
	}
	return blocks
}

func ReadBlocksByCreator(creator string) []*models.Block {
	var blocks []*models.Block
	_, err := orm.NewOrm().QueryTable(&models.Block{}).Filter("creator", creator).All(&blocks)
	if err != nil {
		log.Println(err)
		return nil
	}
	return blocks
}

// 保存新区块 重要！
func NewBlock(block models.Block) error {
	db := orm.NewOrm()

	// 异常处理 rollback
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("更新区块数据错误：", err)
			db.Rollback()
		}
	}()

	db.Begin()
	// 保存交易
	for _, tx := range block.Txs {
		// 保存交易
		tx.BlockNumber = block.Number
		_, err := db.Insert(tx)
		if err != nil {
			panic(err)
		}
		// 两种类型判断
		if tx.Type == 0 {
			// 上传添加file表
			err = insertFile(db, tx.FileID, tx.From, tx.Timestamp)
			if err != nil {
				panic(err)
			}
			// 账户数据更新
			if readAccount(db, tx.From) == nil {
				newAccount(db, tx.From)
			}
			err = addContribution(db, tx.From, 1, 0)
			if err != nil {
				panic(err)
			}
		} else {
			// 生成请求
			transfer := models.TransferFromTx(tx)
			// 判断交易状态
			switch tx.Status {
			case 0:
				// 创建交易请求
				err := InsertTransfer(db, transfer)
				if err != nil {
					panic(err)
				}
			case 1:
				// 确认交易请求
				err := UpdateTransferConfirm(db, transfer)
				if err != nil {
					panic(err)
				}
			case 2:
				// 完成交易请求
				err := UpdateTransferComplete(db, transfer)
				if err != nil {
					panic(err)
				}
				// 交易完成更新文件表
				err = updateFile(db, transfer.FileID, transfer.From, transfer.ID, transfer.Value)
				if err != nil {
					panic(err)
				}

				// 账户数据更新
				if readAccount(db, tx.From) == nil {
					newAccount(db, tx.From)
				}
				// 缺少海绵函数
				err = addContribution(db, tx.From, 1, 0)
				if err != nil {
					panic(err)
				}
			}
		}
	}
	// 保存区块
	_, err := db.Insert(&block)
	if err != nil {
		panic(err)
	}

	db.Commit()
	return nil
}

func GenesisBlock() models.Block {
	genesisBlock := ReadBlockByNumber(0)
	if genesisBlock == nil {
		// 加一个交易解决产生贡献值解决第一个区块生成问题 问题这个admin账户是否会影响？
		genesisBlock = &models.Block{
			Number:       0,
			Hash:         "genesisBlock",
			Creator:      "admin",
			PrevHash:     "",
			Timestamp:    "2006-01-02 15:04:05",
			MerkleRoot:   "",
			Transactions: models.Transactions{},
		}
		_ = NewBlock(*genesisBlock)
	}

	return *genesisBlock
}
