# Satellite Project - Makefile(Golang)
# Copyright(C) 2019, Team Gorgeous Bubble, All Rights Reserved.

# Golang Commands
GO 	= go
GOBUILD = $(GO) build
GOCLEAN = $(GO) clean
GOTEST 	= $(GO) test
GOGET 	= $(GO) get

# Binary Parameters
GOBASE  = $(shell pwd)
GOBIN   = $(GOBASE)/bin
MKBIN   = $(shell mkdir -p $(GOBIN))

# Docker Commands
DOCKER = docker
DOCKERBUILD = $(DOCKER) build
DOCKERRUN	= $(DOCKER) run

# Application
APPNAME	= satellite

# Build
all: test build

build:
	$(MKBIN)
	$(GOBUILD) -o $(GOBIN)

build_image:
	$(DOCKERBUILD) -t $(APPNAME) .

test:
	$(GOTEST) -v -cover -benchmem -bench .

clean:
	$(GOCLEAN)
	rm -rf $(GOBIN)

run:
	$(GOBUILD) -o $(GOBIN)
	./$(GOBIN)

run_container:
	$(DOCKERRUN) -it --rm --name $(APPNAME) -p 8080:8080 -d $(APPNAME)

deps:
	$(GOGET) -v -t -d ./...

# Cross-Compile
