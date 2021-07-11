# webapp
this is webserver by go

main文件是[webserver.go](https://github.com/topfish/webapp/blob/main/webserver.go)，可以直接编译

[webserver](https://github.com/topfish/webapp/blob/main/webserver)是编译好的，Linux环境的可执行二进制文件，可以直接执行，也可以打包为镜像

[webserver.exe](https://github.com/topfish/webapp/blob/main/webserver.exe)同上，是Windows下的可执行文件

# 容器化操作：
参考: [k8s](https://github.com/topfish/k8s)

如果没有安装docker环境的，需要先安装docker。以centos为例，参考：
[Install Docker Engine on CentOS](https://docs.docker.com/engine/install/centos/)

### Dockerfile
    FROM centos
    COPY webserver /home/webserver
    RUN ln -f /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
    CMD /home/webserver

### Dockerfile说明
- 本dockerfile是一个简单的打包可运行程序的事例。

- 比如，把编译完成的二进制可执行文件命名为【webserver】，放在Dockerfile的同目录下。此步骤可以使用[topfish/webapp](https://github.com/topfish/webapp)的代码作为测试。

- 然后运行以下命令来构建镜像:
-     $ docker build -t webserver:1.0 .     //此处镜像名称设置为 webserver:1.0。注意行尾的【.】

## 验证：
### 直接 docker run 方式
- 运行以下命令用刚构建的镜像来启动容器：
-     $ docker run -d -p 80:80 --name myweb webserver:1.0    //启动的容器名称是myweb，容器的80端口映射到主机的80端口上。
- 可以在主机（或公网）请求本机IP：80
### 在k8s上运行
- 需要先把本地的镜像推送到公网的仓库，比如，我使用了阿里云的镜像服务提供的仓库，镜像地址：
-       registry.cn-hangzhou.aliyuncs.com/fishgo/webserver:1.0
- 用镜像部署服务并暴露到公网，可参考下面的yaml文件
-  ps:这个事例用了LoadBalancer的方式，会自动创建一个公网类型的SLB用于暴露服务到公网
- [webapp.yaml](https://github.com/topfish/webapp/blob/main/webapp.yaml)
