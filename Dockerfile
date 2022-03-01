FROM golang:1.17-alpine 

WORKDIR /opt/app

ENV PROTOC_ZIP=protoc-3.19.4-linux-x86_64.zip

RUN apk add --no-cache autoconf git build-base \
    automake libtool make g++ unzip curl

# Install the protocol compiler and pĺugins for Go
RUN curl -OL https://github.com/protocolbuffers/protobuf/releases/download/v3.19.4/$PROTOC_ZIP \
    && unzip -o $PROTOC_ZIP -d /usr/bin/protoc && \
    unzip -o $PROTOC_ZIP -d /usr/bin 'include/*' && \
    chmod +x /usr/bin/protoc && \
    rm -f $PROTOC_ZIP
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1

COPY . .