package todo

import (
	"encoding/json"
	"io/ioutil"
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

	reqBody, err := ioutil.ReadAll(req.Body)

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
	res.Write([]byte("Hi Get by id"))
}

func (th TodoHandler) UpdateHandler(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Hi Update"))
}

func (th TodoHandler) DeleteHandler(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Hi Delete"))
}
