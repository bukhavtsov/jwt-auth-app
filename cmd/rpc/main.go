package main

import (
	"github.com/bukhavtsov/jwt-auth-app/devops/db"
	"github.com/bukhavtsov/jwt-auth-app/rpc/pkg/data"
	authProto "github.com/bukhavtsov/jwt-auth-app/rpc/pkg/proto"
	"github.com/bukhavtsov/jwt-auth-app/rpc/pkg/server"
	"github.com/bukhavtsov/jwt-auth-app/rpc/pkg/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

const (
	engine   = "postgres"
	username = "postgres"
	password = "root"
	name     = "postgres"
	address  = ":8080"
	network  = "tcp"
)

func main() {
	listener, err := net.Listen(network, address)
	if err != nil {
		panic(err)
	}
	srv := grpc.NewServer()
	conn := db.GetConnection(engine, username, password, name)
	defer conn.Close()
	auth := service.NewAuthService(data.NewUserData(conn))
	developer := service.NewDeveloperService(data.NewDeveloperData(conn))
	authProto.RegisterAuthServer(srv, server.NewAuthServer(auth, developer))
	reflection.Register(srv)
	if e := srv.Serve(listener); e != nil {
		panic(e)
	}
}
