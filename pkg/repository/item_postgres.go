package repository

import (
	"fmt"
	"strings"

	gotodo "github.com/grancc/go-to-do-app"
	"github.com/jmoiron/sqlx"
)

type ToDoListItemPostgres struct {
	db *sqlx.DB
}

func NewToDoListItemPostgres(db *sqlx.DB) *ToDoListItemPostgres {
	return &ToDoListItemPostgres{db: db}
}

func (t *ToDoListItemPostgres) Create(listId int, listItem gotodo.ToDoItem) (int, error) {
	tx, err := t.db.Begin()
	if err != nil {
		return 0, err
	}

	var itemId int
	createItemQuery := fmt.Sprintf("insert into %s (title, description) values ($1, $2) returning id",
		todoItemsTable)

	row := tx.QueryRow(createItemQuery, listItem.Title, listItem.Description)
	if err := row.Scan(&itemId); err != nil {
		tx.Rollback()
		return 0, err
	}

	createItemListQuery := fmt.Sprintf("insert into %s (itemid, listid) values ($1, $2)", listsItemsTable)
	_, err = tx.Exec(createItemListQuery, itemId, listId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return itemId, tx.Commit()
}

func (t *ToDoListItemPostgres) GetAllItems(listId int) ([]gotodo.ToDoItem, error) {
	var listItems []gotodo.ToDoItem
	query := fmt.Sprintf("select it.id, it.title, it.description, it.done from %s it inner join %s li on (li.itemid=it.id) where li.listid=$1",
		todoItemsTable, listsItemsTable)

	err := t.db.Select(&listItems, query, listId)

	return listItems, err
}

func (t *ToDoListItemPostgres) GetItemById(userId, itemId int) (gotodo.ToDoItem, error) {
	var listItem gotodo.ToDoItem
	query := fmt.Sprintf("select it.id, it.title, it.description, it.done from %s it inner join %s li on (li.itemid=it.id) inner join %s ul on (li.listid=ul.listid) where ul.userid=$1 and it.id=$2",
		todoItemsTable, listsItemsTable, userToTodolistTable)

	if err := t.db.Get(&listItem, query, itemId, userId); err != nil {
		return listItem, err
	}

	return listItem, nil
}

func (t *ToDoListItemPostgres) DeleteItem(userId, itemId int) error {
	query := fmt.Sprintf("delete from  %s it using %s li, %s ul where li.itemid=it.id and li.listid=ul.listid and ul.userid=$1 and it.id=$2",
		todoItemsTable, listsItemsTable, userToTodolistTable)

	_, err := t.db.Exec(query, userId, itemId)
	return err
}

func (t *ToDoListItemPostgres) UpdateItem(userId, itemId int, input gotodo.UpdateListItemInput) error {
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
	if input.Done != nil {
		setValues = append(setValues, fmt.Sprintf("done=$%d", argId))
		args = append(args, input.Done)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("update %s it set %s from %s li, %s ul where li.itemid=it.id and li.listid=ul.listid and ul.userid=$%d and it.id=$%d",
		todoItemsTable, setQuery, listsItemsTable, userToTodolistTable, argId, argId+1)

	args = append(args, userId, itemId)

	_, err := t.db.Exec(query, args...)
	return err
}
