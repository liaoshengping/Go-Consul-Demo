package main

import (
	"context"
	"fmt"
	hystrixGo "github.com/afex/hystrix-go/hystrix"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/web"
	"github.com/micro/go-plugins/registry/consul/v2"
	"github.com/micro/go-plugins/wrapper/breaker/hystrix/v2"
	opentracing2 "github.com/micro/go-plugins/wrapper/trace/opentracing/v2"
	"github.com/opentracing/opentracing-go"
	"github.com/syyongx/php2go"
	log "log"
	"order-micro/common"
	OrderService "order-micro/proto"
)

func main() {
	// 自定义全局默认超时时间和最大并发数
	hystrixGo.DefaultSleepWindow = 2
	hystrixGo.DefaultMaxConcurrent = 1

	// 针对指定服务接口使用不同熔断配置
	// 第一个参数name=服务名.接口.方法名，这并不是固定写法，而是因为官方plugin默认用这种方式拼接命令name
	// 之后我们自定义wrapper也同样使用了这种格式
	// 如果你采用了不同的name定义方式则以你的自定义格式为准
	hystrixGo.ConfigureCommand("go.micro.service.task.TaskService.Search",
		hystrixGo.CommandConfig{
			MaxConcurrentRequests: 10,
			Timeout:               1000,
		})


	consulRegister := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			"192.168.205.22:8500",
		}
	})

	//链路追踪
	t, io, err := common.NewTracer("order.clinet", "192.168.205.22:6831")
	if err != nil {
		log.Fatal(err)
	}
	defer io.Close()

	opentracing.SetGlobalTracer(t)


	//创建一个新的服务
	server := micro.NewService(
		micro.Name("client"),
		micro.Address(":8080"),
		micro.Registry(consulRegister),
		//绑定链路追踪
		micro.WrapClient(opentracing2.NewClientWrapper(opentracing.GlobalTracer())),
		//添加熔断
		micro.WrapClient(hystrix.NewClientWrapper()),
	)




//
r := gin.Default()

r.GET("/createOrder", func (c *gin.Context) {

	order := OrderService.NewOrderService("order.service", server.Client())

	response, err := order.CreateOrder(context.TODO(), &OrderService.Request{
		GoodsId: php2go.Md5("123"),
		BuyNum:  "1",
	})

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(response)

	c.JSON(200, gin.H{
		"message": response,
	})
})
//r.Run(":8080") // listen and serve on 0.0.0.0:8080
//	if err := server.Run(); err != nil {
//		log.Fatal(err)
//	}

	web_service := web.NewService(
		web.Name("go.micro.api"),
		web.Address(":8888"),
		web.Handler(r),
		web.Registry(consulRegister),
	)

	//初始化
	server.Init()
	if err := web_service.Run(); err != nil {
		log.Fatal(err)
	}
}









