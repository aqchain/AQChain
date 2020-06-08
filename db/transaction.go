package db

import (
	"AQChain/models"
	"github.com/astaxie/beego/orm"
	"log"
)

func ReadTransaction(hash string) *models.Transaction {
	tx := &models.Transaction{
		Hash: hash,
	}
	err := orm.NewOrm().Read(tx)
	if err != nil {
		if err != orm.ErrNoRows {
			log.Println(err)
		}
		return nil
	}
	return tx
}

func ReadTransactions() []models.Transaction {
	var txs []models.Transaction
	_, err := orm.NewOrm().QueryTable(&models.Transaction{}).All(&txs)
	if err != nil {
		if err != orm.ErrNoRows {
			log.Println(err)
		}
		return nil
	}
	return txs
}

func ReadTransactionsByBlockNumber(number int64) []models.Transaction {
	var txs []models.Transaction
	_, err := orm.NewOrm().QueryTable(&models.Transaction{}).Filter("block_number", number).All(&txs)
	if err != nil {
		if err != orm.ErrNoRows {
			log.Println(err)
		}
		return nil
	}
	return txs
}
