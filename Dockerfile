FROM golang:1.13-alpine3.12

#RUN sed -i "s@http://deb.debian.org@http://mirrors.aliyun.com@g" /etc/apt/sources.list && rm -rf /var/lib/apt/lists/* && apt-get update
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
RUN apk update
# arm dependency
#RUN apt install gcc-arm-linux-gnueabi -y
#RUN apt clean
RUN apk add --no-cache --virtual .build-deps \
 		gcc \
 		g++
#RUN apk clean

ENV GOPROXY https://goproxy.io,direct
ENV GO111MODULE on

WORKDIR /go/cache

COPY go.mod .
COPY go.sum .
RUN go mod download

WORKDIR /go/release
COPY . .
WORKDIR /go/release
RUN go generate -v ./...
RUN go build -o app ./cmd/gws/.
RUN pwd
RUN ls -lh

# alpine/scratch/busybox choose one
FROM alpine:3.12
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
RUN apk --no-cache add ca-certificates sqlite
RUN mkdir /data
RUN sqlite3 /data/test.db
WORKDIR /root
COPY --from=0 /go/release/app .
RUN ldd app
CMD ["./app", "server"]
