![proto](https://img-blog.csdnimg.cn/f820d320f7cc40aaa9bb4bf2db4c1d57.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3FxXzIyODIzNTgx,size_16,color_FFFFFF,t_70)



编写 proto 我认为比较显著的优点有：

#### 1、性能好/效率高
时间开销： XML格式化（序列化）的开销还好；但是XML解析（反序列化）的开销就不敢恭维了。 但是protobuf在这个方面就进行了优化。可以使序列化和反序列化的时间开销都减短。、

#### 2、支持多种编程语言
可以生成 java ，go，py ，php 之类的代码，啥时候你不想用Go写RPC服务了，你可以用这个proto文件生成Java的接口，无缝对接。

这个系列我用订单做个微服务helloword

这个proto就像我们变成中的实现接口`（interface）`

# 开始编码
创建一个 `order-micro` 的项目

在proto 目录下创建一个 `order.proto` 文件

```proto
syntax = "proto3";
option go_package = "./;OrderService";
package OrderService;
service OrderService {
    rpc CreateOrder (Request) returns(Response){}
}
message Request {
    string goodsId = 1;
    string buyNum = 2;
}
message Response{
    int32  code = 1;
    string msg =2;
    int64 orderID =3;
}
```
第一行中 `syntax = "proto3";` 定义proto的版本，这边版本为3 ，protoc 就像是一个编程语言，也有自己的语法，每个版本可能有一些不同。

```
option go_package = "./;OrderService"; 
```

为生成的代码文件的目录，以及生成的包名，包名定义为 `OrderService`，在当前目录下生成。

下面的 `service` 就像是一个类，rpc 的方法 CreateOrder 请求和返回的参数有啥

goodsId  = 1 ，不是赋值 goodsId 就是为1 ，这个就像是一个标识符，你多谢几个参数，就累加就行了。

### 利用proto生成golang代码
现在比较多了用docker 生成proto 文件，其实是可以的哈，比如这个库
>znly/protoc

执行

```
 docker run --rm -u $(id -u) -v${PWD}:${PWD} -w${PWD} znly/protoc  -I$(pwd) ${PWD}/*.proto --go_out=${PWD} --micro_out=.
```

本人，执行会有点小小的问题，所以我就不用docker演示 了，但这个是可以的哈，可以[看文档](https://github.com/znly/docker-protobuf)

我本地是windows 就下载[windows版](https://blog.csdn.net/liupeifeng3514/article/details/78985575)

执行： 

```
protoc ./order.proto --go_out=./ --micro_out=./
```

可以 看到目录中已经出现了两个Go文件了

![在这里插入图片描述](https://img-blog.csdnimg.cn/ee872e3b547d46039cbaa75c5201cd19.png)


为了同学们可以更直观的体验，我把源码传到github ，大家可以看commit 查看教程相应的代码
```
git init
git add .
git commit -m "proto编写与生成"
```

源码
==============
[https://github.com/liaoshengping/Go-Consul-Demo](https://github.com/liaoshengping/Go-Consul-Demo)


如果本系列课程能给你带来帮助，可以点一下关注，后面带来更多分享
