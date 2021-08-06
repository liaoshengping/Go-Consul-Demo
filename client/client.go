package main

import (
	"context"
	"fmt"
	"github.com/afex/hystrix-go/hystrix"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/consul/v2"
	opentracing2 "github.com/micro/go-plugins/wrapper/trace/opentracing/v2"
	"github.com/opentracing/opentracing-go"
	"github.com/syyongx/php2go"
	log "log"
	"net"
	"net/http"
	"order-micro/common"
	OrderService "order-micro/proto"
)

func main()  {
	consulRegister := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			"192.168.205.22:8500",
		}
	})

	//链路追踪
	t,io,err := common.NewTracer("order.clinet","192.168.205.22:6831")

	if err !=nil {
		log.Fatal(err)
	}

	defer io.Close()

	opentracing.SetGlobalTracer(t)


	//熔断
	hystrixStreamHander := hystrix.NewStreamHandler();

	hystrixStreamHander.Start()

	go func() {
		err = http.ListenAndServe(net.JoinHostPort("192.168.205.22","9096"),hystrixStreamHander)
		if err != nil {
			log.Fatal(err)
		}
	}()

	//创建一个新的服务
	server := micro.NewService(
		micro.Name("client"),
		micro.Registry(consulRegister),
		//绑定链路追踪
		micro.WrapClient(opentracing2.NewClientWrapper(opentracing.GlobalTracer())),



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