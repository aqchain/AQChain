package core

import (
	"AQChain/crypto"
	"AQChain/db"
	"AQChain/models"
	"AQChain/models/serialize"
	"AQChain/p2p"
	"log"
	"os"
	"sort"
	"strings"
)

const (
	NewTxMessage    = "NTX"
	NewBlockMessage = "NBK"
)

type AQManager struct {
	txsPerBlock int
	chain       Chain
	account     models.Account
	peer        p2p.P2P
	txCache     *TxSet
}

func NewManager(txsPerBlock int, chain Chain, account models.Account, peer p2p.P2P, txCache *TxSet) AQManager {

	return AQManager{
		txsPerBlock: txsPerBlock,
		chain:       chain,
		account:     account,
		peer:        peer,
		txCache:     txCache,
	}
}

func (m *AQManager) HandleMessage(msg string) {

	/*// 小心这个判断strings.Contains可能出错 之后得改
	// 处理请求区块
	if strings.Contains(str, "GBH") {
		fmt.Println(" 同步区块请求")
		// 解析对方发来的区块
		str = strings.Trim(str, "GBH")
		b := serialize.DeserializeBlock([]byte(str))
		// 比较自己的和对方的lastBlock长度
		remoteLen := b.Number
		lastBlock := Blockchain[len(Blockchain)-1]
		localLen := lastBlock.Number
		fmt.Println(" remoteLen", remoteLen)
		fmt.Println(" localLen", localLen)
		bc := make([]models.Block, 0)
		// 返回本地比远程节点多的部分
		if localLen-remoteLen > 0 {
			bc = Blockchain[remoteLen-1 : localLen-1]
		}
		// 相等时返回最后一块用于对方确认
		if localLen-remoteLen == 0 {
			bc = append(bc, lastBlock)
		}
		bytes := serialize.SerializeBlockChain(bc)
		fmt.Println(" BH", string(bytes))
		// 向所有节点发送
		p.P2PSend("BH" + string(bytes))
		return
	}

	// 处理请求交易
	if strings.Contains(str, "GTX") {
		fmt.Println(" 同步交易请求")
		userName := strings.Trim(str, "GTX")
		// 把缓冲里的交易全发给对方
		bytes := serialize.SerializeTransaction(txCache.list)
		// 向所有节点发送 这一步因为是全网发送 需要签名了来判断是谁接收 先做发送了UserName来判断..十分简陋
		p.P2PSend("TX" + userName[0:8] + string(bytes))
		return
	}

	if strings.Contains(str, "TX") {
		fmt.Println(" 同步交易")
		str = strings.Trim(str, "TX")
		id := str[0:8]
		// 只有签名那个人处理这条消息
		if username[0:8] == id {
			txs := serialize.DeserializeTransaction([]byte(str[8:len(str)]))
			for _,tx:= range txs{
				txCache.Add(tx)
			}
		}
		fmt.Println(" 同步交易完成", txCache.Size())
		return
	}
	*/

	if strings.Contains(msg, "mn") {
		//新交易
		str := strings.Trim(msg, "mn")

		//将消息反序列化为数据结构“merkernode”
		newTx := serialize.DeserializeTx([]byte(str))

		// 添加到交易缓存
		err := m.txCache.Add(newTx)
		if err != nil {
			log.Println(err)
		}

		log.Printf(" \n缓存中交易数量	%d\n ", m.txCache.Size())

		// 判断是否生成区块
		if m.txCache.Size()%m.txsPerBlock == 0 {

			log.Println("是时候生成真正的区块了！")

			var validator models.Account
			// 第一块出块者问题
			if m.chain.blockChain[len(m.chain.blockChain)-1].Number == 0 {
				validator.Address = os.Getenv("validator")
			} else {
				// 这个位置可以优化
				validator = *db.ReadMostContributionAccount()
			}

			if m.account.Address == validator.Address { //如果应该生成新区块的是我
				newBlock := m.CreateNewBlock(m.txCache.List())

				a := serialize.SerializeBlock(newBlock)
				b := "block11" + string(a)

				m.peer.P2PSend(b)
				log.Println("发送新区块")
			}
		}
		return
	}

	if strings.Contains(msg, "block11") {
		log.Println("收到新的区块信息")
		str := strings.ReplaceAll(msg, "block11", "")

		//将消息反序列化为数据结构block
		newBlock := serialize.DeserializeBlock([]byte(str))

		// 排序
		sort.Sort(models.SortByHash{newBlock.Transactions})

		log.Println("验证块的校验信息")
		txs := make([]*models.Transaction, 0, len(newBlock.Transactions.Txs))
		for _, mn := range newBlock.Transactions.Txs {

			for _, tx := range m.txCache.List() {
				if mn.FileID == tx.FileID {
					txs = append(txs, tx)
					break
				}
			}
		}
		if len(txs) == len(newBlock.Transactions.Txs) &&
			crypto.SHA256Byte(serialize.SerializeTxs(newBlock.Transactions.Txs)) == crypto.SHA256Byte(serialize.SerializeTxs(txs)) {
			log.Printf("\x1b[32m  块校验通过  \x1b[0m \n")
		} else {
			log.Printf("\x1b[31m  块校验失败  \x1b[0m \n")
		}

		var validator models.Account
		// 第一块出块者问题
		if m.chain.blockChain[len(m.chain.blockChain)-1].Number == 0 {
			validator.Address = os.Getenv("validator")
		} else {
			validator = *db.ReadMostContributionAccount()
		}

		//块的创造者和UserTable中的用户一样时
		log.Println("下面验证块的发送者信息")
		if newBlock.Creator == validator.Address {
			log.Printf("\x1b[32m  发送者校验通过  \x1b[0m \n")
			//LogMsg:="OK  由 "+newBlock.From+" 发送的块 "+strconv.Itoa(newBlock.Number)+" 校验通过"
			//fmt.Fprintln(LogWriter, LogMsg)
			//LogWriter.Flush()
		} else if newBlock.Creator == validator.Address {
			log.Printf("\x1b[32m  发送者校验通过  \x1b[0m \n")
		} else {
			log.Printf("\x1b[31m  发送者校验失败  \x1b[0m \n")
		}

		log.Println("下面验证块的哈希值信息")
		if newBlock.PrevHash == m.chain.blockChain[len(m.chain.blockChain)-1].Hash {
			log.Printf("\x1b[32m  哈希值校验通过  \x1b[0m \n")
			//LogMsg:="序  "+newBlock.From+" 发送的块 "+strconv.Itoa(newBlock.Number)+"Hash校验通过"
			//fmt.Fprintln(LogWriter, LogMsg)
			//LogWriter.Flush()
		} else {
			log.Printf("\x1b[31m  哈希值校验失败  \x1b[0m \n")
		}

		m.chain.blockChain = append(m.chain.blockChain, *newBlock)

		// 更新数据
		_ = db.NewBlock(*newBlock)

		// 清除交易缓存
		clearTxCache(m.txCache, newBlock.Txs)

		return
	}
}

func (m *AQManager) CreateNewBlock(txs []*models.Transaction) models.Block {
	prevBlock := m.chain.blockChain[len(m.chain.blockChain)-1]
	return models.NewBlock(prevBlock, txs, m.account.Address)
}
