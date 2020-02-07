FROM golang:latest

MAINTAINER Alopex6414 "alopex6414@outlook.com"

WORKDIR $GOPATH/src/satellite
COPY . .

RUN go get -v -t -d ./...
RUN go install -v ./...

EXPOSE 80
ENTRYPOINT ["satellite https -ip 0.0.0.0 -port 80"]