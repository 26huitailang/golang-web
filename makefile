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
COVERPROFILE=coverprofile
LOGFILE=main.log
REMOTEIP=192.168.8.217
REMOTEPATH=pi@$(REMOTEIP):/home/pi/go/golang-web/

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
	rm -f $(COVERPROFILE)
	rm -f $(LOGFILE)
run:
	$(GOBUILD) -o $(BINARY_NAME) -v
	./$(BINARY_NAME)
benchmark:
	$(GOBENCHMARK)
cover:
	$(GOTEST) -coverprofile $(COVERPROFILE) ./...
	$(GOTOOL) cover -html $(COVERPROFILE)

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

scp:
	scp config.json $(REMOTEPATH)
	scp supervisor-golang-web.conf $(REMOTEPATH)
	scp main_linux_arm $(REMOTEPATH)
	scp -r templates $(REMOTEPATH)

remote-stop:
	ssh pi@192.168.8.217 -p 22 "\
	sudo supervisorctl stop golang-web && \
	exit"
	echo done!

remote-start:
	ssh pi@192.168.8.217 -p 22 "\
	sudo supervisorctl start golang-web && \
	exit"
	echo done!

# deploy raspberry
deploy-raspberry: docker-build-arm remote-stop scp remote-start
	echo deploy done!