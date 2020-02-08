# Satellite 🌠
The Satellite Golang Project.

## Introduction
Satellite is a tool for both multiple command and web service develop by Golang. This project is my first program design by golang.So I try to use golang to written some useful tools for both Windows and Linux. Now it contains many function such as packet, unpack, compress, decompress, etc. I will continue to contribute to this project about new feature and make it more funny and useful.

The project name is Satellite. 2019 is the 50th anniversary of the Apollo 11 moon landing. So it represents the love of human beings and the fearless exploration spirit of that era. And encourage us to keep going and never give up.

## Installation and usage
Install golang and download package from [https://golang.org](https://golang.org)  
  
  | OS            | Architecture  | Version  | Binary                  |
  | ------------- |:-------------:|:--------:| :---------------------: |
  | Windows       | amd64         | 1.13.6 ↑ | go1.13.6 windows/amd64  |
  | Linux         | adm64         | 1.13.6 ↑ | go1.13.6 linux/amd64    |
  | Linux         | arm           | 1.13.6 ↑ | go1.13.6 linux/arm      |

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
  * Build All in One:
  * `cd $GOPATH/src/satellite`  
  * `docker build -t satellite .`

