CUR_DIR := $(shell dirname $(abspath $(lastword $(MAKEFILE_LIST))))
BINDIR = $(CUR_DIR)/bin
DBDIR = $(CUR_DIR)/devops/db
GATEWAY_IMAGE = gateway-jwt
SERVER_IMAGE = server-jwt
RPC_IMAGE = rpc-jwt
DB_IMAGE = psql-jwt
gen:
	protoc -I /usr/local/include -I. \
	-I $(GOPATH)/src \
	-I $(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
	--go_out=plugins=grpc:. \
	--grpc-gateway_out=logtostderr=true:. \
	--swagger_out=logtostderr=true:. \
	gateway-server/pkg/proto/auth.proto

	protoc -I /usr/local/include -I. \
    	-I $(GOPATH)/src \
    	-I $(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
    	--go_out=plugins=grpc:. \
    	--grpc-gateway_out=logtostderr=true:. \
    	--swagger_out=logtostderr=true:. \
    	rpc/pkg/proto/auth.proto

build:
	go build -o bin/rpc cmd/rpc/main.go
	go build -o bin/gateway cmd/gateway-server/gateway/main.go
	go build -o bin/server cmd/gateway-server/server/main.go
up:
	cp cmd/gateway-server/gateway/Dockerfile $(BINDIR)
	docker build -t $(GATEWAY_IMAGE) $(BINDIR)
	rm $(BINDIR)/Dockerfile
	cp cmd/gateway-server/server/Dockerfile $(BINDIR)
	docker build -t $(SERVER_IMAGE) $(BINDIR)
	rm $(BINDIR)/Dockerfile
	cp cmd/rpc/Dockerfile $(BINDIR)
	docker build -t $(RPC_IMAGE) $(BINDIR)
	rm $(BINDIR)/Dockerfile
	docker build -t $(DB_IMAGE) $(DBDIR)
compose:
	docker-compose up