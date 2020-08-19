.PHONY: build
build:
	env GO111MODULE=off go build -v ./cmd/apiserver/

.DEFAULT_GOAL := build