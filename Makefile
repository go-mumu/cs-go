GOPATH:=$(shell go env GOPATH)

.PHONY:local
local:
	@go run server.go -c "./config/local.toml" -l "./log.log"

# api 依赖注入
.PHONY:inject-service
inject-service:
	@wire ./service/container

# service 依赖注入
.PHONY:inject-api
inject-api:
	@wire ./api/container

.PHONY:gen-model
gen-model:
	@go run ./service/cli/main.go gen-model

.PHONY:proto
proto:
	protoc --proto_path=./proto \
	 	   --proto_path=/Users/weiyi/go/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.16.0/third_party/googleapis \
	 	   --go_out=./proto/pb \
	 	   --go-grpc_out=./proto/pb \
	 	   --grpc-gateway_out=./proto/pb \
	 	   ./proto/*.proto