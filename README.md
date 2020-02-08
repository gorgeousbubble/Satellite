# Satellite 🌠
The Satellite Golang Project.

[![Documentation](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](http://godoc.org/github.com/gorgeousbubble/satellite)
[![Version](https://img.shields.io/github/release/gorgeousbubble/satellite.svg?style=flat-square)](https://github.com/gorgeousbubble/satellite/releases/latest)

## Introduction
Satellite is a tool for both multiple command and web service develop by Golang. This project is my first program design by golang.So I try to use golang to written some useful tools for both Windows and Linux. Now it contains many function such as packet, unpack, compress, decompress, etc. I will continue to contribute to this project about new feature and make it more funny and useful.

The project name is Satellite. 2019 is the 50th anniversary of the Apollo 11 moon landing. So it represents the love of human beings and the fearless exploration spirit of that era. And encourage us to keep going and never give up.

## Installation and usage
Install golang and download package from [https://golang.org](https://golang.org)  
  
  | OS            | Architecture  | Version  | Binary                  |
  | ------------- |:-------------:|:--------:| :---------------------: |
  | Windows       | amd64         | 1.13.6 ▲ | go1.13.6 windows/amd64  |
  | Linux         | adm64         | 1.13.6 ▲ | go1.13.6 linux/amd64    |
  | Linux         | arm           | 1.13.6 ▲ | go1.13.6 linux/arm      |

#### Use Git to clone this project  
  `git clone https://github.com/gorgeousbubble/satellite.git`  

#### Install dependencies  
  You can use `go get -u ...` to install dependencies one by one. In Linux, you can enter `$GOPATH/src` and use `make deps` to install all package.
  
#### Build native or container
Build the project:  
  * Build the project in Windows  
    `cd $GOPATH/src/satellite`  
    `build.bat`  
  * Build the project in Linux or ARM  
    `cd $GOPATH/src/satellite`  
    `make build`
    
Build docker image:  
  * Build all in one  
    `cd $GOPATH/src/satellite`  
    `docker build -t satellite .` or `make build_image`
  * Build application  
    first build native application, then `cp Dockerfile_BIN bin\Dockerfile` and use command `docker build -t satellite .` build image.
    
Run docker container:  
  `docker run --name=satellite -p 8080:80 -d satellite`  

Build Docker image need download base image 'ubuntu:latest' or 'golang:latest' from Docker Hub.

#### Usage of Satellite
Use command `./satellite --help` see how to use it. Start HTTPS service and listen on port 8080:  
  `./satellite https -ip 127.0.0.1 -port 8080`  
  
#### Test the project
Test the project:  
  * Test the project in Windows  
    `cd $GOPATH/src/satellite`  
    `test.bat`  
  * Test the project in Linux or ARM  
    `cd $GOPATH/src/satellite`  
    `make test`

#### Format the code
Format the project:  
  * Format the project in Windows  
    `cd $GOPATH/src/satellite`  
    `fmt.bat`  
  * Format the project in Linux or ARM  
    `cd $GOPATH/src/satellite`  
    `make fmt`
    
#### Clean environment
Clean environment:  
  * Clean the project in Windows  
    `cd $GOPATH/src/satellite`  
    `clean.bat`  
  * Clean the project in Linux or ARM  
    `cd $GOPATH/src/satellite`  
    `make clean`
    
## API document
If opened in a browser, the import path itself leads to the API documentation:  
  * http://godoc.org/github.com/gorgeousbubble/satellite
  
You can also refer to doc folder and I will add markdown file to give some description.
    
## License
The Satellite project is licensed under the Apache License 2.0. Please see the LICENSE file for details.
