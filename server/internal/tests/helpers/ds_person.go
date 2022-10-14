package helpers

import (
	"fmt"
	"rsoi-1/internal/model"
)

type PersonDataSource struct {
	data       map[int32]*model.Person
	IdNotFound int32
	IdFound    int32
}

func (s *PersonDataSource) Init() {
	s.data = make(map[int32]*model.Person)
	for i := int32(1); i <= 5; i++ {
		age := 20 + i
		name := fmt.Sprintf("name %d", i)
		addr := fmt.Sprintf("addr %d", i)
		work := fmt.Sprintf("work %d", i)
		s.data[i] = &model.Person{
			i,
			name,
			&age,
			&addr,
			&work,
		}
	}
	s.IdFound = 2
	s.IdNotFound = 10
	//fmt.Println(s.data[1])
	//fmt.Println(s.data[2])
}

func (s *PersonDataSource) Get(id int32) *model.Person {
	p, ok := s.data[id]
	if !ok {
		return &model.Person{}
	}
	return p
}
