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

type TodoListItem interface {
	Create(listId int, listItem gotodo.ToDoItem) (int, error)
	GetAllItems(listId int) ([]gotodo.ToDoItem, error)
	GetItemById(userId, itemId int) (gotodo.ToDoItem, error)
	UpdateItem(userId, itemId int, input gotodo.UpdateListItemInput) error
	DeleteItem(userId, itemId int) error
}

type Repository struct {
	Authorization
	TodoList
	TodoListItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		TodoList:      NewToDoListPostgres(db),
		TodoListItem:  NewToDoListItemPostgres(db),
	}
}
