# This Makefile is meant to be used by people that do not usually work
# with Go source code. If you know what GOPATH is then you probably
# don't need to bother with make.

.PHONY: all clean
OS = $(shell uname)
GOBIN = build/bin
ifeq ($(OS), Linux)
endif

ifeq ($(OS), Darwin)
endif

all:
	mkdir -p $(GOBIN)
	go build -v -o $(GOBIN)/bobo
clean:
	go clean -cache
	rm -rf $(GOBIN)/*
