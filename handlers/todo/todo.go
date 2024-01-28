package todo

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/src/todoService/handlers"
	"github.com/src/todoService/model"
	"github.com/src/todoService/service"
)

type TodoHandler struct {
	todoSvc service.Todo
}

func NewHttpHandler(todoSvc service.Todo) TodoHandler {
	return TodoHandler{todoSvc: todoSvc}
}

func (th TodoHandler) CreateHandler(res http.ResponseWriter, req *http.Request) {

	res.Header().Set("Content-Type", "application/json")

	reqBody, err := io.ReadAll(req.Body)

	var todo model.Todo

	err = json.Unmarshal(reqBody, &todo)

	if err != nil {
		handlers.ErrorResponseWriter(res, model.Error{Stage: "http", Error: err, Message: "Error while unmarshalling"}, http.StatusInternalServerError)
		return
	}

	_, err = th.todoSvc.Create(&todo)

	if err != nil {
		handlers.ErrorResponseWriter(res, model.Error{Stage: "http", Error: err, Message: "Error while creating"}, http.StatusBadRequest)
		return
	}
}

func (th TodoHandler) GetHandler(res http.ResponseWriter, req *http.Request) {

	res.Header().Set("Content-Type", "application/json")

	queryParams := req.URL.Query()

	filter, err := handlers.ValidateFilters(queryParams)
	if err != nil {
		handlers.ErrorResponseWriter(res, model.Error{Stage: "http", Error: err, Message: "Invalid Filter Error"}, http.StatusBadRequest)
		return
	}

	todoList, err := th.todoSvc.Get(filter)
	if err != nil {
		handlers.ErrorResponseWriter(res, model.Error{Stage: "http", Error: err, Message: "Error while unmarshalling"}, http.StatusInternalServerError)
		return
	}

	handlers.ResponseWrapper(res, todoList)

	return
}

func (th TodoHandler) GetByIDHandler(res http.ResponseWriter, req *http.Request) {

	res.Header().Set("Content-Type", "application/json")

	queryParams := req.URL.Query()

	id := queryParams.Get("id")

	result, err := th.todoSvc.GetByID(id)
	if err != nil {
		handlers.ErrorResponseWriter(res, model.Error{Stage: "http", Error: err, Message: "Error while getting todo"}, http.StatusInternalServerError)
		return
	}

	handlers.ResponseWrapper(res, []*model.Todo{result})
}

func (th TodoHandler) UpdateHandler(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	reqBody, err := io.ReadAll(req.Body)

	var todo model.Todo

	err = json.Unmarshal(reqBody, &todo)

	if err != nil {
		handlers.ErrorResponseWriter(res, model.Error{Stage: "http", Error: err, Message: "Error while unmarshalling"}, http.StatusInternalServerError)
		return
	}

	updatedTodo, err := th.todoSvc.Update(&todo)
	if err != nil {
		handlers.ErrorResponseWriter(res, model.Error{Stage: "http", Error: err, Message: "Error while updating"}, http.StatusBadRequest)
		return
	}

	handlers.ResponseWrapper(res, []*model.Todo{updatedTodo})
}

func (th TodoHandler) DeleteHandler(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Hi Delete"))
}
