package todo

import (
	"github.com/todoService/model"
	"github.com/todoService/service"
	"github.com/todoService/store"
)

type todo struct {
	todoStore store.Todo
}

func NewService(todoStore store.Todo) service.Todo {
	return todo{todoStore: todoStore}
}

func (td todo) Create(todo *model.Todo) (*model.Todo, error) {
	return td.todoStore.Create(todo)
}

func (td todo) Get(filter model.Filter) ([]*model.Todo, error) {
	return td.todoStore.Get(filter)
}

func (td todo) GetByID(todoID string) (*model.Todo, error) {
	return td.todoStore.GetByID(todoID)
}

func (td todo) Update(todo *model.Todo) (*model.Todo, error) {
	todo, err := td.GetByID(todo.Id)
	if err != nil {
		return nil, err
	}

	return td.todoStore.Update(todo)
}

func (td todo) Delete(todoID string) error {
	todo, err := td.GetByID(todoID)
	if err != nil {
		return err
	}

	return td.todoStore.Delete(todo.Id)
}
