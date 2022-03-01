FROM golang:1.17-alpine 

WORKDIR /opt/app

RUN apk add --no-cache autoconf git build-base \
    automake libtool make g++

COPY . .

## download deps and build golang bin
RUN go mod download
RUN go build -o bin/app main.go

## run binary
CMD ["./bin/app"]
