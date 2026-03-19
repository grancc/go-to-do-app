package repository

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

func NewRepository() *Repository {
	return &Repository{}
}
