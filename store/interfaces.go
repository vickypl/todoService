package store

import "github.com/todoService/model"

type Todo interface {
	Create(todo *model.Todo) (*model.Todo, error)
	Get(filter model.Filter) ([]*model.Todo, error)
	GetByID(id string) (*model.Todo, error)
	Update(todo *model.Todo) (*model.Todo, error)
	Delete(id string) error
}
