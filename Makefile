# Satellite Project - Makefile(Golang)
# Copyright(C) 2019, Team Gorgeous Bubble, All Rights Reserved.

# Golang Command
GO 	= go
GOBUILD = $(GO) build
GOCLEAN = $(GO) clean
GOTEST 	= $(GO) test
GOGET 	= $(GO) get

# Binary Parameters
GOBASE  = $(shell pwd)
GOBIN   = $(GOBASE)/bin

# Build
all: test build

build: 
	$(GOBUILD) -o $(GOBIN)

test:
	$(GOTEST) -v -cover -benchmem -bench .

clean:
	$(GOCLEAN)
	rm -rf $(GOBIN)

run:
	$(GOBUILD) -o $(GOBIN)
	./$(GOBIN)

deps:
	$(GOGET) -v -t -d ./...

# Cross-Compile
