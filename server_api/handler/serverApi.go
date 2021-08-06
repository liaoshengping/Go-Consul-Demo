package handler

import (
	"context"
	"errors"
	"github.com/prometheus/common/log"
	ServerApiPkg "order-micro/server_api/proto"
	service "order-micro/services"
)

type ServerApi struct {
	orderServer service.OrderService
}

func (e *ServerApi) FindAll(ctx context.Context, req *ServerApiPkg.Request, rsp *ServerApiPkg.Response) error {
	log.Info("接受到访问请求")
	if _, ok := req.Get["goods_id"]; !ok {
		rsp.StatusCode = 500
		return errors.New("参数异常")
	}

	msg := e.orderServer.CreateOrder("12312")

	rsp.StatusCode = 200
	rsp.Body = msg

	return nil
}
