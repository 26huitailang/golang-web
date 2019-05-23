# golang-web

## requirements

- httprouter

## docker

- 测试阿里云的触发

## todo

- [ ] 添加sqlite，初始化的数据放入sqlite，支持一些数据库的排序操作
- [ ] sql使用Gorm来支持
- [ ] 考虑使用Echo

## profiling

使用三方的route，要注册`net/http/pprof`的路由，然后通过以下方式使用：

- 使用ab或wrk伪造持续的请求，以ab为例，持续120秒，每次发送10个请求：

    ab -t 120 -c 10 -m GET http://127.0.0.1:8080
- 在网站访问`http://127.0.0.1:8080/debug/pprof/profile`或者在命令行：

    go tool pprof http://127.0.0.1:8080/debug/pprof/profile

- 默认30s持续时间，可以用`?time=60`来控制，结束后浏览器访问会下载一个profile文件，命令行会直接进入交互模式。（下载的文件用`go tool pprof profile`来使用）