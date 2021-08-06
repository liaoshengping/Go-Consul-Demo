package services

type OrderService struct {
}

//这个才是真正的服务类
func (h *OrderService) CreateOrder(goods_id string) string {
	return "服务处理了goods_id为" + goods_id
}
