package server

import (
	"github.com/bukhavtsov/jwt-auth-app/rpc/pkg/data"
	"github.com/bukhavtsov/jwt-auth-app/rpc/pkg/proto"
	"github.com/bukhavtsov/jwt-auth-app/rpc/pkg/service"
	"golang.org/x/net/context"
	"log"
)

type server struct {
	auth      service.AuthService
	developer service.DeveloperService
}

func NewAuthServer(auth service.AuthService, developer service.DeveloperService) *server {
	return &server{auth: auth, developer: developer}
}

func (s server) ReadAllDevelopers(ctx context.Context, req *authProto.ReadAllDevelopersRequest) (*authProto.ReadAllDevelopersResponse, error) {
	developers, err := s.developer.GetDevelopers()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var respDevelopers []*authProto.Developer
	for i := 0; i < len(developers); i++ {
		respDeveloper := &authProto.Developer{
			Id:           developers[i].Id,
			Name:         developers[i].Name,
			Age:          developers[i].Age,
			PrimarySkill: developers[i].PrimarySkill,
		}
		respDevelopers = append(respDevelopers, respDeveloper)
	}
	return &authProto.ReadAllDevelopersResponse{Developers: respDevelopers}, nil
}

func (s server) CreateDeveloper(ctx context.Context, req *authProto.CreateDeveloperRequest) (*authProto.CreateDeveloperResponse, error) {
	developer := data.Developer{
		Id:           req.Developer.Id,
		Name:         req.Developer.Name,
		Age:          req.Developer.Age,
		PrimarySkill: req.Developer.PrimarySkill,
	}
	id, err := s.developer.CreateDeveloper(developer)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &authProto.CreateDeveloperResponse{Id: id}, nil
}

func (s server) ReadDeveloper(ctx context.Context, req *authProto.ReadDeveloperRequest) (*authProto.ReadDeveloperResponse, error) {
	developer, err := s.developer.GetDeveloper(req.Id)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	respDeveloper := &authProto.Developer{
		Id:           developer.Id,
		Name:         developer.Name,
		Age:          developer.Age,
		PrimarySkill: developer.PrimarySkill,
	}
	return &authProto.ReadDeveloperResponse{Developer: respDeveloper}, nil
}

func (s server) UpdateDeveloper(ctx context.Context, req *authProto.UpdateDeveloperRequest) (*authProto.UpdateDeveloperResponse, error) {
	developer := data.Developer{
		Id:           req.Developer.Id,
		Name:         req.Developer.Name,
		Age:          req.Developer.Age,
		PrimarySkill: req.Developer.PrimarySkill,
	}
	updated, err := s.developer.UpdateDeveloper(req.Id, developer)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	respUpdated := &authProto.Developer{
		Id:           updated.Id,
		Name:         updated.Name,
		Age:          updated.Age,
		PrimarySkill: updated.PrimarySkill,
	}
	return &authProto.UpdateDeveloperResponse{Developer: respUpdated}, nil
}

func (s server) DeleteDeveloper(ctx context.Context, req *authProto.DeleteDeveloperRequest) (*authProto.DeleteDeveloperResponse, error) {
	err := s.developer.DeleteDeveloper(req.Id)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return nil, nil //what need to return?
}

func (s server) SignIn(ctx context.Context, req *authProto.SignInRequest) (*authProto.SignInResponse, error) {
	token, err := s.auth.SingIn(req.Login, req.Password)
	log.Println("token is: ", token)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &authProto.SignInResponse{AccessToken: token.Refresh, RefreshToken: token.Access}, nil
}

func (s server) SignUp(ctx context.Context, req *authProto.SignUpRequest) (*authProto.SignUpResponse, error) {
	user := data.User{
		Id:           req.User.Id,
		Login:        req.User.Login,
		Email:        req.User.Email,
		Password:     req.User.Password,
		Role:         req.User.Role,
		RefreshToken: req.User.RefreshToken,
	}
	token, err := s.auth.SignUp(user)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &authProto.SignUpResponse{AccessToken: token.Refresh, RefreshToken: token.Access}, nil
}
