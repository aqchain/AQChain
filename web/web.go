package web

import (
	"AQChain/p2p"
	"AQChain/web/impl"
	"github.com/astaxie/beego"
	"github.com/gorilla/websocket"
	"github.com/pkg/browser"
	"log"
	"net/http"
	"os"
	"sync"
)

// web 负责管理了与web端交互
type Web struct {
	Address string

	p2pnet p2p.P2P

	// 连接
	wsConn *websocket.Conn
	conn   *impl.Connection
	//转换器
	upGrader websocket.Upgrader
	// 锁 在一个页面打开websocket时锁定
	wsLock sync.Mutex
}

// 初始化Web
func NewWeb(p p2p.P2P) (web Web) {
	web = Web{
		Address: os.Getenv("address"),
		p2pnet:  p,
		wsLock:  sync.Mutex{},
		upGrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool { //允许跨域
				return true
			},
		},
	}
	return
}

func (w Web) Start() {

	go func() {
		//HandleFunc 的第一个参数指的是请求路径，第二个参数是一个函数类型，表示这个请求需要处理的事情。
		http.HandleFunc("/ws/set", w.WSHandler) //配置路由,ws代表websocket

		http.HandleFunc("/http/account/myAccount", w.AccountMyAccountHandler)
		http.HandleFunc("/http/blockChain/block", w.BlockChainBlockHandler)
		http.HandleFunc("/http/blockChain/myBlock", w.BlockChainMyBlockHandler)
		http.HandleFunc("/http/blockChain/transaction", w.BlockChainTransactionHandler)
		http.HandleFunc("/http/file/upload", w.FileUploadHandler)
		http.HandleFunc("/http/file/list", w.FileListHandler)
		http.HandleFunc("/http/file/myFile", w.FileMyFileHandler)
		http.HandleFunc("/http/file/create", w.TransferCreateHandler)
		http.HandleFunc("/http/transfer/create", w.TransferCreateHandler)
		http.HandleFunc("/http/transfer/confirm", w.TransferConfirmHandler)
		http.HandleFunc("/http/transfer/complete", w.TransferCompleteHandler)
		http.HandleFunc("/http/transfer/unconfirmedList", w.TransferUnconfirmedListHandler)
		http.HandleFunc("/http/transfer/uncompletedList", w.TransferUncompletedListHandler)
		log.Println("AQChain client address :" + w.Address)
		//第一个参数是监听的端口、第二个参数是根页面的处理函数，可以为空。
		http.ListenAndServe(w.Address, nil)
	}()

	go beego.Run("127.0.0.1:8080")

	browser.OpenURL("http://" + "127.0.0.1:8080" + "/static/views/index.html")
}
