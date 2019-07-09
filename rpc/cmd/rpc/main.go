package main

import (
	"net"
	"os"

	"github.com/bukhavtsov/jwt-auth-app/devops/db"
	"github.com/bukhavtsov/jwt-auth-app/rpc/pkg/data"
	authProto "github.com/bukhavtsov/jwt-auth-app/rpc/pkg/proto"
	"github.com/bukhavtsov/jwt-auth-app/rpc/pkg/server"
	"github.com/bukhavtsov/jwt-auth-app/rpc/pkg/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var addr string
var host string
var port string
var user string
var dbname string
var password string
var sslmode string

func init() {
	addr = os.Getenv("listen")
	port = os.Getenv("database.port")
	user = os.Getenv("database.user")
	dbname = os.Getenv("database.name")
	password = os.Getenv("database.password")
	sslmode = os.Getenv("database.ssl")
	host = os.Getenv("database.host")
}

func main() {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}
	srv := grpc.NewServer()
	conn := db.GetConnection(host, port, user, dbname, password, sslmode)
	defer conn.Close()
	auth := service.NewAuthService(data.NewUserData(conn))
	developer := service.NewDeveloperService(data.NewDeveloperData(conn))
	authProto.RegisterAuthServer(srv, server.NewAuthServer(auth, developer))
	reflection.Register(srv)
	if e := srv.Serve(listener); e != nil {
		panic(e)
	}
}
