package service

import "github.com/src/todoService/model"

type Todo interface {
	Create(todo *model.Todo) (*model.Todo, error)
	Get(filter model.Filter) ([]*model.Todo, error)
	GetByID(todoID string) (*model.Todo, error)
	Update(todo *model.Todo) (*model.Todo, error)
	Delete(todoID string) error
}
