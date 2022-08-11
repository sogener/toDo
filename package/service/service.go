package service

import (
	"toDo"
	"toDo/package/repository"
)

type Authorization interface {
	CreateUser(user toDo.User) (int, error)
}

type ToDoList interface{}

type ToDoItem interface{}

type Service struct {
	Authorization
	ToDoList
	ToDoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
