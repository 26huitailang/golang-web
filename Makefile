# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOBENCHMARK=$(GOTEST) -bench .
GOGET=$(GOCMD) get
GOTOOL=$(GOCMD) tool
BINARY_NAME=main
BINARY_LINUX=$(BINARY_NAME)_linux
BINARY_ARM=$(BINARY_LINUX)_arm
PROFILE=profile
COVERPROFILE=coverprofile
COVERAGE_FILE=coverage.html
LOGFILE=main.log
REMOTEIP=pi
REMOTEPATH=pi@$(REMOTEIP):/home/pi/go/golang-web/
DOCKERTAG=golang:1.13-stretch

generate:
	$(GOCMD) generate ./...
all: generate test build
build: generate
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
	$(GOTOOL) cover -html $(COVERPROFILE) -o $(COVERAGE_FILE)


# Cross compilation
build-linux: generate
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_LINUX) -v
build-arm: generate
	$(GOBUILD) -o $(BINARY_ARM) -v ./server/...

# docker 不要升级到1.14 依赖GLIBC 2.28 树莓派没有，可以试试stretch版本
# docker run --rm -it golang:1.14 ldd --version
docker-build: # 准备docker用于go build的环境
	docker build -f Dockerfile.cross -t $(DOCKERTAG) .
docker-build-arm:
	docker run --rm -it -v "$(PWD)":/go/release -w /go/release -e CGO_ENABLED=1 -e GOOS=linux -e GOARCH=arm -e CC=arm-linux-gnueabi-gcc $(DOCKERTAG) make build-arm

# documentation
doc:
	echo "access http://127.0.0.1:6060"
	godoc -http=:6060

# go tool
pprof:
	$(GOTOOL) pprof $(PROFILE)

scp:
	scp supervisor-golang-web.conf $(REMOTEPATH)
	scp main_linux_arm $(REMOTEPATH)

remote-stop:
	ssh pi@pi -p 22 "\
	sudo supervisorctl stop golang-web && \
	exit"
	echo done!

remote-start:
	ssh pi@pi -p 22 "\
	sudo supervisorctl start golang-web && \
	exit"
	echo done!

# deploy raspberry
deploy-raspberry: docker-build-arm remote-stop scp remote-start
	echo deploy done!