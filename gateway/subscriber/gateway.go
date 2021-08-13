package subscriber

import (
	"context"
	log "github.com/micro/go-micro/v2/logger"

	gateway "gateway/proto/gateway"
)

type Gateway struct{}

func (e *Gateway) Handle(ctx context.Context, msg *gateway.Message) error {
	log.Info("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *gateway.Message) error {
	log.Info("Function Received message: ", msg.Say)
	return nil
}
