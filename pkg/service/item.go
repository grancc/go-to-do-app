package service

import (
	gotodo "github.com/grancc/go-to-do-app"
	"github.com/grancc/go-to-do-app/pkg/repository"
)

type TodoListItemService struct {
	repo     repository.TodoListItem
	listRepo repository.TodoList
}

func NewListItemService(repo repository.TodoListItem, listrepo repository.TodoList) *TodoListItemService {
	return &TodoListItemService{
		repo:     repo,
		listRepo: listrepo,
	}
}

func (s *TodoListItemService) Create(userId, listId int, listItem gotodo.ToDoItem) (int, error) {
	_, err := s.listRepo.GetById(userId, listId)
	if err != nil {
		return 0, err
	}
	return s.repo.Create(listId, listItem)
}

func (s *TodoListItemService) GetAllItems(userId, listId int) ([]gotodo.ToDoItem, error) {
	_, err := s.listRepo.GetById(userId, listId)
	if err != nil {
		return nil, err
	}
	return s.repo.GetAllItems(listId)
}

func (s *TodoListItemService) GetItemById(userId, itemId int) (gotodo.ToDoItem, error) {
	return s.repo.GetItemById(userId, itemId)
}

func (s *TodoListItemService) DeleteItem(userId, itemId int) error {
	return s.repo.DeleteItem(userId, itemId)
}

func (s *TodoListItemService) UpdateItem(userId, itemId int, input gotodo.UpdateListItemInput) error {
	return s.repo.UpdateItem(userId, itemId, input)
}
