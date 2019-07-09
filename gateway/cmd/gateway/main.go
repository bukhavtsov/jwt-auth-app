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

var addr string
var svcGateway string

func init() {
	addr = os.Getenv("listen")
	svcGateway = os.Getenv("svc-gateway")
}

// TODO: unary interceptor
func HTTPProxy(proxyAddr string, serviceAddr string) {
	log.Println("gateway work in:", addr)
	log.Println("svc-gateway work in:", svcGateway)
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
	HTTPProxy(addr, svcGateway)
}
