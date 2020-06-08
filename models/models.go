package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
)

// 记录每轮的贡献值
type Contribution struct {
	ID               int     `orm:"pk;auto"`                    // 数据库自增id
	UserID           string  `orm:"column(user_id)"`            // 节点的ID
	Hash             string  `orm:"column(hash)"`               // 最后一轮的生成区块hash
	Contribution     float64 `orm:"size(300);column(con)"`      // ContributionFile+ContributionTx
	ContributionFile float64 `orm:"size(300);column(con_file)"` // 知识产权所得贡献值
	ContributionTx   float64 `orm:"size(300);column(con_tx)"`   // 交易所得贡献值
}

type Log struct {
	ID       int    `orm:"pk;auto"`          // 数据库自增id
	UserID   string `orm:"column(user_id)"`  // 节点的ID
	LogTime  string `orm:"column(log_time)"` // 记录时间
	Activity string `orm:"column(activity)"` // 活动的描述
}

func init() {
	orm.RegisterModel(new(Transaction))
	orm.RegisterModel(new(Block))
	orm.RegisterModel(new(Account))
	orm.RegisterModel(new(Contribution))
	orm.RegisterModel(new(Log))
	orm.RegisterModel(new(File))
	orm.RegisterModel(new(Transfer))
}
