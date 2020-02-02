FROM golang

WORKDIR /go/src/satellite
COPY . .

RUN go get -v -t -d ./...
RUN go install -v ./...

CMD ["satellite https -ip 127.0.0.1 -port 8080"]