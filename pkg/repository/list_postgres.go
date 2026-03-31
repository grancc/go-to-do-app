package repository

import (
	"fmt"
	"strings"

	gotodo "github.com/grancc/go-to-do-app"
	"github.com/jmoiron/sqlx"
)

type ToDoListPostgres struct {
	db *sqlx.DB
}

func NewToDoListPostgres(db *sqlx.DB) *ToDoListPostgres {
	return &ToDoListPostgres{db: db}
}

func (t *ToDoListPostgres) Create(userId int, list gotodo.ToDoList) (int, error) {
	tx, err := t.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int
	createListQuery := fmt.Sprintf("Insert into %s (title, description) values ($1, $2) returning id ", todoListsTable)
	row := tx.QueryRow(createListQuery, list.Title, list.Description)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	createUsersListQuery := fmt.Sprintf("Insert into %s (userid, listid) values ($1, $2) ", userToTodolistTable)
	_, err = tx.Exec(createUsersListQuery, userId, id)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}

func (t *ToDoListPostgres) GetAll(userId int) ([]gotodo.ToDoList, error) {
	var lists []gotodo.ToDoList
	query := fmt.Sprintf("select tl.id, tl.title, tl.description from %s tl inner join %s ul on (tl.id=ul.listid) where ul.userid = $1", todoListsTable, userToTodolistTable)
	err := t.db.Select(&lists, query, userId)
	return lists, err
}

func (t *ToDoListPostgres) GetById(userId, listid int) (gotodo.ToDoList, error) {
	var list gotodo.ToDoList
	query := fmt.Sprintf("select tl.id, tl.title, tl.description from %s tl inner join %s ul on (tl.id=ul.listid) where ul.userid = $1 and tl.id = $2",
		todoListsTable, userToTodolistTable)
	err := t.db.Get(&list, query, userId, listid)
	return list, err
}

func (t *ToDoListPostgres) Delete(userId, listid int) error {
	query := fmt.Sprintf("delete from %s tl using %s ul where tl.id=ul.listid and ul.userid = $1 and tl.id = $2",
		todoListsTable, userToTodolistTable)
	_, err := t.db.Exec(query, userId, listid)
	return err
}

func (t *ToDoListPostgres) UpdateList(userId, listid int, input gotodo.UpdateListInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
		args = append(args, input.Title)
		argId++
	}
	if input.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argId))
		args = append(args, input.Description)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")
	query := fmt.Sprintf("update %s tl set %s from %s ul where tl.id=ul.listid and ul.listid=$%d and ul.userid = $%d",
		todoListsTable, setQuery, userToTodolistTable, argId, argId+1)
	args = append(args, listid, userId)

	_, err := t.db.Exec(query, args...)
	return err
}
