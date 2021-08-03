
先初始化下编写main文件

```php
	consulRegister := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			"192.168.205.22:8500",
		}
	})

	//创建一个新的服务
	server := micro.NewService(
		micro.Name("client"),
		micro.Registry(consulRegister),
	)
	//初始化
	server.Init()
```

以上和服务端长得几乎是一模一样，`micro.Registry(consulRegister)`,这个参数能调用consul 中的服务。

为了更贴近真实的案例，我们用Gin 框架做一个🌰栗子
```php
go get github.com/gin-gonic/gin
```
# 完整代码

```php
package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/consul/v2"
	"github.com/syyongx/php2go"
	OrderService "order-micro/proto"
)

func main()  {

	consulRegister := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			"192.168.205.22:8500",
		}
	})

	//创建一个新的服务
	server := micro.NewService(
		micro.Name("client"),
		micro.Registry(consulRegister),
	)
	//初始化
	server.Init()
	//
	r := gin.Default()

	r.GET("/createOrder", func(c *gin.Context) {

		order := OrderService.NewOrderService("order.service", server.Client())

		response, err := order.CreateOrder(context.TODO(), &OrderService.Request{
			GoodsId: php2go.Md5("123"),
			BuyNum:"1",
		})
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(response)

		c.JSON(200, gin.H{
			"message":response,
		})
	})
	r.Run(":8080") // listen and serve on 0.0.0.0:8080

}
```

其中

```go

order := OrderService.NewOrderService("order.service", server.Client())

response, err := order.CreateOrder(context.TODO(), &OrderService.Request{
	GoodsId: php2go.Md5("123"),
	BuyNum:"1",
})
if err != nil {
	fmt.Println(err)
}
fmt.Println(response)
```

这块代码调用了服务端 `order.service`


Go Go Go ！！！跑起来看看

![在这里插入图片描述](http://p1.itc.cn/images01/20201207/52a9bf649675424c9c10481791404b2e.gif)

访问

```shell
http://127.0.0.1:8080/createOrder
```
![在这里插入图片描述](https://img-blog.csdnimg.cn/32061f797d0c4ac780441e5c39cb1b9f.png)

完美，成功请求到了服务端。

```
git add .
git commit -m "客户端编写"
```



