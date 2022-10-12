package service

import (
	"rsoi-1/internal/model"
	"rsoi-1/internal/repository"
)

type PersonServiceImpl struct {
	repo repository.PersonRepo
}

func NewPersonService(repo repository.PersonRepo) IPersonService {
	return &PersonServiceImpl{repo}
}

func (s *PersonServiceImpl) ListPersons() ([]model.PersonResponse, error) {
	persons, err := s.repo.SelectAll()
	l := len(persons)
	pr := make([]model.PersonResponse, l)
	for i := 0; i < l; i++ {
		pr[i] = persons[i].ToResponse()
	}
	return pr, err
}

func (s *PersonServiceImpl) GetPerson(id int32) (model.PersonResponse, error) {
	p, err := s.repo.SelectById(id)
	return p.ToResponse(), err
}

func (s *PersonServiceImpl) CreatePerson(pr *model.PersonRequest) (int32, error) {
	p := model.Person{}
	p.FromRequest(pr)
	return s.repo.Insert(&p)
}

func (s *PersonServiceImpl) EditPerson(id int32, pr *model.PersonRequest) error {
	p := model.Person{Id: id}
	p.FromRequest(pr)

	return s.repo.UpdateById(&p)
}

func (s *PersonServiceImpl) DeletePerson(id int32) error {
	return s.repo.DeleteById(id)
}
