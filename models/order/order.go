package order

const(
	OrderStatusCreated= 1
	OrderStatusPaid = 2
	OrderStatusFinish = 3
	OrderStatusClose = 4

	PayStatusNot = 0
	PayStatusSuccess = 1
	PayStatusRefund =2

)

type Orders struct {
	OrderNo string `json:"order_no"`
	OrderStatus int8 `json:"order_status"`
	PayStatus int8 	`json:"pay_status"`
}
