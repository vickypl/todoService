package todo

import (
	"database/sql"

	"github.com/todoService/model"
	"github.com/todoService/store"
)

type todo struct {
	db *sql.DB
}

func NewTodoStore(db *sql.DB) store.Todo {
	return todo{db: db}
}

func (st todo) Create(todo *model.Todo) (*model.Todo, error) {
	return nil, nil
}

func (st todo) Get(filter model.Filter) ([]*model.Todo, error) {
	return nil, nil
}

func (st todo) GetByID(id string) (*model.Todo, error) {
	return nil, nil
}

func (st todo) Update(todo *model.Todo) (*model.Todo, error) {
	return nil, nil
}

func (st todo) Delete(id string) error {
	return nil
}
