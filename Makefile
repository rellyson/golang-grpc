generate:
	@protoc --proto_path=proto --go_out=plugins=grpc:proto/gen --go_opt=paths=source_relative \
	proto/currencies.proto

build_client:
	@go build -o ./bin/client-bin ./main.go

build_server:
	@go build -o ./bin/server-bin ./main.go

# .PHONY is used for reserving tasks words
.PHONY: generate build_client build_server

# supress echo commands on cli
.SILENT: generate build_client build_server