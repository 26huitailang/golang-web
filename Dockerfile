FROM golang:1.12

ADD . /go/src/github.com/26huitailang/golang-web
WORKDIR /go/src/github.com/26huitailang/golang-web

# arm build
RUN apt update
RUN apt install gcc-arm-linux-gnueabi -y
RUN apt clean

# Go dep!
#RUN go get -u github.com/golang/dep/...
#RUN dep ensure

# build
# RUN go install github.com/26huitailang/golang-web

# ENTRYPOINT /go/bin/golang-web

# EXPOSE 8080