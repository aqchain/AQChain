package db

import (
	"AQChain/models"
	"github.com/astaxie/beego/orm"
	"log"
)

func ReadFile(fileID string) *models.File {
	file := &models.File{
		ID: fileID,
	}
	err := orm.NewOrm().Read(file)
	if err != nil {
		log.Println(err)
		return nil
	}
	return file
}

func ReadFiles() []*models.File {
	var files []*models.File
	_, err := orm.NewOrm().QueryTable(&models.File{}).All(&files)
	if err != nil {
		log.Println(err)
		return nil
	}
	return files
}

func ReadFilesByOwner(owner string) []*models.File {
	var files []*models.File
	_, err := orm.NewOrm().QueryTable(&models.File{}).Filter("owner", owner).All(&files)
	if err != nil {
		log.Println(err)
		return nil
	}
	return files
}

func insertFile(o orm.Ormer, fileID, creator, createTime string) error {
	file := models.File{
		ID:         fileID,
		Creator:    creator,
		CreateTime: createTime,
		Owner:      creator,
	}
	_, err := o.Insert(&file)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func updateFile(o orm.Ormer, fileID, owner, transferID string, value float64) error {
	file := models.File{
		ID: fileID,
	}
	err := o.Read(&file)
	if err != nil {
		if err != orm.ErrNoRows {
			log.Println(err)
		}
		return err
	}
	file.Owner = owner
	file.Value = value
	file.TransferRecord += "/" + transferID
	_, err = o.Update(&file, "owner", "value", "transfer_record")
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
