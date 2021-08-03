![在这里插入图片描述](https://img-blog.csdnimg.cn/eb144e13347c474faa5af88c75d00c41.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3FxXzIyODIzNTgx,size_16,color_FFFFFF,t_70)


上一节，完成了consul的搭建，这节的目标是 写一个服务类，用于客户端调用。

## 编写服务端
创建 `/server/server.go`

```go
type OrderServer struct {}

func (h *OrderServer)CreateOrder(ctx context.Context,req *OrderService.Request, rp *OrderService.Response) error  {
	defer func() {
		if err :=recover(); err != nil{
			return
		}
	}()

	goodsId := req.GoodsId;

	rp.Code = 200

	generateOrderId := php2go.Rand(0, 121212); //github.com/syyongx/php2go
	//创建订单的逻辑
	generateOrderId_ := int64(generateOrderId);

	rp.OrderID = generateOrderId_;
	rp.Msg  = fmt.Sprintf("提交订单的goodsid为%s",goodsId)
	return nil
}
```
编写main

>go get github.com/micro/go-micro/v2

获取微服务

server.go 的完整代码：

```go
package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/v2"
	"github.com/syyongx/php2go"
	OrderService "order-micro/proto"
)

type OrderServer struct{}

func (h *OrderServer) CreateOrder(ctx context.Context, req *OrderService.Request, rp *OrderService.Response) error {
	defer func() {
		if err := recover(); err != nil {
			return
		}
	}()

	goodsId := req.GoodsId

	rp.Code = 200

	generateOrderId := php2go.Rand(0, 121212) //github.com/syyongx/php2go
	//创建订单的逻辑
	generateOrderId_ := int64(generateOrderId)

	rp.OrderID = generateOrderId_

	rp.Msg = fmt.Sprintf("提交订单的goodsid为%s生成的订单id为%d", goodsId, generateOrderId)

	return nil
}

func main() {

	service := micro.NewService(
		micro.Name("order.service"),
	)

	service.Init()

	OrderService.RegisterOrderServiceHandler(service.Server(), new(OrderServer))

	if err := service.Run(); err != nil {
		fmt.Println("service err", err)
	}
}

```

## 解释一下

```go
	service := micro.NewService(
		micro.Name("order.service"),
	)
```
这边是实现一个微服务，在NewService 可以加一些Options ，比如Name ，这边的Name 是待会注册到Consul 的名字，也是客户端调用的名字
也可以添加自定义的端口号比如：

```shell
micro.Address(":8001")
```
接下来是服务的初始化
```go
	service.Init()

	OrderService.RegisterOrderServiceHandler(service.Server(), new(OrderServer))
```

并把刚刚的结构体里面的OrderServer 的结构注册到这个服务中去，当然也会把CreateOrder 注册进去，这样客户端就可以调用了方法了

和很多框架比如Gin一样，Run启动服务。
```go
	if err := service.Run(); err != nil {
		fmt.Println("service err", err)
	}
```

>go run server.go

```shell
E:\linuxdir\go\src\order-micro\server>go run server.go
2021-08-02 16:28:06  file=v2@v2.9.1/service.go:200 level=info Starting [service] order.service
2021-08-02 16:28:06  file=grpc/grpc.go:864 level=info Server [grpc] Listening on [::]:51567
2021-08-02 16:28:06  file=grpc/grpc.go:697 level=info Registry [mdns] Registering node: order.service-f59428a9-24d9-4ea2-967b-8fdb9351327d
```
你将看到以上信息，说你启动了一个 order.service 的服务，并且在监听了，监听端口`51567`是系统随机的，刚刚说了，也可以自定义



```shell
git add .
git commit -m "编写服务端并启动"
```
