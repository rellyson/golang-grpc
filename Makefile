prepare:
	go mod download

generate:
	@protoc -I=proto --go_out=proto/gen-code --go-grpc_out=proto/gen-code \
	 --go-grpc_opt=paths=source_relative -I. proto/*.proto


build_client: prepare
	@go build -o ./bin/client-bin ./client/*.go

build_server: prepare
	@go build -o ./bin/server-bin ./server/*.go

# .PHONY is used for reserving tasks words
.PHONY: generate build_client build_server prepare

# supress echo commands on cli
.SILENT: generate