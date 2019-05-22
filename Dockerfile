FROM golang:1.12-alpine

ADD . /go/src/github.com/26huitailang/golang-web
WORKDIR /go/src/github.com/26huitailang/golang-web

# Go dep!
RUN go get -u github.com/golang/dep/...
RUN dep ensure

# build
RUN go install github.com/26huitailang/golang-web

ENTRYPOINT /go/bin/golang-web

EXPOSE 8080