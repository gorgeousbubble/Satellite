# Satellite Project - Makefile(Golang)
# Copyright(C) 2019, Team Gorgeous Bubble, All Rights Reserved.

# Golang Command
GO 		= go
GOBUILD = $(GO) build
GOCLEAN = $(GO) clean
GOTEST 	= $(GO) test
GOGET 	= $(GO) get

# Binary Parameters
BIN 	= satellite

# Build
all: test build

build: 
	$(GOBUILD) -o $(BIN) -v

test:
	$(GOTEST) -v ./...

clean:
	$(GOCLEAN)
	rm -rf $(BIN)

run:
	$(GOBUILD) -o $(BIN) -v ./...
	./$(BIN)

deps:

# Cross-Compile
