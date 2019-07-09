CUR_DIR := $(shell dirname $(abspath $(lastword $(MAKEFILE_LIST))))
BINDIR = $(CUR_DIR)/bin
DBDIR = $(CUR_DIR)/devops/db
GATEWAY_IMAGE = gateway
SVC_GATEWAY_IMAGE = svc-gateway
RPC_IMAGE = rpc
DB_IMAGE = db-rpc
gen:
	protoc -I /usr/local/include -I. \
	-I $(GOPATH)/src \
	-I $(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
	--go_out=plugins=grpc:. \
	--grpc-gateway_out=logtostderr=true:. \
	--swagger_out=logtostderr=true:. \
	gateway/pkg/proto/auth.proto

	protoc -I /usr/local/include -I. \
    	-I $(GOPATH)/src \
    	-I $(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
    	--go_out=plugins=grpc:. \
    	--grpc-gateway_out=logtostderr=true:. \
    	--swagger_out=logtostderr=true:. \
    	rpc/pkg/proto/auth.proto

build:
	go build -o bin/rpc rpc/cmd/rpc/main.go
	go build -o bin/gateway gateway/cmd/gateway/main.go
	go build -o bin/svc-gateway gateway/cmd/svc-gateway/main.go
image:
	cp gateway/cmd/gateway/Dockerfile $(BINDIR)
	docker build -t $(GATEWAY_IMAGE) $(BINDIR)
	rm $(BINDIR)/Dockerfile
	cp gateway/cmd/svc-gateway/Dockerfile $(BINDIR)
	docker build -t $(SVC_GATEWAY_IMAGE) $(BINDIR)
	rm $(BINDIR)/Dockerfile
	cp rpc/cmd/rpc/Dockerfile $(BINDIR)
	docker build -t $(RPC_IMAGE) $(BINDIR)
	rm $(BINDIR)/Dockerfile
	docker build -t $(DB_IMAGE) $(DBDIR)
up:
	docker-compose up

#down:
#	docker kill $(docker ps -a -q)
#	docker rm $(docker ps -a -q)
#	docker rmi $(docker images -a -q)