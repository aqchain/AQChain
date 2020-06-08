package models

type Account struct {
	Address          string  `orm:"column(address);pk"`        // 账户名
	Contribution     float64 `orm:"column(contribution)"`      // ContributionFile+ContributionTx
	ContributionFile float64 `orm:"column(contribution_file)"` // 知识产权所得贡献值
	ContributionTx   float64 `orm:"column(contribution_tx)"`   // 交易所得贡献值
}
