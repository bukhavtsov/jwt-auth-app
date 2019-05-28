package main

import (
	"context"
	"fmt"
	"github.com/bukhavtsov/jwt-auth-app/gateway-server/pkg/proto"
	"github.com/bukhavtsov/restful-app/pkg/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

type server struct {
	client authProto.AuthClient
}

func (s *server) ReadAllDevelopers(ctx context.Context, req *serverProto.ReadAllDevelopersRequest) (resp *serverProto.ReadAllDevelopersResponse, err error) {
	authReq := authProto.ReadAllDevelopersRequest{}
	authResp, err := s.client.ReadAllDevelopers(ctx, &authReq)
	if err != nil {
		log.Println(err)
	}
	var respDevelopers []*serverProto.Developer
	for _, currentDev := range authResp.Developers {
		developer := &serverProto.Developer{
			Id:           currentDev.Id,
			Name:         currentDev.Name,
			Age:          currentDev.Age,
			PrimarySkill: currentDev.PrimarySkill,
		}
		respDevelopers = append(respDevelopers, developer)
	}
	return &serverProto.ReadAllDevelopersResponse{Developers: respDevelopers}, nil
}

func (s *server) CreateDeveloper(ctx context.Context, req *serverProto.CreateDeveloperRequest) (*serverProto.CreateDeveloperResponse, error) {
	developer := authProto.Developer{
		Name:         req.Developer.Name,
		Age:          req.Developer.Age,
		PrimarySkill: req.Developer.PrimarySkill,
	}
	authReq := authProto.CreateDeveloperRequest{
		Developer: &developer,
	}
	authResp, err := s.client.CreateDeveloper(ctx, &authReq)
	if err != nil {
		log.Println(err)
	}
	return &serverProto.CreateDeveloperResponse{Id: authResp.Id}, nil
}

func (s *server) ReadDeveloper(ctx context.Context, req *serverProto.ReadDeveloperRequest) (*serverProto.ReadDeveloperResponse, error) {
	authReq := authProto.ReadDeveloperRequest{Id: req.Id}
	authResp, err := s.client.ReadDeveloper(ctx, &authReq)
	if err != nil {
		log.Println(err)
	}
	respDeveloper := &serverProto.Developer{
		Id:           authResp.Developer.Id,
		Name:         authResp.Developer.Name,
		Age:          authResp.Developer.Age,
		PrimarySkill: authResp.Developer.PrimarySkill,
	}
	return &serverProto.ReadDeveloperResponse{Developer: respDeveloper}, nil
}

func (s *server) UpdateDeveloper(ctx context.Context, req *serverProto.UpdateDeveloperRequest) (*serverProto.UpdateDeveloperResponse, error) {
	developer := authProto.Developer{
		Id:           req.Developer.Id,
		Name:         req.Developer.Name,
		Age:          req.Developer.Age,
		PrimarySkill: req.Developer.PrimarySkill,
	}
	authReq := authProto.UpdateDeveloperRequest{
		Developer: &developer,
		Id:        req.Id,
	}
	authResp, err := s.client.UpdateDeveloper(ctx, &authReq)
	if err != nil {
		log.Println(err)
	}
	respDeveloper := &serverProto.Developer{
		Id:           authResp.Developer.Id,
		Name:         authResp.Developer.Name,
		Age:          authResp.Developer.Age,
		PrimarySkill: authResp.Developer.PrimarySkill,
	}
	return &serverProto.UpdateDeveloperResponse{Developer: respDeveloper}, nil
}

func (s *server) DeleteDeveloper(ctx context.Context, req *serverProto.DeleteDeveloperRequest) (*serverProto.DeleteDeveloperResponse, error) {
	authReq := authProto.DeleteDeveloperRequest{
		Id: req.Id,
	}
	_, err := s.client.DeleteDeveloper(ctx, &authReq)
	if err != nil {
		log.Println(err)
	}
	return &serverProto.DeleteDeveloperResponse{}, nil
}

func (s *server) SignIn(ctx context.Context, req *serverProto.SignInRequest) (*serverProto.SignInResponse, error) {
	authReq := authProto.SignInRequest{
		Login:    req.Login,
		Password: req.Password,
	}
	authResp, err := s.client.SignIn(ctx, &authReq)
	if err != nil {
		log.Println(err)
	}
	return &serverProto.SignInResponse{RefreshToken: authResp.RefreshToken, AccessToken: authResp.AccessToken}, nil
}

func (s *server) SignUp(ctx context.Context, req *serverProto.SignUpRequest) (*serverProto.SignUpResponse, error) {
	user := authProto.User{
		Id:           req.User.Id,
		Login:        req.User.Login,
		Email:        req.User.Email,
		Password:     req.User.Password,
		Role:         req.User.Role,
		RefreshToken: req.User.RefreshToken,
	}
	authReq := authProto.SignUpRequest{
		User: &user,
	}
	authResp, err := s.client.SignUp(ctx, &authReq)
	if err != nil {
		log.Println(err)
	}
	return &serverProto.SignUpResponse{RefreshToken: authResp.RefreshToken, AccessToken: authResp.AccessToken}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":2020")
	if err != nil {
		log.Fatal(err)
	}
	srv := grpc.NewServer()
	cc, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	client := authProto.NewAuthClient(cc)
	serverProto.RegisterRestAppServer(srv, &server{client})
	reflection.Register(srv)
	fmt.Println("server")
	if e := srv.Serve(listener); e != nil {
		log.Fatal(err)
	}
}
