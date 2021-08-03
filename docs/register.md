![在这里插入图片描述](https://img-blog.csdnimg.cn/f1789e4ad7c247b8b8cf1b7c9c2f5998.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3FxXzIyODIzNTgx,size_16,color_FFFFFF,t_70)


> 上节把服务端讲完了，现在我们需要把服务注册到consul中去。

```shell
go get github.com/micro/go-micro/v2/registry
go get github.com/micro/go-plugins/registry/consul/v2
```
引入微服务注册包和consul 注册包

我们在server.go 中修改代码：

```go
	consulRegister := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			"192.168.205.22:8500",
		}
	})
	service := micro.NewService(
		micro.Name("order.service"),
		micro.Registry(consulRegister),
	)
```
这边的`192.168.205.22` 是我本地虚拟机的ip，同学们可以根据自己实际情况修改，如果是本地的一般是172.0.0.1

`8500` 是第二节中的的consul 的端口，根据实际情况修改哈

在NewService服务中，把consulRegister 信息注册到服务当中去。

这样就完成了服务的注册了，我们打开cousul 查看一下，这个order.service 是否在服务列表中
![在这里插入图片描述](https://img-blog.csdnimg.cn/63c471983dd94b1aa0c7d319a6e16164.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3FxXzIyODIzNTgx,size_16,color_FFFFFF,t_70)
嗯，果然，非常好！
![在这里插入图片描述](https://img-blog.csdnimg.cn/img_convert/9479e46281f6d16f68c85e73c6ca64d0.gif)
接下来就是调用注册再consul 中的服务了

```shell
git add .
git commit -m "consul 服务注册"
```

