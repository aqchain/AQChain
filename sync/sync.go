package sync

import (
	"AQChain/db"
	"AQChain/models"
	"fmt"
	"strings"
)

func Start() {
	lastBlock := db.ReadLastBlock()
	bc := make([]models.Block, 0)
	// 计数器 取 循环次数过多 直接错误中断
	var count = 0
	// 同步区块 单开协程的前提 是要设置好通道传递 终止协程
	for {
		// 向一个节点发送请求 !没有做到循环中到不断向下一个节点发送
		P2PSendOne("GBH" + string(SerializeBlock(lastBlock)))
		count++
		if count > len(NodeList) {
			panic("过多的请求错误")
		}
		str := P2PRecv()
		// 处理接收区块请求
		if strings.Contains(str, "BH") {
			// 解析对方发来的区块
			s := str
			s = strings.Trim(s, "BH")
			bc = DeserializeBlockChain([]byte(s))
			fmt.Println("收到回复", len(bc))
			if len(bc) == 1 && lastBlock.Number == bc[0].Number {
				fmt.Println("区块已经完全同步")
				break
			}
			fmt.Println("收到区块")
			// 检查一下
			if lastBlock.Number == bc[0].Number && lastBlock.Hash == bc[0].Hash {

			}
			// 保存获取的链 接到自己的链
			for i := 0; i < len(bc); i++ {
				// 数据库得开事务 之后改 保存Block
				db.InsertBlockDB(bc[i])
				// 保存块中merkleNode
				for j := 0; j < len(bc[i].Transactions.Txs); j++ {
					db.InsertTransaction(bc[i].Transactions.Txs[j])
				}
				Blockchain = append(Blockchain, bc[i])
			}
			for i := 0; i < len(Blockchain); i++ {
				fmt.Println("收到区块", Blockchain[i].Hash)
				fmt.Println("收到区块", Blockchain[i].Number)
			}

			// 修改最后一块位置 继续循环 再次发送同步请求 为了确保完全同步
			// 循环的出口 在len(bc) == 1 即对方返回了和自己相同的最后区块
			lastBlock = Blockchain[len(Blockchain)-1]
		}
	}

	// 发送同步交易请求 要对方缓冲里的所有交易 如果可以做到两点之间通信是最好的 就不需要发那么多没有用的消息
	P2PSendOne("GTX" + id)
}
