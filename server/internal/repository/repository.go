package repository

import (
	"database/sql"
	"rsoi-1/internal/model"
	errors "rsoi-1/internal/model/error"
)

type PersonRepo interface {
	SelectAll() ([]model.Person, error)
	SelectById(id int32) (model.Person, error)
	Insert(person *model.Person) (int32, error)
	UpdateById(person *model.Person) error
	DeleteById(id int32) error
}

type Repo struct {
	Person PersonRepo
}

func NewSqlRepository(db *sql.DB) *Repo {
	return &Repo{
		Person: NewPersonSqlRepo(db),
	}
}

// MultiScanner created to support sql.Row and sql.Rows
type MultiScanner interface {
	Scan(dest ...any) error
}

func handleRowsAffected(res sql.Result) error {
	count, err := res.RowsAffected()
	if err == nil && count == 0 {
		err = errors.NoAffected
	}
	return err
}
