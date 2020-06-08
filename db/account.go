package db

import (
	"AQChain/models"
	"github.com/astaxie/beego/orm"
	"log"
)

func ReadAccount(address string) *models.Account {
	account := &models.Account{
		Address: address,
	}
	err := orm.NewOrm().Read(account)
	if err != nil {
		if err != orm.ErrNoRows {
			log.Println(err)
		}
		return nil
	}
	return account
}

func readAccount(db orm.Ormer, address string) *models.Account {
	account := &models.Account{
		Address: address,
	}
	err := db.Read(account)
	if err != nil {
		if err != orm.ErrNoRows {
			log.Println(err)
		}
		return nil
	}
	return account
}

// 这里的查询方法有疑问 排序是否随机 是否需要order 相同贡献值怎么办？
func ReadMostContributionAccount() *models.Account {
	account := &models.Account{}
	err := orm.NewOrm().Raw("select * from account where contribution = (select max(contribution) from account) ").QueryRow(account)
	if err != nil {
		log.Println(err)
		return nil
	}
	return account
}

// 保存新账户
func newAccount(db orm.Ormer, address string) *models.Account {
	account := &models.Account{
		Address:          address,
		Contribution:     0,
		ContributionFile: 0,
		ContributionTx:   0,
	}
	_, err := db.Insert(account)
	if err != nil {
		log.Println(err)
		return nil
	}
	return account
}

// 更新贡献值
func setContribution(db orm.Ormer, account models.Account) error {
	_, err := db.Update(&account, "contribution", "contribution_file", "contribution_tx")
	if err != nil {
		return err
	}
	return nil
}

// 添加贡献值
func addContribution(db orm.Ormer, address string, contributionFile, contributionTx float64) error {
	account := models.Account{
		Address: address,
	}
	err := db.Read(&account)
	if err != nil {
		if err != orm.ErrNoRows {
			log.Println(err)
		}
		return err
	}
	if contributionFile > 0 {
		account.ContributionFile += contributionFile
	}
	if contributionTx > 0 {
		account.ContributionTx += contributionTx
	}
	account.Contribution = account.ContributionTx + account.ContributionFile

	err = setContribution(db, account)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

// 减少贡献值 现在是清零
func subContribution(db orm.Ormer, address string, contributionFile, contributionTx float64) error {
	account := models.Account{
		Address: address,
	}
	err := db.Read(&account)
	if err != nil {
		log.Println(err)
		return err
	}
	account.ContributionFile = 0
	account.ContributionTx = 0
	account.Contribution = 0
	err = setContribution(db, account)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
