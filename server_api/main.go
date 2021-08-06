package server_api

import (
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/consul/v2"
	opentracing2 "github.com/micro/go-plugins/wrapper/trace/opentracing/v2"
	"github.com/opentracing/opentracing-go"
	"log"
	"order-micro/common"
)

func main()  {
	consulRegister := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			"192.168.205.22:8500",
		}
	})

	//链路追踪
	t,io,err := common.NewTracer("order.api","192.168.205.22:6831")

	if err !=nil {
		log.Fatal(err)
	}

	defer io.Close()

	opentracing.SetGlobalTracer(t)

	//熔断器



	//创建一个新的服务
	server := micro.NewService(
		micro.Name("api.orderApi"),
		micro.Registry(consulRegister),
		//绑定链路追踪
		micro.WrapClient(opentracing2.NewClientWrapper(opentracing.GlobalTracer())),
	)



	//初始化
	server.Init()
}