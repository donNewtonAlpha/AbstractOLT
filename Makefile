SERVER_OUT	 := "bin/AbstractOLT"
CLIENT_OUT	 := "bin/client"
API_OUT		 := "api/abstract_olt_api.pb.go"
API_REST_OUT     := "api/abstract_olt_api.pb.gw.go"
SWAGGER_OUT      := "api/abstract_olt_api.swagger.json"
PKG	         := "github.com/donNewtonAlpha/AbstractOLT"
SERVER_PKG_BUILD := "${PKG}/cmd/AbstractOLT"
CLIENT_PKG_BUILD := "${PKG}/client"
PKG_LIST := $(shell go list ${PKG}/... | grep -v /vendor/)


.PHONY: all api server client

all: server client

api/abstract_olt_api.pb.go: api/abstract_olt_api.proto
	@protoc -I api/ \
	-I${GOPATH}/src \
	-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
	--go_out=plugins=grpc:api \
	api/abstract_olt_api.proto

api/abstract_olt_api.pb.gw.go : api/abstract_olt_api.proto
	  @protoc -I api/ \
	-I${GOPATH}/src \
	-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
	--grpc-gateway_out=logtostderr=true:api \
	 api/abstract_olt_api.proto

swagger:
	protoc -I api/ \
  -I${GOPATH}/src \
  -I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  --swagger_out=logtostderr=true:api \
  api/abstract_olt_api.proto

api: api/abstract_olt_api.pb.go api/abstract_olt_api.pb.gw.go swagger

dep: ## Get the dependencies
	@go get -v -d ./...

server: dep api ## Build the binary file for server
	@go build -i -v -o $(SERVER_OUT) $(SERVER_PKG_BUILD)

client: dep api ## Build the binary file for client
	@go build -i -v -o $(CLIENT_OUT) $(CLIENT_PKG_BUILD)

clean: ## Remove previous builds
	@rm $(SERVER_OUT) $(CLIENT_OUT) $(API_OUT) $(API_REST_OUT) $(SWAGGER_OUT)

help: ## Display this help screen
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
