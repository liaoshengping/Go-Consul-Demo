package main

import (
	"fmt"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/consul/v2"
	ratelimit "github.com/micro/go-plugins/wrapper/ratelimiter/uber/v2"
	opentracing2 "github.com/micro/go-plugins/wrapper/trace/opentracing/v2"
	"github.com/opentracing/opentracing-go"
	"log"
	"order-micro/common"
	"order-micro/handler"
	model "order-micro/pkg/model"
	OrderService "order-micro/proto"
)

var (
	MysqlConfig *common.MysqlConfig
	QPS = 100
)



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
	//链路追踪
	t,io,err := common.NewTracer("order.service","192.168.205.22:6831")

	if err !=nil {
		log.Fatal(err)
	}

	defer io.Close()

	opentracing.SetGlobalTracer(t)


	service := micro.NewService(
		micro.Name("order.service"),
		//注册中心
		micro.Registry(consulRegister),
		//绑定链路追踪
		micro.WrapHandler(opentracing2.NewHandlerWrapper(opentracing.GlobalTracer())),
		//限流
		micro.WrapHandler(ratelimit.NewHandlerWrapper(QPS)),

	)
	//获取配置中心数据
	MysqlConfig = common.GetMysqlFromConsul(consulConfig, "mysql")
	model.InitDatabase(*MysqlConfig)
	service.Init()
	OrderService.RegisterOrderServiceHandler(service.Server(), new(handler.OrderHandler))

	if err := service.Run(); err != nil {
		fmt.Println("service err", err)
	}
}
