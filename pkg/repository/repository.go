package repository

import (
	gotodo "github.com/grancc/go-to-do-app"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user gotodo.User) (int, error)
}

type TodoList interface {
}

type ToodItem interface {
}

type Repository struct {
	Authorization
	TodoList
	ToodItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
