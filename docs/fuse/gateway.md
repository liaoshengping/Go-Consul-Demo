Api 网关
============

>新建一个网关

```go
docker run --rm -v $(pwd):$(pwd) -w $(pwd) micro/micro:v2.9.3 new gateway
```

生成proto 文件 （这个我是本地环境）

```go
protoc ./gateway.proto --go_out=./ --micro_out=./
```


```go
docker run --rm -p 8088:8088 -v $(pwd):$(pwd) -w $(pwd)  micro/micro:v2.9.3   api --handler=api
```



