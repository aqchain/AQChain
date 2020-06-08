package serialize

import (
	"AQChain/models"
	msg "AQChain/models/pb"
	"bytes"
	"encoding/gob"
	"fmt"
	"github.com/golang/protobuf/proto"
	"log"
)

func SerializeTx(tx *models.Transaction) []byte {
	msgTx := tx.ToMessage()
	data, err := proto.Marshal(msgTx)
	if err != nil {
		log.Fatalln("Marshal Data error: ", err)
	}
	return data
}

func DeserializeTx(b []byte) *models.Transaction {
	var tx models.Transaction
	var msgTx msg.Transaction

	err := proto.Unmarshal(b, &msgTx)
	if err != nil {
		log.Fatalln("Unmarshal Data error: ", err)
	}
	tx.Hash = msgTx.Hash
	tx.From = msgTx.From
	tx.To = msgTx.To
	tx.FileID = msgTx.FileID
	tx.Value = msgTx.Value
	tx.Timestamp = msgTx.Timestamp
	tx.Type = msgTx.Type
	tx.Status = msgTx.Status

	return &tx
}

func SerializeTxs(b []*models.Transaction) []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)
	err := encoder.Encode(b)
	if err != nil {
		log.Panic(err)
	}
	return result.Bytes()
}

func SerializeBlock(block models.Block) []byte {
	/*var result bytes.Buffer
	encoder := gob.NewEncoder(&result)
	err := encoder.Encode(b)
	if err != nil {
		log.Panic(err)
	}
	return result.Bytes()*/
	blockMsg := msg.Block{
		Number:     block.Number,
		Hash:       block.Hash,
		Creator:    block.Creator,
		PrevHash:   block.PrevHash,
		Timestamp:  block.Timestamp,
		MerkleRoot: block.MerkleRoot,
	}

	txs := make([]*msg.Transaction, 0, len(block.Txs))

	for _, tx := range block.Txs {
		txmsg := tx.ToMessage()
		txs = append(txs, txmsg)
	}
	blockMsg.Txs = txs
	result, err := proto.Marshal(&blockMsg)
	if err != nil {
		log.Fatalln("Marshal Data error: ", err)
	}
	return result
}

func DeserializeBlock(d []byte) *models.Block {
	/*var b models.Block
	decoder := gob.NewDecoder(bytes.NewReader(d))
	err := decoder.Decode(&b)
	if err != nil {
		log.Panic(err)
	}
	return &b*/
	var target msg.Block

	err := proto.Unmarshal(d, &target)
	if err != nil {
		log.Fatalln("Unmarshal Data error: ", err)
	}

	b := &models.Block{
		Number:       target.Number,
		Hash:         target.Hash,
		Creator:      target.Creator,
		PrevHash:     target.PrevHash,
		Timestamp:    target.Timestamp,
		MerkleRoot:   target.MerkleRoot,
		Transactions: models.Transactions{},
	}

	txs := make([]*models.Transaction, 0, len(target.Txs))

	for _, msgTx := range target.Txs {
		tx := &models.Transaction{
			Hash:      msgTx.Hash,
			From:      msgTx.From,
			To:        msgTx.To,
			Value:     msgTx.Value,
			FileID:    msgTx.FileID,
			Timestamp: msgTx.Timestamp,
			Type:      msgTx.Type,
			Status:    msgTx.Status,
		}
		txs = append(txs, tx)
	}
	b.Txs = txs
	return b
}

func SerializeAccount(u models.Account) []byte {
	user := msg.Account{
		Address:          u.Address,
		Contribution:     u.Contribution,
		ContributionFile: u.ContributionFile,
		ContributionTx:   u.ContributionTx,
	}

	result, err := proto.Marshal(&user)
	if err != nil {
		log.Fatalln("Marshal Data error: ", err)
	}

	return result
}

func DeserializeAccount(d []byte) *models.Account {
	var u models.Account
	var target msg.Account

	err := proto.Unmarshal(d, &target)
	if err != nil {
		log.Fatalln("Unmarshal Data error: ", err)
	}

	u.Contribution = target.Contribution
	u.ContributionTx = target.ContributionTx
	u.ContributionFile = target.ContributionFile
	u.Address = target.Address

	return &u
}

func SerializeBlockChain(bc []models.Block) []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)
	err := encoder.Encode(bc)
	if err != nil {
		log.Panic(err)
	}
	return result.Bytes()
}

func DeserializeBlockChain(b []byte) (bc []models.Block) {
	bc = make([]models.Block, 0)
	decoder := gob.NewDecoder(bytes.NewReader(b))
	fmt.Println(decoder)
	err := decoder.Decode(&bc)
	if err != nil {
		log.Panic(err)
	}
	return bc
}

func SerializeTransaction(mns []*models.Transaction) []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)
	err := encoder.Encode(mns)
	if err != nil {
		log.Panic(err)
	}
	return result.Bytes()
}

func DeserializeTransaction(b []byte) (mns []*models.Transaction) {
	mns = make([]*models.Transaction, 0)
	decoder := gob.NewDecoder(bytes.NewReader(b))
	fmt.Println(decoder)
	err := decoder.Decode(&mns)
	if err != nil {
		log.Panic(err)
	}
	return mns
}
