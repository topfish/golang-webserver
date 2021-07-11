# webapp
this is webserver by go

main文件是[webserver.go](https://github.com/topfish/webapp/blob/main/webserver.go)，可以直接编译

[webserver](https://github.com/topfish/webapp/blob/main/webserver)是编译好的，Linux环境的可执行二进制文件，可以直接执行，也可以打包进进行

[webserver.exe](https://github.com/topfish/webapp/blob/main/webserver.exe)同上，是Windows下的可执行文件

# 打包为镜像：
参考: [k8s](https://github.com/topfish/k8s)

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

- 运行以下命令用刚构建的镜像来启动容器：
-     $ docker run -d -p 80:80 --name myweb webserver:1.0    //启动的容器名称是myweb，容器的80端口映射到主机的80端口上。

## 验证：
- 可以在主机（或公网）请求本机IP：80
