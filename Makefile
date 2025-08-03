.DEFAULT_GOAL := help

DST ?= SnakeGame

.PHONY: run build build-windows-amd64 build-linux-amd64 build-darwin-amd64 build-darwin-arm64 build-all

run:
	@go run cmd/main.go

build:
	@go build -o $(DST) cmd/main.go

build-windows-amd64:
	@GOOS=windows GOARCH=amd64 go build -o bin/windows/SnakeGame.exe cmd/main.go

# build-linux-amd64:
# 	@GOOS=linux GOARCH=amd64 go build -o bin/linux/SnakeGame cmd/main.go

# build-darwin-amd64:
# 	@GOOS=darwin GOARCH=amd64 go build -o bin/macOS/SnakeGame_x64 cmd/main.go

build-darwin-arm64:
	@GOOS=darwin GOARCH=arm64 go build -o bin/macOS/SnakeGame cmd/main.go

build-all: build-windows-amd64 build-linux-amd64 build-darwin-amd64 build-darwin-arm64

help:
	@echo "Usage:"