package p2p

import (
	"AQChain/models"
	"AQChain/models/serialize"
	"bufio"
	"container/list"
	"context"
	"encoding/json"
	"fmt"
	"github.com/libp2p/go-libp2p"
	core "github.com/libp2p/go-libp2p-core"
	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/peerstore"
	ma "github.com/multiformats/go-multiaddr"
	"log"
	"net"
	"strconv"
	"strings"
	"sync"
	"time"
)

// 地址本
type Peer struct {
	Account   string
	Address   string
	PublicKey string
	IsOnline  bool
}

type P2P struct {
	// 账户地址
	Account   string
	PublicKey string
	// libp2p地址
	Address string

	Peers         map[string]Peer
	ServerAddress string

	recv chan string
	send *list.List

	host core.Host

	mutex sync.Mutex
}

func NewP2P(account string) P2P {
	return P2P{
		Account:       account,
		Peers:         make(map[string]Peer, 0),
		ServerAddress: "127.0.0.1:19093",
		recv:          make(chan string, 100),
		send:          list.New(),
		mutex:         sync.Mutex{},
	}
}

// 创建Host
func (p *P2P) makeBasicHost(listenPort int) core.Host {

	opts := []libp2p.Option{
		libp2p.ListenAddrStrings(fmt.Sprintf("/ip4/127.0.0.1/tcp/%d", listenPort)),
	}
	host, _ := libp2p.New(context.Background(), opts...)

	ma, _ := ma.NewMultiaddr(fmt.Sprintf("/ipfs/%s", host.ID().Pretty()))
	p.Address = host.Addrs()[0].Encapsulate(ma).String()
	log.Printf("我的连接地址是 %s\n", p.Address)

	host.SetStreamHandler("/p2p/1.0.0", p.handleStream)

	return host
}

// 流处理
func (p *P2P) handleStream(s network.Stream) {
	rw := bufio.NewReadWriter(bufio.NewReader(s), bufio.NewWriter(s))
	go p.readData(rw)
	go p.writeData(rw)
}

// 读数据存入recv
func (p *P2P) readData(rw *bufio.ReadWriter) {
	var str string
	for {
		str, _ = rw.ReadString(23)

		if str == "" {
			return
		}
		if str != string(23) {
			p.recv <- str
		}
	}
}

func (p *P2P) writeData(rw *bufio.ReadWriter) {
	var sendchan = make(chan string, 100)
	p.send.PushBack(&sendchan)

	var str string
	for {
		str = <-sendchan
		p.mutex.Lock()
		rw.WriteString(str)
		rw.Flush()
		p.mutex.Unlock()
	}
}

// 选通道中第一个发送
func (p *P2P) P2PSendOne(str string) {
	if strings.Contains(str, string(17)+string(18)) {
		return
	}
	str = strings.Replace(str, string(23), string(17)+string(18), -1)
	str = str + string(23)

	for p := p.send.Front(); p != nil; p = p.Next() {
		sc, ok := (p.Value).(*chan string)
		if ok {
			//将str的数据存入到sc中
			*sc <- str
			break
		}
	}
}

func (p *P2P) P2PSend(str string) {

	//判断 string(17)+string(18)是否再str中
	//string(17)是设备控制一，string(18)是设备控制二
	if strings.Contains(str, string(17)+string(18)) {
		return
	}
	//将str中的string(23)用string(17)+string(18)替换
	//string(23)是区块传输结束
	str = strings.Replace(str, string(23), string(17)+string(18), -1)
	//在str的末尾加上string(23)
	str = str + string(23)
	//如果SendList列表为空，则返回第一个
	for p := p.send.Front(); p != nil; p = p.Next() {
		sc, ok := (p.Value).(*chan string)
		if ok {
			//将str的数据存入到sc中
			*sc <- str
		}
	}
	p.recv <- str
}

func (p *P2P) P2PRecv() string {
	//从通道中取出消息
	str := <-p.recv
	str = strings.TrimRight(str, string(23))
	str = strings.Replace(str, string(17)+string(18), string(23), -1)
	return str
}

func (p *P2P) connectTarget(target string) {
	// 从target参数解析ipfs地址
	ipfsaddr, _ := ma.NewMultiaddr(target)
	pid, _ := ipfsaddr.ValueForProtocol(ma.P_IPFS)
	peerid, _ := peer.IDB58Decode(pid)

	// 解析PeerID和TargetID
	targetPeerAddr, _ := ma.NewMultiaddr(
		fmt.Sprintf("/ipfs/%s", peer.IDB58Encode(peerid)))

	targetAddr := ipfsaddr.Decapsulate(targetPeerAddr)

	// 加入节点列表
	p.host.Peerstore().AddAddr(peerid, targetAddr, peerstore.PermanentAddrTTL)

	// 连接目标建立流
	s, err := p.host.NewStream(context.Background(), peerid, "/p2p/1.0.0")
	if err != nil {
		fmt.Println("连接失败")
		fmt.Println(err)
		return
	}
	fmt.Println("已连接")

	// 建立流
	rw := bufio.NewReadWriter(bufio.NewReader(s), bufio.NewWriter(s))

	// 开始读写
	go p.writeData(rw)
	go p.readData(rw)

}

func (p *P2P) UpdatePeersFromServer() {

	for {
		for {
			time.Sleep(1e9)
			conn, err := net.Dial("tcp", p.ServerAddress)
			if err != nil {
				log.Println("Server连接错误")
				log.Println(err)
				return
			}
			str, err := bufio.NewReader(conn).ReadString('\n')
			if err != nil {
				log.Println("读取Server数据错误")
				log.Println(err)
				return
			}
			str = strings.Replace(str, "\n", "", -1)
			conn.Close()

			peersJson, _ := json.Marshal(p.Peers)

			if str != string(peersJson) {
				log.Println("更新Peers")
				err = json.Unmarshal([]byte(str), &p.Peers)
				if err != nil {
					log.Print(err)
				}
				p.PrintNodeList()
				break
			}
			time.Sleep(1e9)
		}
	}
}

func (p *P2P) PrintNodeList() {
	log.Printf("Peers的长度为：%v \n", len(p.Peers))
	for _, peer := range p.Peers {
		log.Printf("ID: %s  在线情况: %v", peer.Account, peer.IsOnline)
	}
}

// 从Server獲取節點
func (p *P2P) GetPeersFromServer() {
	conn, err := net.Dial("tcp", p.ServerAddress)
	if err != nil {
		log.Println("Server连接错误")
		log.Println(err)
		return
	}
	str, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		log.Print("读取Server数据錯誤")
		log.Println(err)
		return
	}
	str = strings.Replace(str, "\n", "", -1)
	conn.Close()
	_ = json.Unmarshal([]byte(str), &p.Peers)
}

// 將自己的信息發到Server
func (p *P2P) SendPeerToServer() {
	conn, err := net.Dial("tcp", p.ServerAddress)
	if err != nil {
		log.Println("Server连接错误")
		log.Println(err)
		return
	}
	_, err = bufio.NewReader(conn).ReadString('\n')
	_, err = fmt.Fprintln(conn, p.Account+"|"+p.Address+"|"+p.PublicKey)
	if err != nil {
		log.Println("寫入Server數據错误")
		log.Println(err)
		return
	}
	conn.Close()
}

func (p *P2P) Connect() {

	// 從Server更新peers
	p.GetPeersFromServer()
	// 判斷是否有之前的登陸
	peer := p.Peers[p.Account]
	var port int
	if len(peer.Account) > 0 {
		port, _ = strconv.Atoi(strings.Split(peer.Address, "/")[4])
	} else {
		port = 1000 + len(p.Peers)
	}
	// 創建Host
	p.host = p.makeBasicHost(port)
	// 發送新地址
	p.SendPeerToServer()

	// 連接節點
	if len(p.Peers) != 0 {
		for _, peer := range p.Peers {
			if peer.IsOnline == true {
				p.connectTarget(peer.Address)
			}
		}
	}
}

func (p *P2P) SendTransaction(tx *models.Transaction) {
	msg := string(serialize.SerializeTx(tx))
	msg = "mn" + msg
	p.P2PSend(msg)
	log.Printf("发送新交易 交易类型: %v", tx.Type)
}
