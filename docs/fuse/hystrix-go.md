熔断器

==================

go get github.com/micro/go-plugins/wrapper/breaker/hystrix/v2

###hystrix-go  

>[hɪst'rɪks]（海丝吹克丝）

#### 作用：

* 1.阻止故障的连锁反应
* 2.快速失败并迅速恢复
* 3.回退并优雅降级
* 4.提供近实时的监控与告警

实质就是解决微服务带来的雪崩的隐患

#### 使用原则：

* 1.防止单独资源依赖耗尽资源
* 2.过载立即切断并迅速失败，防止排队
* 3.回退机制
* 4.警报，监控

#### 原理

* 计数上报

#### hystrix-go 状态

* CLOSED 关闭
* Open 开启
* HALF_OPEN 半开，允许某些流量通过，如果超时再进入打开的状态

打开意味着熔断了



>go get github.com/micro/go-plugins/wrapper/breaker/hystrix/v2


参考：[断路器Hystrix](https://blog.csdn.net/hotcoffie/article/details/108844945)


直接添加服务即可

```go
	//创建一个新的服务
	server := micro.NewService(
		micro.Name("client"),
		micro.Address(":8080"),
		micro.Registry(consulRegister),
		//绑定链路追踪
		micro.WrapClient(opentracing2.NewClientWrapper(opentracing.GlobalTracer())),
		//添加熔断
		micro.WrapClient(hystrix.NewClientWrapper()),
	)
```

熔断还是挺有意思的，默认1秒 如果调用的服务端响应时间超过了1秒，则会断开

服务端 ``createOrder``添加了sleep 3秒

```go
time.Sleep(3 * time.Second)
```

```go
var count int = 0;



//其实这个是handler
func (h *OrderHandler) CreateOrder(ctx context.Context, req *OrderService.Request, rp *OrderService.Response) error {

	time.Sleep(3 * time.Second)
	count++
	fmt.Println("请求到了"+strconv.Itoa(count));

}
```

当我在client 端请求的时候，发现服务返回的接口是为null了

```go
{
  "message": null
}
```

也就是已经被熔断了，没等到服务端返回值，就已经被截断并且返回了

####问题来了
client端到底有没有触发这个服务呢，答案是肯定的，请求到了服务端，服务端超时之后hystrix再做判断的。

那是否每次都会请求呢？ 每次都做超时处理？

我们在刚刚createOrder 的方法中打印请求到了的响应，并做一个实验，我快速的用浏览器请求10次，看下打印出来是否是我们猜想的？

```go
2021-08-13 17:35:42  file=grpc/grpc.go:697 level=info Registry [consul] Registering node: order.service-05d329fb-ce93-4bde-a2f7-15e8ff383bd4
请求到了1
请求到了2
请求到了3
请求到了4
```
可以看到，请求到了4次，这个只是我手动f5 刷新浏览器的，hystrix 自己会计算是否到了下次触发server的时间。
