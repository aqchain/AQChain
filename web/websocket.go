package web

import (
	"AQChain/web/impl"
	"encoding/json"
	"log"
	"net/http"
)

const (
	BlockChain = "Blockchain"
)

// websocket 请求的消息格式 页面请求时需要注意格式
type reqMsg struct {
	// 请求的标识  通过不同的标识判断使用什么方法响应 同时转换json数据时需要判断结构体
	Code string `json:"code"`
	// 请求参数 json格式
	Data string `json:"data"`
}

type respMsg struct {
	// 错误标识 发生错误时向web
	Err string `json:"err"`
	// 返回数据 json格式
	Data interface{} `json:"data"`
}

type AccountInfo struct {
	AccountID string `json:"accountID"`
	Address   string `json:"address"`
}

// wsHandler中一定要注意 running的修改 一旦有error 不能panic
func (web *Web) WSHandler(w http.ResponseWriter, r *http.Request) {

	log.Println("WebSocket")
	var err error

	if web.wsConn, err = web.upGrader.Upgrade(w, r, nil); err != nil {
		log.Println(err)
		return
	}
	if web.conn, err = impl.InitConnection(web.wsConn); err != nil { //方法的调用为包名.方法名
		log.Println(err)
		return
	}

	// 连接后返回用户信息
	jsonMsg, _ := json.Marshal(respMsg{Data: AccountInfo{
		AccountID: web.p2pnet.Account,
		Address:   web.Address,
	}})
	web.conn.WriteMassage(jsonMsg)

	// 读取消息
	var msg []byte
	for {
		if msg, err = web.conn.ReadMessage(); err != nil {
			if err != nil {
				log.Println(err)
				web.reset()
				return
			}
		}
		// 转换消息格式
		req := &reqMsg{}
		log.Println(string(msg))
		err = json.Unmarshal(msg, req)
		if err != nil {
			log.Println(err)
			continue
		}
		// 处理请求
		if req.Code != "" {
			switch req.Code {

			case BlockChain: //区块建立记录页面
				web.wsLock.Lock()
				log.Println("前端打开区块建立记录页面")
				web.blockChainPage()
				web.wsLock.Unlock()
			default:
				web.errMsg(req)
			}

		}

	}
}

func (web *Web) blockChainPage() {
	jsonByte, err := json.Marshal(web.p2pnet.Account)
	if err != nil {
		log.Println(err)
		return
	}
	resp := respMsg{
		Err:  "",
		Data: string(jsonByte),
	}
	msg, err := json.Marshal(resp)
	if err != nil {
		log.Println(err)
	}
	_ = web.conn.WriteMassage(msg)
}

func (web *Web) reset() {
	web.conn.Close()
	_ = web.wsConn.Close()
}

// 处理错误消息
func (web *Web) errMsg(req *reqMsg) {
	log.Println("错误的web请求标识")
	log.Println(req.Code)
	msg, err := json.Marshal(respMsg{
		Err:  "错误的web请求标识",
		Data: "",
	})
	if err != nil {
		log.Println(err)
	}
	_ = web.conn.WriteMassage(msg)
}
