package repository

import (
	"database/sql"
	"rsoi-1/internal/model"
	errors "rsoi-1/internal/model/error"
)

type PersonSqlRepo struct {
	db *sql.DB
}

const (
	// returning - Postgresql
	SelectAllQuery  = "select * from Persons"
	SelectByIdQuery = "select * from Persons where id=$1"
	InsertQuery     = "insert into Persons(name, age, work, address) values ($1, $2, $3, $4) returning id"
	UpdateByIdQuery = `update Persons set
		name=coalesce($2,name), age=coalesce($3,age), work=coalesce($4,work), address=coalesce($5,address)
		where id=$1`
	DeleteByIdQuery = "delete from Persons where id=$1"
)

func NewPersonSqlRepo(db *sql.DB) IPersonRepo {
	return &PersonSqlRepo{db}
}

func scanPerson(row MultiScanner, p *model.Person) error {
	err := row.Scan(
		&p.Id,
		&p.Name,
		&p.Age,
		&p.Work,
		&p.Address,
	)
	return err
}

func (r *PersonSqlRepo) SelectAll() (ps []model.Person, err error) {
	rows, err := r.db.Query(SelectAllQuery)
	if err != nil {
		return nil, err
	}

	var p model.Person
	defer rows.Close()
	for rows.Next() {
		err = scanPerson(rows, &p)
		if err != nil {
			return
		}
		ps = append(ps, p)
	}
	return
}

func (r *PersonSqlRepo) SelectById(id int32) (p model.Person, err error) {
	row := r.db.QueryRow(SelectByIdQuery, id)
	err = scanPerson(row, &p)
	if err == sql.ErrNoRows {
		err = errors.NotFound
	}
	return p, err
}

func (r *PersonSqlRepo) Insert(p *model.Person) (id int32, err error) {
	row := r.db.QueryRow(InsertQuery, p.Name, p.Age, p.Work, p.Address)
	err = row.Scan(&id)
	return
}

func (r *PersonSqlRepo) UpdateById(p *model.Person) error {
	res, err := r.db.Exec(UpdateByIdQuery, p.Id, p.Name, p.Age, p.Work, p.Address)
	if err == nil {
		err = handleRowsAffected(res)
	}
	return err
}

func (r *PersonSqlRepo) DeleteById(id int32) error {
	res, err := r.db.Exec(DeleteByIdQuery, id)
	if err == nil {
		err = handleRowsAffected(res)
	}
	return err
}
