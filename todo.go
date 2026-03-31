package gotodo

import "errors"

type ToDoList struct {
	Id          int    `json:"-" db:"id"`
	Title       string `json:"title" binding:"required" db:"title"`
	Description string `json:"description" db:"description"`
}

type UserList struct {
	Id     int
	UserId int
	ListId int
}

type ToDoItem struct {
	Id          int    `json:"-"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Done        bool   `json:"done"`
}

type ListItem struct {
	Id     int
	ItemId int
	ListId int
}

type UpdateListInput struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
}

func (i UpdateListInput) Validate() error {
	if i.Title == nil && i.Description == nil {
		return errors.New("update structure has no values")
	}

	return nil
}

type UpdateListItemInput struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
	Done        *bool   `json:"done"`
}

func (i UpdateListItemInput) Validate() error {
	if i.Title == nil && i.Description == nil && i.Done == nil {
		return errors.New("update structure has no values")
	}

	return nil
}
