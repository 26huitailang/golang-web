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
deps:
	# todo 不用go get，改为dep
	$(GOGET) github.com/markbates/goth
	$(GOGET) github.com/markbates/pop
# todo benchmark go test -bench .
benchmark:
	$(GOBENCHMARK)

# Cross compilation
build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_LINUX) -v
build-arm:
	CGO_ENABLED=0 GOOS=linux GOARCH=arm $(GOBUILD) -o $(BINARY_ARM) -v
docker-build:
	docker run --rm -it -v "$(GOPATH)":/go -w /go/src/github.com/26huitailang/golang-web golang:latest go build -o "$(BINARY_LINUX)" -v

# documentation
doc:
	echo "access http://127.0.0.1:6060"
	godoc -http=:6060

# go tool
pprof:
	$(GOTOOL) pprof $(PROFILE)