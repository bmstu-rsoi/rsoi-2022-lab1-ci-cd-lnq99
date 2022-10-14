package model

type Person struct {
	Id      int32
	Name    string
	Age     *int32
	Address *string
	Work    *string
}

type PersonRequest struct {
	Name    string  `json:"name" validate:"required"`
	Age     *int32  `json:"age,omitempty"`
	Address *string `json:"address,omitempty"`
	Work    *string `json:"work,omitempty"`
}

type PersonResponse struct {
	Id      int32   `json:"id"`
	Name    string  `json:"name"`
	Age     *int32  `json:"age,omitempty"`
	Address *string `json:"address,omitempty"`
	Work    *string `json:"work,omitempty"`
}

func (p *Person) FromRequest(r *PersonRequest) {
	p.Name = r.Name
	p.Age = r.Age
	p.Address = r.Address
	p.Work = r.Work
}

func (p *Person) ToResponse() PersonResponse {
	return PersonResponse{
		Id:      p.Id,
		Name:    p.Name,
		Age:     p.Age,
		Address: p.Address,
		Work:    p.Work,
	}
}
