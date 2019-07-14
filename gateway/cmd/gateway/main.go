package main

import (
	"context"
	"fmt"
	pb "github.com/bukhavtsov/jwt-auth-app/gateway/pkg/proto"
	"github.com/bukhavtsov/jwt-auth-app/rpc/pkg/jwt"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"log"
	"net/http"
	"os"
	"strings"
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
	connGRPC, err := grpc.Dial(serviceAddr, grpc.WithInsecure(), )
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
	mux.Handle("/", authHandler(gatewayConn))
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

func authHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.RequestURI)
		signUpReq := r.RequestURI == "/api/v1/signup" && r.Method == "POST"
		signInReq := strings.Contains(r.RequestURI, "/api/v1/signin/") && r.Method == "GET"
		if !signUpReq && !signInReq {
			accessToken := r.Header.Get("accessToken")
			refreshToken := r.Header.Get("refreshToken")
			err := jwt.Validate(&accessToken, &refreshToken)
			if err != nil {
				log.Println("tokens are not found")
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("tokens are not found"))
				return
			}
			w.Header().Set("accessToken", accessToken)
		}
		h.ServeHTTP(w, r)
	})
}
