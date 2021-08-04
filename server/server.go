package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/consul/v2"
	"github.com/syyongx/php2go"
	"order-micro/common"
	"order-micro/models/order"
	model2 "order-micro/pkg/model"
	OrderService "order-micro/proto"
)

var (
	MysqlConfig *common.MysqlConfig
)

type OrderServer struct{}

func (h *OrderServer) CreateOrder(ctx context.Context, req *OrderService.Request, rp *OrderService.Response) error {
	defer func() {
		if err := recover(); err != nil {
			return
		}
	}()

	goodsId := req.GoodsId

	rp.Code = 200





	generateOrderId := php2go.Rand(0, 121212) //github.com/syyongx/php2go
	//创建订单的逻辑
	generateOrderId_ := int64(generateOrderId)

	rp.OrderID = generateOrderId_

	rp.Msg = fmt.Sprintf("提交订单的goodsid为%s生成的订单id为%d 数据库host：%s", goodsId, generateOrderId, MysqlConfig.Host)

	return nil
}

func main() {
	//配置中心
	consulConfig, err := common.GetConsulConfig("192.168.205.22", 8500, "")

	if err != nil {
		fmt.Println(err)
		return
	}

	consulRegister := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			"192.168.205.22:8500",
		}
	})
	service := micro.NewService(
		micro.Name("order.service"),
		micro.Registry(consulRegister),
	)
	//获取配置中心数据
	MysqlConfig = common.GetMysqlFromConsul(consulConfig, "mysql")

	model2.InitDatabase(*MysqlConfig)

	_order := order.Orders{}
	_order.OrderNo= "kskdjf"
	err = model2.Db.Create(_order).Error

	if err != nil {
		fmt.Println(err)
	}

	service.Init()

	OrderService.RegisterOrderServiceHandler(service.Server(), new(OrderServer))

	if err := service.Run(); err != nil {
		fmt.Println("service err", err)
	}
}
