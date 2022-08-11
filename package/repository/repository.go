package repository

import (
	"github.com/jmoiron/sqlx"
	"toDo"
)

type Authorization interface {
	CreateUser(user toDo.User) (int, error)
}

type ToDoList interface{}

type ToDoItem interface{}

type Repository struct {
	Authorization
	ToDoList
	ToDoItem
}

// NewRepository this is constructor
func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
