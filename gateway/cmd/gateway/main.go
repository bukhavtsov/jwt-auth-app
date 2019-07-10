package main

import (
	"fmt"
	pb "github.com/bukhavtsov/jwt-auth-app/gateway/pkg/proto"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"net/http"
	"os"
)

var gwAddr string
var svcGWAddr string

func init() {
	gwAddr = os.Getenv("listen")
	svcGWAddr = os.Getenv("svc-gateway")
}

func isEmptyAddr() bool {
	if gwAddr == "" || svcGWAddr == "" {
		return true
	}
	return false
}

func getDefaultValues() {
	gwAddr = "127.0.0.1:8081"
	svcGWAddr = "127.0.0.1:2020"
}

func emptyEnvMessage() {
	log.Println("env variables not found")
	log.Println("default addrs have been setted")
	log.Println("gateway:", gwAddr)
	log.Println("svc-gateway:", svcGWAddr)
}

// TODO: unary interceptor
func HTTPProxy(proxyAddr string, serviceAddr string) {
	log.Println("gateway work in:", gwAddr)
	log.Println("svc-gateway work in:", svcGWAddr)
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
	if isEmptyAddr() {
		getDefaultValues()
		emptyEnvMessage()
	}
	HTTPProxy(gwAddr, svcGWAddr)
}
