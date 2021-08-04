package model

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"order-micro/common"
	"order-micro/model/order"
)

var db *gorm.DB

func InitDatabase(config common.MysqlConfig) error {
	dsnString := config.User + ":" + config.Pwd + "@tcp(" + config.Host + ":" + config.Port + ")/"+config.Database+"?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsnString), &gorm.Config{})
	if err != nil {
		return err
	}
	db = db

	Automigrate(*db)

	return nil
}

func Automigrate(db gorm.DB)  {
	db.AutoMigrate(
		&order.Orders{},
	)
}

