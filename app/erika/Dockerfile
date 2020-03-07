FROM ubuntu:latest

MAINTAINER Alopex6414 "alopex6414@outlook.com"

RUN mkdir -p /app
COPY ./bin/erika /app

EXPOSE 80
ENTRYPOINT ["/app/erika","start","-ip","0.0.0.0","-port","80"]