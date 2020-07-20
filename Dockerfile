FROM golang:1.13-stretch

WORKDIR /go/src

RUN sed -i "s@http://deb.debian.org@http://mirrors.aliyun.com@g" /etc/apt/sources.list && rm -rf /var/lib/apt/lists/* && apt-get update
# arm dependency
RUN apt install gcc-arm-linux-gnueabi -y
RUN apt clean

COPY go.mod .
COPY go.sum .
RUN go env -w GOPROXY=https://goproxy.io,direct
RUN go mod download
RUN go mod vendor
RUN go get github.com/GeertJohan/go.rice/rice
COPY . .
WORKDIR /go/src/server
RUN go generate
RUN go build -o app

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=0 /go/src/server/app .
CMD ["./app"]
