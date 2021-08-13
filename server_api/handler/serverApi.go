package handler

import (
	"context"
	"errors"
	"github.com/prometheus/common/log"
	OrderService "order-micro/proto"
	ServerApiPkg "order-micro/server_api/proto"
)

type ServerApi struct {
	orderServer OrderService.OrderService
}

func (e *ServerApi) FindAll(ctx context.Context, req *ServerApiPkg.Request, rsp *ServerApiPkg.Response) error {

	log.Info("接受到访问请求")
	if _, ok := req.Get["goods_id"]; !ok {
		rsp.StatusCode = 500
		return errors.New("参数异常")
	}
	rsp.StatusCode = 200
	rsp.Body = "kkkk"

	return nil
}
