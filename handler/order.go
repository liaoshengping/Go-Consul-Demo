package handler

import (
	"context"
	"fmt"
	"github.com/syyongx/php2go"
	"order-micro/models/order"
	"order-micro/pkg/model"
	OrderService "order-micro/proto"
)

var count int = 0;

type OrderHandler struct{}

//其实这个是handler
func (h *OrderHandler) CreateOrder(ctx context.Context, req *OrderService.Request, rp *OrderService.Response) error {

	//time.Sleep(3 * time.Second)
	//
	//count++
	//fmt.Println("请求到了"+strconv.Itoa(count));


	defer func() {
		if err := recover(); err != nil {
			return
		}
	}()

	goodsId := req.GoodsId

	rp.Code = 200

	_order := &order.Orders{}
	_order.OrderNo = "asdf"
	model.Db.Create(_order)

	generateOrderId := php2go.Rand(0, 121212) //github.com/syyongx/php2go
	generateOrderId_ := int64(generateOrderId)
	rp.OrderID = generateOrderId_

	rp.Msg = fmt.Sprintf("提交订单的goodsid为%s生成的订单id为%d", goodsId, generateOrderId)
	//创建订单的逻辑

	return nil
}
