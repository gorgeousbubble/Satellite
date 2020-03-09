# Erika Project - Makefile(Golang)
# Copyright(C) 2020, Team Gorgeous Bubble, All Rights Reserved.

# Golang Commands
GO 	= go
GOBUILD = $(GO) build
GOGET 	= $(GO) get

# Binary Parameters
GOBASE  = $(shell pwd)
GOBIN   = $(GOBASE)/bin
MKBIN   = $(shell mkdir -p $(GOBIN))

# Docker Commands
DOCKER      = docker
DOCKERBUILD = $(DOCKER) build
DOCKERRUN   = $(DOCKER) run

# Application
APPNAME	= erika
APPDIST = erika.tar.gz
APPPATH = bin/$(APPNAME)

# Build
all: build

build:
	$(MKBIN)
	$(GOBUILD) -o $(GOBIN)

build_image:
	$(MKBIN)
	$(GOBUILD) -o $(GOBIN)
	$(DOCKERBUILD) -t $(APPNAME) .

deps:
	$(GOGET) -v -t -d ./...

# Cross-Compile
