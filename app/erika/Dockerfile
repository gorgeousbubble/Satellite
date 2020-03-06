FROM golang:latest

MAINTAINER Alopex6414 "alopex6414@outlook.com"

WORKDIR $GOPATH/src/satellite
COPY . .

cd ./app/erika/

RUN go get -v -t -d ./...
RUN go install -v ./...

EXPOSE 80
ENTRYPOINT ["$GOPATH/bin/erika","start","-ip","0.0.0.0","-port","80"]