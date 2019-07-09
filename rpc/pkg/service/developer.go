package service

import (
	"github.com/bukhavtsov/jwt-auth-app/rpc/pkg/data"
	"log"
)

type DeveloperService interface {
	GetDevelopers() ([]*data.Developer, error)
	GetDeveloper(id int64) (*data.Developer, error)
	CreateDeveloper(developer data.Developer) (int64, error)
	UpdateDeveloper(id int64, developer data.Developer) (*data.Developer, error)
	DeleteDeveloper(id int64) error
}

type developerService struct {
	data data.DeveloperData
}

func NewDeveloperService(data data.DeveloperData) DeveloperService {
	return developerService{data: data}
}
func (s developerService) GetDevelopers() ([]*data.Developer, error) {
	developers, err := s.data.ReadAll()
	if err != nil {
		log.Println("developers haven't been read")
		return nil, err
	}
	return developers, nil
}

func (api developerService) GetDeveloper(id int64) (*data.Developer, error) {
	developer, err := api.data.Read(id)
	if err != nil {
		log.Println("developer hasn't been read")
		return nil, err
	}
	return developer, nil
}

func (s developerService) CreateDeveloper(developer data.Developer) (int64, error) {
	id, err := s.data.Create(&developer)
	if err != nil {
		log.Println("developer hasn't been created")
		return -1, err
	}
	return id, nil
}

func (s developerService) UpdateDeveloper(id int64, developer data.Developer) (*data.Developer, error) {
	developer.Id = id
	updated, err := s.data.Update(&developer)
	if err != nil {
		return nil, err
	}
	return updated, nil
}

func (s developerService) DeleteDeveloper(id int64) error {
	err := s.data.Delete(id)
	if err != nil {
		log.Println("developer hasn't been removed")
		return err
	}
	return nil
}
