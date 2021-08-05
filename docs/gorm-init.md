引入GORM并创建数据库

[GORM中文文档](https://learnku.com/docs/gorm/v2)

初始化 `pkg/model/model_init.go` 

```go
package model

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"order-micro/common"
	"order-micro/models/order"
	gormlogger "gorm.io/gorm/logger"
)

var Db *gorm.DB

func InitDatabase(config common.MysqlConfig) error {
	dsnString := config.User + ":" + config.Pwd + "@tcp(" + config.Host + ":" + config.Port + ")/"+config.Database+"?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsnString), &gorm.Config{
		Logger: gormlogger.Default.LogMode(gormlogger.Warn),
	})
	if err != nil {
		return err
	}
	Db = db

	Automigrate(*Db)

	return nil
}

func Automigrate(db gorm.DB)  {
	db.AutoMigrate(
		&order.Orders{},
	)
}

```

在 ``server/server.go``

添加

```
	MysqlConfig = common.GetMysqlFromConsul(consulConfig, "mysql")
	model.InitDatabase(*MysqlConfig)
```

初始化数据库，同时同步数据结构，这样就可以全局调用了

比如创建一个订单

```go
	_order := &order.Orders{}
	_order.OrderNo= "asdf"
	model.Db.Create(_order)
```
提交

```go
git add .
git commit -m 创建订单数据库
```