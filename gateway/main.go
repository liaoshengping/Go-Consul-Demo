package main

import (
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2"
	"gateway/handler"
	"gateway/subscriber"

	gateway "gateway/proto/gateway"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.service.gateway"),
		micro.Version("latest"),
		micro.Address(":8080"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	gateway.RegisterGatewayHandler(service.Server(), new(handler.Gateway))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.service.gateway", service.Server(), new(subscriber.Gateway))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
