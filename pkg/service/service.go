package service

import (
	gotodo "github.com/grancc/go-to-do-app"
	"github.com/grancc/go-to-do-app/pkg/repository"
)

type Authorization interface {
	CreateUser(user gotodo.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type TodoList interface {
	Create(userId int, list gotodo.ToDoList) (int, error)
	GetAll(userId int) ([]gotodo.ToDoList, error)
	GetById(userId, listid int) (gotodo.ToDoList, error)
	UpdateList(userId, listid int, input gotodo.UpdateListInput) error
	Delete(userId, listid int) error
}

type ToodItem interface {
	Create(userId, listId int, listItem gotodo.ToDoItem) (int, error)
	GetAllItems(userId, listId int) ([]gotodo.ToDoItem, error)
	GetItemById(userId, itemId int) (gotodo.ToDoItem, error)
	DeleteItem(userId, itemId int) error
	UpdateItem(userId, itemId int, input gotodo.UpdateListItemInput) error
}

type Service struct {
	Authorization
	TodoList
	ToodItem
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo.Authorization),
		TodoList:      NewListService(repo.TodoList),
		ToodItem:      NewListItemService(repo.TodoListItem, repo.TodoList),
	}
}
