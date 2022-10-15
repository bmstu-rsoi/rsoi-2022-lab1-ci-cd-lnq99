package service

import (
	"rsoi-1/internal/model"
	"rsoi-1/internal/repository"
)

var services Services

type IPersonService interface {
	ListPersons() ([]model.PersonResponse, error)
	GetPerson(id int32) (model.PersonResponse, error)
	CreatePerson(person *model.PersonRequest) (int32, error)
	EditPerson(id int32, person *model.PersonRequest) (model.PersonResponse, error)
	DeletePerson(id int32) error
}

type Services struct {
	Person IPersonService
}

func (*Services) Init(repo *repository.Repo) {
	services = Services{
		Person: NewPersonService(repo.Person),
	}
}

func GetServices() *Services {
	return &services
}

func InitServices(repo *repository.Repo) *Services {
	services.Init(repo)
	return &services
}
