package main

import (
	"fmt"
	pb "github.com/bukhavtsov/jwt-auth-app/gateway-server/pkg/proto"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"net/http"
)

const (
	proxyAddr   = ":8081"
	serviceAddr = "127.0.0.1:2020"
)

//TODO: unary interceptor
func HTTPProxy(proxyAddr string, serviceAddr string) {
	connGRPC, err := grpc.Dial(serviceAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalln("failed to connect to GRPC ", err)
	}
	defer connGRPC.Close()
	gatewayConn := runtime.NewServeMux()
	err = pb.RegisterRestAppHandler(context.Background(), gatewayConn, connGRPC)
	if err != nil {
		log.Fatalln("failed to start HTTP server", err)
	}
	mux := http.NewServeMux()
	mux.Handle("/", gatewayConn)
	fmt.Println("gateway")
	log.Fatalln(http.ListenAndServe(proxyAddr, mux))
}

func main() {
	HTTPProxy(proxyAddr, serviceAddr)
}
