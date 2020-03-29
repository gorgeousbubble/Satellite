FROM ubuntu:latest

MAINTAINER Alopex6414 "alopex6414@outlook.com"

WORKDIR /app

RUN apt-get update
RUN apt-get install -y\
    git
RUN mkdir -p .
RUN mkdir -p ./task
COPY ./bin/erika .
COPY ./repos/resources ./task

EXPOSE 80
ENTRYPOINT ["./erika","start","-ip","0.0.0.0","-port","80"]