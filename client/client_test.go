package client

import (
	"AQChain/db"
	"AQChain/models"
	"AQChain/models/serialize"
	"AQChain/p2p"
	"AQChain/utils"
	web2 "AQChain/web"
	"github.com/davecgh/go-spew/spew"
	assert2 "github.com/magiconair/properties/assert"
	"github.com/pkg/browser"
	"github.com/stretchr/testify/assert"
	"net/http"
	"os"
	"strings"
	"testing"
)

func init() {
	username := "test"
	path := utils.CaculatePath() + username + ".db"
	os.Remove(path)
	db.InitDB(username)
}

func mockBlock() models.Block {
	file := models.File{
		ID:             "1",
		Creator:        "test",
		CreateTime:     "1",
		Owner:          "test",
		Value:          0,
		TransferRecord: "",
		ContentHash:    "test",
	}
	// 上传交易
	tx := models.NewTransaction(file.Creator, "", file.ID, 0, 0, 0)
	// 购买交易
	tx1 := models.NewTransaction("test2", "test", file.ID, 10, 1, 0)
	// 确认交易
	tx2 := models.NewTransaction("test2", "test", file.ID, 10, 1, 1)
	// 完成交易
	tx3 := models.NewTransaction("test2", "test", file.ID, 10, 1, 2)

	// ge
	gb := db.GenesisBlock()

	return models.NewBlock(gb, []*models.Transaction{tx, tx1, tx2, tx3}, "test")

}

func TestHttpHandler(t *testing.T) {
	web := web2.NewWeb(p2p.P2P{})
	http.HandleFunc("/http/file/upload", web.FileUploadHandler)
	http.HandleFunc("/http/file/list", web.FileListHandler)
	http.HandleFunc("/http/transfer/create", web.TransferCreateHandler)
	http.HandleFunc("/http/transfer/confirm", web.TransferConfirmHandler)
	http.HandleFunc("/http/transfer/complete", web.TransferCompleteHandler)
	http.HandleFunc("/http/transfer/unconfirmedList", web.TransferUnconfirmedListHandler)
	http.HandleFunc("/http/transfer/uncompletedList", web.TransferUncompletedListHandler)

	http.HandleFunc("/ws/set", web.WSHandler)
	http.ListenAndServe("127.0.0.1:9999", nil)
	select {}
}

func TestSerializeBlock(t *testing.T) {
	b := mockBlock()
	bts := serialize.SerializeBlock(b)

	str := "block11" + string(bts)
	// str2 := strings.ReplaceAll(str, "block11","")
	str2 := strings.Trim(str, "block11")
	assert2.Equal(t, string(bts), str2)
	b2 := serialize.DeserializeBlock([]byte(str2))

	assert2.Equal(t, b, b2)
}

func TestInsertNewBlock(t *testing.T) {
	err := db.NewBlock(mockBlock())
	assert.Nil(t, err)
	/*lb:= db.ReadLastBlock()
	spew.Dump(lb)*/

	var blocks []models.Block
	//_, _ = orm.NewOrm().QueryTable(&models.Block{}).Limit(2, 0).All(&blocks)
	blocks = db.ReadBlocks(0, 2)
	spew.Dump(blocks)

}

func TestHandleMessage(t *testing.T) {
	browser.OpenURL("http://localhost:7777/AQChain/page/views/index.html")
}
