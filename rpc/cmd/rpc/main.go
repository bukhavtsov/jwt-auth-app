package main

import (
	"log"
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

var rpcAddr string
var host string
var port string
var user string
var dbname string
var password string
var sslmode string

func init() {
	rpcAddr = os.Getenv("listen")
	port = os.Getenv("database.port")
	user = os.Getenv("database.user")
	dbname = os.Getenv("database.name")
	password = os.Getenv("database.password")
	sslmode = os.Getenv("database.ssl")
	host = os.Getenv("database.host")
}

func isEmptyAddr() bool {
	if rpcAddr == "" {
		return true
	}
	return false
}

func isEmptyDBAddr() bool {
	if host == "" || port == "" || user == "" || dbname == "" || password == "" || sslmode == "" {
		return true
	}
	return false
}

func getDefaultAddr() {
	rpcAddr = "127.0.0.1:8080"
}

func getDefaultDBAddr() {
	host = "localhost"
	port = "5432"
	user = "postgres"
	dbname = "postgres"
	password = "root"
	sslmode = "disable"
}

func emptyEnvMessage() {
	log.Println("env variables are not found")
	log.Println("default address has been setted")
	log.Println("rpc addr:", rpcAddr)
}

func emptyEnvDBMessage() {
	log.Println("env db variables are not found")
	log.Println("default db address has been setted")
	log.Println("host:", host)
	log.Println("port:", port)
	log.Println("user:", user)
	log.Println("dbname:", dbname)
	log.Println("password:", password)
	log.Println("sslmode:", sslmode)
}

func main() {
	if isEmptyAddr() {
		getDefaultAddr()
		emptyEnvMessage()
	}
	if isEmptyDBAddr() {
		getDefaultDBAddr()
		emptyEnvDBMessage()
	}
	listener, err := net.Listen("tcp", rpcAddr)
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
