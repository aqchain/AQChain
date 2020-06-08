package main

import (
	"AQChain/p2p"
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
	"time"
)

var port = "19093"

func listenLoop(peers map[string]p2p.Peer) {

	listener, _ := net.Listen("tcp", "127.0.0.1:"+port)
	for {
		conn, _ := listener.Accept()
		go handle(conn, peers)
	}

}

func handle(conn net.Conn, peers map[string]p2p.Peer) {

	// 将NodeList发送
	nl, _ := json.Marshal(peers)
	_, _ = fmt.Fprintln(conn, string(nl))

	// 读取
	recv, _ := bufio.NewReader(conn).ReadString('\n')
	// 没有内容结束
	if recv == "" {
		conn.Close()
		return
	}
	recv = recv[:len(recv)-1]
	// 截取信息
	infos := strings.Split(recv, "|")
	account := infos[0]
	address := infos[1]
	publicKey := infos[2]
	// 更新
	peers[account] = p2p.Peer{
		Account:   account,
		Address:   address,
		PublicKey: publicKey,
		IsOnline:  true,
	}

	fmt.Println("新上线的节点是 账户" + account)
	fmt.Println("p2p地址 " + address)
	conn.Close()
}

func main() {
	peers := make(map[string]p2p.Peer)
	go listenLoop(peers)
	go detectLoop(peers)
	select {}
}

func detectLoop(peers map[string]p2p.Peer) {
	fmt.Println("探测节点是否在线")
	for {
		time.Sleep(3e9)
		for _, peer := range peers {
			if peer.IsOnline == true {
				str := strings.Split(peer.Address, "/")
				port, _ := strconv.Atoi(str[4])
				addr := &net.TCPAddr{IP: net.ParseIP(str[2]), Port: port}
				conn, err := net.DialTCP("tcp", nil, addr)
				if err != nil {
					peer.IsOnline = false
					peers[peer.Account] = peer
					log.Print("节点 " + peer.Account + "已离线")
				} else {
					conn.Close()
				}
			}

		}
	}
}
