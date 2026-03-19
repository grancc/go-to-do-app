package repository

import "github.com/jmoiron/sqlx"

type Authorization interface {
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
	return &Repository{}
}
