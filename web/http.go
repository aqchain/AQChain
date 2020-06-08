package web

import (
	"AQChain/crypto"
	"AQChain/db"
	"AQChain/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

type result struct {
	Code  int         `json:"code"`
	Msg   string      `json:"msg"`
	Count int         `json:"count"`
	Data  interface{} `json:"data"`
}

func (web *Web) AccountMyAccountHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	account := db.ReadAccount(web.p2pnet.Account)
	var accounts = make([]*models.Account, 0)
	if account != nil {
		accounts = append(accounts, account)
	}
	result := result{
		Code:  0,
		Msg:   "",
		Count: len(accounts),
		Data:  accounts,
	}
	jsonResult, err := json.Marshal(result)
	if err != nil {
		log.Println(err)
	}
	_, _ = w.Write(jsonResult)
}

func (web *Web) BlockChainBlockHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	lastBlock := db.ReadLastBlock()
	blocks := db.ReadBlocks(0, lastBlock.Number+1)

	result := result{
		Code:  0,
		Msg:   "",
		Count: len(blocks),
		Data:  blocks,
	}
	jsonResult, err := json.Marshal(result)
	if err != nil {
		log.Println(err)
	}
	_, _ = w.Write(jsonResult)
}

func (web *Web) BlockChainMyBlockHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	blocks := db.ReadBlocksByCreator(web.p2pnet.Account)

	result := result{
		Code:  0,
		Msg:   "",
		Count: len(blocks),
		Data:  blocks,
	}
	jsonResult, err := json.Marshal(result)
	if err != nil {
		log.Println(err)
	}
	_, _ = w.Write(jsonResult)
}

func (web *Web) BlockChainTransactionHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	var txs []models.Transaction
	str := r.FormValue("blockNumber")
	if str == "false" {
		txs = db.ReadTransactions()
	} else {
		blockNumber, _ := strconv.ParseInt(str, 10, 64)
		txs = db.ReadTransactionsByBlockNumber(blockNumber)
	}

	result := result{
		Code:  0,
		Msg:   "",
		Count: len(txs),
		Data:  txs,
	}
	jsonResult, err := json.Marshal(result)
	if err != nil {
		log.Println(err)
	}
	_, _ = w.Write(jsonResult)
}

type fileResult struct {
	FileID string `json:"file_id"`
	TxHash string `json:"tx_hash"`
}

func (web *Web) FileUploadHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	uploadFile, _, err := r.FormFile("file")
	if err != nil {
		fmt.Println(err)
	}
	now := time.Now().String()
	// 文件内容hash
	contentHash, _ := crypto.SHA256UploadFile(uploadFile)
	file := models.File{
		Creator:     web.p2pnet.Account,
		CreateTime:  now,
		ContentHash: contentHash,
	}
	// 文件ID
	fileID, err := file.EncodeFile()
	if err != nil {
		fmt.Println(err)
	}

	// 创建并发送交易
	tx := models.NewTransaction(web.p2pnet.Account, "", fileID, 0, 0, 0)
	web.p2pnet.SendTransaction(tx)

	jsonResult, _ := json.Marshal(result{
		Code: 0,
		Msg:  "",
		Data: fileResult{
			FileID: fileID,
			TxHash: tx.Hash,
		},
	})
	w.Write(jsonResult)

}

func (web *Web) TransferCreateHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	fileID := r.FormValue("fileID")
	price, err := strconv.ParseFloat(r.FormValue("price"), 2)
	if err != nil {
		log.Println(err)
		return
	}

	// 根据fileID获取file记录
	file := db.ReadFile(fileID)
	// 有个问题没有考虑 就是已经产生最终交易但未出块的情况

	if file == nil {
		log.Println("未查询到fileID")
		return
	}
	// 创建交易请求并发送
	tx := models.NewTransaction(web.p2pnet.Account, file.Owner, file.ID, price, 1, 0)
	web.p2pnet.SendTransaction(tx)

	jsonResult, _ := json.Marshal(result{
		Code: 0,
		Msg:  "",
		Data: tx.Hash,
	})
	w.Write(jsonResult)
}

func (web *Web) TransferConfirmHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	transferID := r.FormValue("id")
	transfer := db.ReadTransfer(transferID)
	// 创建交易请求并发送
	tx := models.NewTransaction(transfer.From, transfer.To, transfer.FileID, transfer.Value, 1, 1)
	web.p2pnet.SendTransaction(tx)

	jsonResult, _ := json.Marshal(result{
		Code: 0,
		Msg:  "",
		Data: tx.Hash,
	})
	w.Write(jsonResult)
}

func (web *Web) TransferCompleteHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	transferID := r.FormValue("id")
	transfer := db.ReadTransfer(transferID)
	// 创建交易请求并发送
	tx := models.NewTransaction(transfer.From, transfer.To, transfer.FileID, transfer.Value, 1, 2)
	web.p2pnet.SendTransaction(tx)

	jsonResult, _ := json.Marshal(result{
		Code: 0,
		Msg:  "",
		Data: tx.Hash,
	})
	w.Write(jsonResult)
}

func (web *Web) TransferUnconfirmedListHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	transfers := db.ReadUnconfirmedTransfers(web.p2pnet.Account)
	result := result{
		Code:  0,
		Msg:   "",
		Count: len(transfers),
		Data:  transfers,
	}
	jsonResult, err := json.Marshal(result)
	if err != nil {
		log.Println(err)
	}
	_, _ = w.Write(jsonResult)
}

func (web *Web) TransferUncompletedListHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	transfers := db.ReadUncompletedTransfers(web.p2pnet.Account)
	result := result{
		Code:  0,
		Msg:   "",
		Count: len(transfers),
		Data:  transfers,
	}
	jsonResult, err := json.Marshal(result)
	if err != nil {
		log.Println(err)
	}
	_, _ = w.Write(jsonResult)
}

func (web *Web) FileListHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	files := db.ReadFiles()

	result := result{
		Code:  0,
		Msg:   "",
		Count: len(files),
		Data:  files,
	}
	jsonResult, err := json.Marshal(result)
	if err != nil {
		log.Println(err)
	}
	_, _ = w.Write(jsonResult)
}

func (web *Web) FileMyFileHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	files := db.ReadFilesByOwner(web.p2pnet.Account)
	result := result{
		Code:  0,
		Msg:   "",
		Count: len(files),
		Data:  files,
	}
	jsonResult, err := json.Marshal(result)
	if err != nil {
		log.Println(err)
	}
	_, _ = w.Write(jsonResult)
}
