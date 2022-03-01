FROM golang:1.17-alpine 

WORKDIR /opt/app

RUN apk add --no-cache autoconf git build-base \
    automake libtool make g++

COPY . .
