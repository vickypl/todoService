package todo

import (
	"database/sql"

	"github.com/src/todoService/model"
	"github.com/src/todoService/store"
)

const (
	insertQuery = "insert into todo (id, userid, title, discription, priority, status) values (?, ?, ?, ?, ?, ?);"
	updateQuery = "update todo set userid=?, title=?, discription=?, priority=?, status=? where id=?;"
	selectQuery = "select * from todo where id=?;"
	deleteQuery = "delete from todo where id=?;"
)

type todo struct {
	db *sql.DB
}

func NewTodoStore(db *sql.DB) store.Todo {
	return todo{db: db}
}

func (st todo) Create(todo *model.Todo) (*model.Todo, error) {
	todoID := store.GenerateID()
	_, err := st.db.Exec(insertQuery, todo.UserID, todo.Title, todo.Discription, todo.Priority, todo.Status, todoID)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (st todo) Get(filter model.Filter) ([]*model.Todo, error) {
	query, err := store.QueryGenerator(filter)
	if err != nil {
		return nil, err
	}

	var todoList []*model.Todo
	row, err := st.db.Query(query)
	if err != nil {
		return nil, err
	}

	for row.Next() {
		var todo model.Todo
		err := row.Scan(&todo.Id, &todo.UserID, &todo.Title, &todo.Discription, &todo.Priority, &todo.Status)
		if err != nil {
			return nil, err
		}

		todoList = append(todoList, &todo)
	}

	return todoList, nil
}

func (st todo) GetByID(todoID string) (*model.Todo, error) {
	row := st.db.QueryRow(selectQuery, todoID)
	var todo model.Todo

	err := row.Scan(&todo.Id, &todo.UserID, &todo.Title, &todo.Discription, &todo.Priority, &todo.Status)
	if err != nil {
		return nil, err
	}

	return &todo, nil
}

func (st todo) Update(todo *model.Todo) (*model.Todo, error) {
	_, err := st.db.Exec(updateQuery, todo.UserID, todo.Title, todo.Discription, todo.Priority, todo.Status, todo.Id)
	if err != nil {
		return nil, err
	}

	updatedTodo, err := st.GetByID(todo.Id)
	if err != nil {
		return nil, err
	}

	return updatedTodo, nil
}

func (st todo) Delete(todoID string) error {
	_, err := st.db.Exec(deleteQuery, todoID)
	if err != nil {
		return err
	}

	return nil
}
