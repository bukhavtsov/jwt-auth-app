package service

import (
	"errors"
	"github.com/bukhavtsov/jwt-auth-app/rpc/pkg/data"
	"github.com/bukhavtsov/jwt-auth-app/rpc/pkg/jwt"
	"log"
)

type AuthService interface {
	SingIn(login, password string) (*token, error)
	SignUp(user data.User) (*token, error)
}

type token struct {
	Access, Refresh string
}

type authService struct {
	data data.UserData
}

func NewAuthService(data data.UserData) AuthService {
	return authService{data: data}
}

func (s authService) SingIn(login, password string) (*token, error) {
	user, err := s.data.Get(login, password)
	if err != nil {
		log.Println("user hasn't been found")
		return nil, err
	}
	refresh, err := jwt.GenerateRefresh(*user)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	_, err = s.data.Update(user, refresh)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	access, err := jwt.GenerateAccess(*user)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &token{access, refresh}, nil
}

func (s authService) SignUp(user data.User) (*token, error) {
	_, err := s.data.Get(user.Login, user.Password)
	if err == nil {
		log.Println("user has been found")
		return nil, errors.New("user has been found")
	}
	_, err = s.data.Create(&user)
	if err != nil {
		log.Println("user hasn't been created")
		return nil, err
	}
	user.RefreshToken, err = jwt.GenerateRefresh(user)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	access, err := jwt.GenerateAccess(user)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	refresh := user.RefreshToken
	return &token{access, refresh}, nil
}
