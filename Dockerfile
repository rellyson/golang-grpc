FROM golang:1.17-alpine 

WORKDIR /opt/app

RUN apk add --no-cache autoconf git build-base \
    automake libtool make g++

# Install the protocol compiler plugins for Go
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1

COPY . .