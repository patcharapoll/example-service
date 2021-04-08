PROJECTNAME := $(shell basename "$(PWD)")
GOCMD = go
OS := $(shell uname -s | awk '{print tolower($$0)}')
GOARCH := amd64
GOBUILD = build

## stringer: generate error code
.PHONY: stringer
stringer:
	stringer -type ErrorCode internal/constants/error_code.go

## bin: build go server to binary
.PHONY: build
build:
	env CGO_ENABLED=0 GOOS=$(OS) GOARCH=${GOARCH} go build -a -installsuffix cgo -o bin/server cmd/server/main.go

.PHONY: watch
watch:
	CompileDaemon -include=Makefile --build="make build" --command=./bin/server --color=true --log-prefix=false

## pbgen: generate protobuf file
.PHONY: pbgen
pbgen:
	protoc --proto_path=internal/api/v1 --go_out=plugins=grpc:pkg/grpc/health/v1 health.proto
	protoc --proto_path=internal/api/v1 --go_out=plugins=grpc:pkg/api/v1 ping_pong.proto
	protoc-go-inject-tag -input=pkg/api/v1/ping_pong.pb.go

.PHONY: help
all: help
help: Makefile
	@echo
	@echo " Choose a command run in "$(PROJECTNAME)":"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo