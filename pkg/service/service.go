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
}

type ToodItem interface {
}

type Service struct {
	Authorization
	TodoList
	ToodItem
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo.Authorization),
	}
}
