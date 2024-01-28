package todo

import (
	"github.com/src/todoService/model"
	"github.com/src/todoService/service"
	"github.com/src/todoService/store"
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
	_, err := td.GetByID(todo.Id)
	if err != nil {
		return nil, err
	}

	UpdatedTodo, err := td.todoStore.Update(todo)
	if err != nil {
		return nil, err
	}

	return UpdatedTodo, nil
}

func (td todo) Delete(todoID string) error {
	todo, err := td.GetByID(todoID)
	if err != nil {
		return err
	}

	return td.todoStore.Delete(todo.Id)
}
