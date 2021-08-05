熔断器

==================

github：https://github.com/afex/hystrix-go

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

### 安装

>docker pull cap1573/hystrix-dashboard

```
docker run -d -p 9002:9002 cap1573/hystrix-dashboard
```







