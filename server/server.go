package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/consul/v2"
	"github.com/syyongx/php2go"
	OrderService "order-micro/proto"
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

	rp.Msg = fmt.Sprintf("提交订单的goodsid为%s生成的订单id为%d", goodsId, generateOrderId)

	return nil
}

func main() {
	consulRegister := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			"192.168.205.22:8500",
		}
	})
	service := micro.NewService(
		micro.Name("order.service"),
		micro.Registry(consulRegister),
	)

	service.Init()

	OrderService.RegisterOrderServiceHandler(service.Server(), new(OrderServer))

	if err := service.Run(); err != nil {
		fmt.Println("service err", err)
	}
}