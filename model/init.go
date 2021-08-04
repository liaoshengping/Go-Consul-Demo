package model

import (
	"gorm.io/gorm"
	"order-micro/common"
	"gorm.io/driver/mysql"
)

var DB *gorm.DB

func InitDatabase(config common.MysqlConfig) error {
	dsnString := config.User + ":" + config.Pwd + "@tcp(" + config.Host + ":" + config.Port + ")/goblog?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsnString), &gorm.Config{})
	if err != nil {
		return err
	}
	DB = db
	return nil
}
