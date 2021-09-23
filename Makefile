.PHONY: build
build:
	go build -v ./cmd/tcpchat

.DEFAULT_GOAL := build