package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/consul/v2"
	"github.com/syyongx/php2go"
	OrderService "order-micro/proto"
)

func main()  {
	consulRegister := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			"192.168.205.22:8500",
		}
	})

	//创建一个新的服务
	server := micro.NewService(
		micro.Name("client"),
		micro.Registry(consulRegister),
	)
	//初始化
	server.Init()
	//
	r := gin.Default()

	r.GET("/createOrder", func(c *gin.Context) {

		order := OrderService.NewOrderService("order.service", server.Client())

		response, err := order.CreateOrder(context.TODO(), &OrderService.Request{
			GoodsId: php2go.Md5("123"),
			BuyNum:"1",
		})
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(response)

		c.JSON(200, gin.H{
			"message":response,
		})
	})
	r.Run(":8080") // listen and serve on 0.0.0.0:8080

}