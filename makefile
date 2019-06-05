# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOBENCHMARK=$(GOTEST) -bench .
GOGET=$(GOCMD) get
GOTOOL=$(GOCMD) tool
GODEP=dep
BINARY_NAME=main
BINARY_LINUX=$(BINARY_NAME)_linux
BINARY_ARM=$(BINARY_LINUX)_arm
PROFILE=profile

all: test build
build:
	$(GOBUILD) -o $(BINARY_NAME) -v
test:
	$(GOTEST) -v ./...
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_LINUX)
	rm -f $(BINARY_ARM)
run:
	$(GOBUILD) -o $(BINARY_NAME) -v
	./$(BINARY_NAME)
benchmark:
	$(GOBENCHMARK)

# dependence
deps:
	$(GODEP) ensure
deps-install:
	curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

# Cross compilation
build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_LINUX) -v
build-arm:
	CGO_ENABLED=0 GOOS=linux GOARCH=arm $(GOBUILD) -o $(BINARY_ARM) -v

# docker
docker-build: # 准备docker用于go build的环境
	docker build -t golang:1.12 .
docker-build-arm:
	docker run --rm -it -v "$(GOPATH)":/go -w /go/src/github.com/26huitailang/golang-web -e CGO_ENABLED=1 -e GOOS=linux -e GOARCH=arm -e CC=arm-linux-gnueabi-gcc golang:1.12 go build -o $(BINARY_ARM) -v

# documentation
doc:
	echo "access http://127.0.0.1:6060"
	godoc -http=:6060

# go tool
pprof:
	$(GOTOOL) pprof $(PROFILE)