![在这里插入图片描述](https://img-blog.csdnimg.cn/b86b723591b040439c62a1bb92c574bf.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3FxXzIyODIzNTgx,size_16,color_FFFFFF,t_70)

consoul 是微服务服务于发现,其他的不多说了

# Docker安装
>docker pull consul:latest

# 启动
>docker run -d -p 8500:8500 consul:latest

我们访问本地地址 端口8500 的就可以看到上图的效果了

下面我们需要做的是往 consul 注册服务

