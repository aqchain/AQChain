package db

import (
	"AQChain/models"
	"errors"
	"github.com/astaxie/beego/orm"
	"log"
)

var RequestMissingID = errors.New("request miss request.ID")
var RequestMissingCreateTx = errors.New("request miss request.CreateTx")
var RequestMissingConfirmTx = errors.New("request miss request.ConfirmTx")
var RequestMissingCompleteTx = errors.New("request miss request.CompleteTx")

func ReadTransfer(id string) *models.Transfer {
	req := &models.Transfer{
		ID: id,
	}
	err := orm.NewOrm().Read(req)
	if err != nil {
		log.Println(err)
		return nil
	}
	return req
}

func ReadTransfersByFileID(fileID string) []*models.Transfer {
	var transfers []*models.Transfer
	_, err := orm.NewOrm().QueryTable(&models.Transfer{}).Filter("file_id", fileID).All(&transfers)
	if err != nil {
		log.Println(err)
		return nil
	}
	return transfers
}

// 查询账户未确认Transfer
func ReadUnconfirmedTransfers(account string) []*models.Transfer {
	var transfers []*models.Transfer
	_, err := orm.NewOrm().QueryTable(&models.Transfer{}).Filter("to", account).Filter("status", 0).All(&transfers)
	if err != nil {
		log.Println(err)
		return nil
	}
	return transfers
}

// 查询账户未确认Transfer
func ReadUncompletedTransfers(account string) []*models.Transfer {
	var transfers []*models.Transfer
	_, err := orm.NewOrm().QueryTable(&models.Transfer{}).Filter("from", account).Filter("status", 1).All(&transfers)
	if err != nil {
		log.Println(err)
		return nil
	}
	return transfers
}

// 添加Transfer
func InsertTransfer(orm orm.Ormer, transfer *models.Transfer) error {
	if len(transfer.ID) == 0 {
		return RequestMissingID
	}

	if len(transfer.CreateTx) == 0 {
		return RequestMissingCreateTx
	}

	_, err := orm.Insert(transfer)
	if err != nil {
		return err
	}
	return nil
}

// 更新Request到确认状态
func UpdateTransferConfirm(orm orm.Ormer, request *models.Transfer) error {
	if len(request.ID) == 0 {
		return RequestMissingID
	}

	if len(request.ConfirmTx) == 0 {
		return RequestMissingConfirmTx
	}

	request.Status = 1

	_, err := orm.Update(request, "status", "confirm_tx")
	if err != nil {
		return err
	}
	return nil
}

// 更新Request到确认状态
func UpdateTransferComplete(orm orm.Ormer, request *models.Transfer) error {
	if len(request.ID) == 0 {
		return RequestMissingID
	}

	if len(request.CompleteTx) == 0 {
		return RequestMissingConfirmTx
	}

	request.Status = 2

	_, err := orm.Update(request, "status", "complete_tx")
	if err != nil {
		return err
	}
	return nil
}
