package db

import (
	"AQChain/models"
	"AQChain/utils"
	"github.com/astaxie/beego/orm"
	"log"
)

//初始化数据库
func InitDB(account string) {
	path := utils.CaculatePath() + account + ".db"
	// 注册一个数据库驱动程序使用指定的驱动程序名称，这可以定义驱动程序是哪种数据库类型。
	orm.RegisterDriver("sqlite", orm.DRSqlite)
	//设置数据库连接参数。
	orm.RegisterDataBase("default", "sqlite3", path)
	//运行syncdb命令行名称意味着表的别名默认为“default”
	orm.RunSyncdb("default", true, false)
}

func InsertContribution(c []*models.Contribution) {
	o := orm.NewOrm()
	for i := 0; i < len(c); i++ {
		record := &models.Contribution{UserID: c[i].UserID, Hash: c[i].Hash}
		err := o.Read(record)
		if err != nil {
			_, err = o.Insert(c[i])
			if err != nil {
				log.Println(err)
			}
		} else {
			record.Contribution = c[i].Contribution
			record.ContributionFile = c[i].ContributionFile
			record.ContributionTx = c[i].ContributionTx
			_, err = o.Update(record, "Hash", "Contribution", "ContributionFile", "ContributionTx")
			if err != nil {
				log.Println(err)
			}
		}
	}
}
