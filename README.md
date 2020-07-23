# golang-web

![](https://github.com/26huitailang/golang-web/workflows/Build%26Test/badge.svg)

## todo

- [ ] 使用cobra 作为命令入口
    - 移植mq的consumer和producer到cobra
- [ ] 整理文件组织
- [ ] 使用 traefik
- [ ] docker swarm

## docker

- 测试阿里云的触发

## profiling

使用三方的route，要注册`net/http/pprof`的路由，然后通过以下方式使用：

- 使用ab或wrk伪造持续的请求，以ab为例，持续120秒，每次发送10个请求：

    ab -t 120 -c 10 -m GET http://127.0.0.1:8080

- 在网站访问`http://127.0.0.1:8080/debug/pprof/profile`或者在命令行：

    go tool pprof http://127.0.0.1:8080/debug/pprof/profile

- 默认30s持续时间，可以用`?time=60`来控制，结束后浏览器访问会下载一个profile文件，命令行会直接进入交互模式。（下载的文件用`go tool pprof profile`来使用）

## sqlite3

使用sqlite的话，交叉编译要把CGO_ENABLED=1，因为依赖了C代码，但是又需要额外的gcc参数，`CC=arm-linux-gnueabi-gcc`，要额外安装，比较麻烦，可以研究使用docker编译。

## docker build

用了树莓派，为了更方便的build对应的平台，考虑用docker来build。

步骤：

- 准备一个镜像，安装了go和gcc-arm-linux-gnueabi等工具（记得apt clean）
  - [dockerfile参考](https://github.com/cloudfoundry-incubator/diego-dockerfiles/tree/master/golang-ci/Dockerfile)
  - go build -t golang:1.12 . ，构建新的镜像
- 将makefile调用的docker换到本地构建好的docker上，构建好的文件会出现在本地目录

```dockerfile
FROM golang:1.12

ADD . /go/src/github.com/26huitailang/golang-web
WORKDIR /go/src/github.com/26huitailang/golang-web

# arm build
RUN apt update
RUN apt install gcc-arm-linux-gnueabi -y
RUN apt clean
```

## Sup

Super simple deployment tool - think of it like 'make' for a network of servers https://pressly.github.io/sup

用来自动执行网络命令。
