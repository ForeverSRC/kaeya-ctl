SHELL := /usr/bin/env bash

GO_BIN=$(shell go env GOPATH)/bin
GO_CMD=go
GO_TEST=$(GO_BIN) test

build_win:
	GOOS=windows GOARCH=amd64 go build -o ./bin/windows/kaeya-ctl.exe ./cmd/kaeya-ctl

build_linux:
	GOOS=linux GOARCH=amd64 go build -o ./bin/linux/kaeya-ctl ./cmd/kaeya-ctl

build_mac:
	GOOS=darwin GOARCH=amd64 go build -o ./bin/macos/kaeya-ctl $(BUILD_ARGS) ./cmd/kaeya-ctl

test:
	@$(GO_CMD) clean -testcache
	@$(GO_CMD) test -v -short --tags test -timeout 120s ./...

build: build_mac build_linux build_win

all: test build