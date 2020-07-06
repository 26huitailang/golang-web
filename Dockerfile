FROM golang:1.13-stretch

WORKDIR /go/src

# arm build
RUN sed -i "s@http://deb.debian.org@http://mirrors.aliyun.com@g" /etc/apt/sources.list && rm -rf /var/lib/apt/lists/* && apt-get update
RUN apt install gcc-arm-linux-gnueabi -y
RUN apt clean

COPY go.mod .
COPY go.sum .
RUN go mod download
RUN go mod vendor
RUN go env -w GOPROXY=https://goproxy.io,direct
RUN go get github.com/GeertJohan/go.rice/rice

# Go dep!
#RUN go get -u github.com/golang/dep/...
#RUN dep ensure

# build
# RUN go install github.com/26huitailang/golang-web

# ENTRYPOINT /go/bin/golang-web

# EXPOSE 8080