///go:build unit

package service

import (
	"rsoi-1/internal/model"
	errors "rsoi-1/internal/model/error"
	"rsoi-1/internal/tests/helpers"
	"rsoi-1/internal/tests/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type PersonServiceSuite struct {
	suite.Suite
	repo    *mocks.IPersonRepo
	service IPersonService
	ds      helpers.PersonDataSource
}

func (s *PersonServiceSuite) SetupTest() {
	s.repo = new(mocks.IPersonRepo)
	s.service = NewPersonService(s.repo)
	s.ds.Init()
}

func (s *PersonServiceSuite) TearDownTest() {
}

func TestPersonServiceSuite(t *testing.T) {
	suite.Run(t, new(PersonServiceSuite))
}

func (s *PersonServiceSuite) TestListPersons() {
	s.T().Run("Not Empty", func(t *testing.T) {
		// Arrange
		s.repo.On("SelectAll").Return([]model.Person{
			*s.ds.Get(1),
			*s.ds.Get(2),
		}, nil)
		// Act
		res, err := s.service.ListPersons()
		// Assert
		assert.NoError(t, err)
		assert.Equal(t, res, []model.PersonResponse{
			s.ds.Get(1).ToResponse(),
			s.ds.Get(2).ToResponse(),
		})
	})
}

func (s *PersonServiceSuite) TestGetPerson() {
	s.T().Run("Id found", func(t *testing.T) {
		// Arrange
		id := s.ds.IdFound
		s.repo.On("SelectById", id).Return(*s.ds.Get(id), nil)
		// Act
		res, err := s.service.GetPerson(id)
		// Assert
		assert.NoError(t, err)
		assert.Equal(t, res, s.ds.Get(id).ToResponse())
	})

	s.T().Run("Id not found", func(t *testing.T) {
		// Arrange
		id := s.ds.IdNotFound
		s.repo.On("SelectById", id).Return(model.Person{}, errors.NotFound)
		// Act
		_, err := s.service.GetPerson(id)
		// Assert
		assert.Error(t, err)
		assert.Equal(t, err, errors.NotFound)
	})
}

func (s *PersonServiceSuite) TestCreatePerson() {
	p1 := s.ds.Get(1)

	s.T().Run("Every fields filled", func(t *testing.T) {
		// Arrange
		req := model.PersonRequest{
			Name:    p1.Name,
			Age:     p1.Age,
			Address: p1.Address,
			Work:    p1.Work,
		}
		p := model.Person{}
		p.FromRequest(&req)
		s.repo.On("Insert", &p).Return(int32(0), nil)
		// Act
		_, err := s.service.CreatePerson(&req)
		// Assert
		assert.NoError(t, err)
	})

	s.T().Run("Required fields filled", func(t *testing.T) {
		// Arrange
		req := model.PersonRequest{
			Name: "new user",
			Age:  p1.Age,
		}
		p := model.Person{}
		p.FromRequest(&req)
		s.repo.On("Insert", &p).Return(int32(0), nil)
		// Act
		_, err := s.service.CreatePerson(&req)
		// Assert
		assert.NoError(t, err)
	})

	s.T().Run("Required fields not filled", func(t *testing.T) {
		// Arrange
		req := model.PersonRequest{
			Age: p1.Age,
		}
		p := model.Person{}
		p.FromRequest(&req)
		s.repo.On("Insert", &p).Return(int32(0), errors.Unknown)
		// Act
		_, err := s.service.CreatePerson(&req)
		// Assert
		assert.Error(t, err)
	})
}

func (s *PersonServiceSuite) TestEditPerson() {
	p1 := s.ds.Get(1)

	s.T().Run("Id found", func(t *testing.T) {
		// Arrange
		req := model.PersonRequest{
			Name: "new user",
			Age:  p1.Age,
		}
		p := model.Person{}
		p.FromRequest(&req)
		s.repo.On("UpdateById", &p).Return(nil).Once()
		s.repo.On("SelectById", p.Id).Return(p, nil).Once()
		// Act
		pr, err := s.service.EditPerson(p.Id, &req)
		// Assert
		assert.NoError(t, err)
		assert.Equal(t, pr, p.ToResponse())
	})

	s.T().Run("Id not found", func(t *testing.T) {
		// Arrange
		req := model.PersonRequest{
			Name: "new user",
			Age:  p1.Age,
		}
		p := model.Person{}
		p.FromRequest(&req)
		s.repo.On("UpdateById", &p).Return(errors.NoAffected).Once()
		// Act
		pr, err := s.service.EditPerson(p.Id, &req)
		// Assert
		assert.Error(t, err)
		assert.Equal(t, err, errors.NoAffected)
		assert.Equal(t, pr, model.PersonResponse{})
	})
}

func (s *PersonServiceSuite) TestDeletePerson() {
	s.T().Run("Id found", func(t *testing.T) {
		// Arrange
		id := s.ds.IdFound
		s.repo.On("DeleteById", id).Return(nil)
		// Act
		err := s.service.DeletePerson(id)
		// Assert
		assert.NoError(t, err)
	})

	s.T().Run("Id not found", func(t *testing.T) {
		// Arrange
		id := s.ds.IdNotFound
		s.repo.On("DeleteById", id).Return(errors.NoAffected)
		// Act
		err := s.service.DeletePerson(id)
		// Assert
		assert.Error(t, err)
		assert.Equal(t, err, errors.NoAffected)
	})
}
