package service

import (
	gotodo "github.com/grancc/go-to-do-app"
	"github.com/grancc/go-to-do-app/pkg/repository"
)

type TodoListService struct {
	repo repository.TodoList
}

func NewListService(repo repository.TodoList) *TodoListService {
	return &TodoListService{repo: repo}
}

func (s *TodoListService) Create(userId int, list gotodo.ToDoList) (int, error) {
	return s.repo.Create(userId, list)
}

func (s *TodoListService) GetAll(userId int) ([]gotodo.ToDoList, error) {
	return s.repo.GetAll(userId)
}

func (s *TodoListService) GetById(userId, listid int) (gotodo.ToDoList, error) {
	return s.repo.GetById(userId, listid)
}

func (s *TodoListService) Delete(userId, listid int) error {
	return s.repo.Delete(userId, listid)
}

func (s *TodoListService) UpdateList(userId, listid int, input gotodo.UpdateListInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repo.UpdateList(userId, listid, input)
}
