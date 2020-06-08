package client

func TimeTest() {
	/*if len(Blockchain)==1{
		fmt.Println("当前还没生成块。")
	}else {*/

	/*time.Sleep(time.Second * 10)
	for {
		a := time.Now().Format(TimeType)
		lastBlock := Blockchain[len(Blockchain)-1]
		b := lastBlock.Timestamp
		timeLayout := "2006-01-02 15:04:05"
		loc, _ := time.LoadLocation("Local")
		startUnix, _ := time.ParseInLocation(timeLayout, a, loc)
		endUnix, _ := time.ParseInLocation(timeLayout, b, loc)
		startTime := startUnix.Unix()
		endTime := endUnix.Unix()
		date := (startTime - endTime) / 60
		fmt.Print("当前的时间间隔为：")
		fmt.Println(date)
		if date > 0 && username == UserTable[maxCIndex].Address {
			//这里写逻辑
			//若在这段时间里面 ，没有生成新的交易，则直接生成一个空块
			var i int

			if len(txCache.Txs) == 0 {
				fmt.Println("在3分钟内没有交易也没有上传文件")
				//获取第一个用户
				for i = 0; i < len(NodeList); i++ {
					if NodeList[i].IsOnline == true {
						break
					}
				}
				//user:=NodeList[i].ID[0:8]
				//txCache := models.Transactions{}
				newblock := generateNewBlock(txCache.Txs)
				//BlockValidList <- SHA256Byte(SerializeTxs(txCache.Transactions))
				fmt.Println("正在生成新区块 ")
				a := SerializeBlock(newblock)
				b := "11block" + string(a)
				fmt.Println("向通道发送成块信息成功")
				P2PSend(b)
			}
			if len(txCache.Txs) == 1 {
				fmt.Println("在3分钟内块内有一笔交易")
				//user:=UserTable[maxCIndex].From[0:8]
				newblock := generateNewBlock(txCache.Txs)
				//BlockValidList <- SHA256Byte(SerializeTxs(txCache.Transactions))
				fmt.Println("正在生成新区块 ")
				a := SerializeBlock(newblock)
				b := "block11" + string(a)
				fmt.Println("向通道发送成块信息成功")
				P2PSend(b)
			}
			if len(txCache.Txs) == 2 {
				fmt.Println("在3分钟内块内有两笔交易")
				//user:=UserTable[maxCIndex].From[0:8]
				newblock := generateNewBlock(txCache.Txs)
				//BlockValidList <- SHA256Byte(SerializeTxs(txCache.Transactions))
				fmt.Println("正在生成新区块 ")
				a := SerializeBlock(newblock)
				b := "block11" + string(a)
				fmt.Println("向通道发送成块信息成功")
				P2PSend(b)
			}
		}
		time.Sleep(2 * time.Second)
	}*/

}
