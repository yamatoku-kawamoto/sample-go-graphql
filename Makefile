GOVERSION=$(shell go version)
GOOS:=$(shell go env GOOS)
GOARCH:=$(shell go env GOARCH)

OUTPUT:="bin/app.exe"

SYSTEM_VERSION="1.0.0-alpha"
BUILD_FLAGS:=-trimpath -ldflags="-s -w -X main.VERSION=$(SYSTEM_VERSION) -X main.MODE=Release"

BUILD_FILES=

.PHONY: generate
generate:
	@go run github.com/99designs/gqlgen generate
	
.PHONY: go/run
go/run:
	@go run cmd/server.go

.PHONY: build/go
go/build:
	go build -o $(OUTPUT) $(BUILD_FLAGS) $(BUILD_FILES)
	
.PHONY: build/container
build/container:
	podman build -t go-graphql:latest .
