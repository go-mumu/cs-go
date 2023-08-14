GOPATH:=$(shell go env GOPATH)

.PHONY:local-api
local-api:
	@go run api.go \
		--consul-addr="http://127.0.0.1:8500" \
		--consul-config-path="config/local"

.PHONY:local-service1
local-service1:
	@go run server.go \
		--log-path="./log.log" \
		--consul-addr="http://127.0.0.1:8500" \
		--consul-config-path="config/local" \
		--node-ip="127.0.0.1" \
		--node-port=9991 \
		--node-id="node1" \
		--http-port=10001

.PHONY:local-service2
local-service2:
	@go run server.go \
		--log-path="./log.log" \
		--consul-addr="http://127.0.0.1:8500" \
		--consul-config-path="config/local" \
		--node-ip="127.0.0.1" \
		--node-port=9992 \
		--node-id="node2" \
		--http-port=10002

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