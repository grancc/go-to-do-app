package repository

import (
	gotodo "github.com/grancc/go-to-do-app"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user gotodo.User) (int, error)
	GetUser(username, password string) (gotodo.User, error)
}

type TodoList interface {
	Create(userId int, list gotodo.ToDoList) (int, error)
	GetAll(userId int) ([]gotodo.ToDoList, error)
	GetById(userId, listid int) (gotodo.ToDoList, error)
	UpdateList(userId, listid int, input gotodo.UpdateListInput) error
	Delete(userId, listid int) error
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
		TodoList:      NewToDoListPostgres(db),
	}
}
