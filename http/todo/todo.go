package todo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/src/todoService/model"
	"github.com/src/todoService/service"
)

type TodoHandler struct {
	todoSvc service.Todo
}

func NewHttpHandler(todoSvc service.Todo) TodoHandler {
	return TodoHandler{todoSvc: todoSvc}
}

func (th TodoHandler) TodoHandler(res http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodPost:
		th.CreateHandler(res, req)
	case http.MethodGet:
		th.GetHandler(res, req)
	case http.MethodPut:
		th.UpdateHandler(res, req)
	case http.MethodDelete:
		th.DeleteHandler(res, req)
	default:
		http.Error(res, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (th TodoHandler) CreateHandler(res http.ResponseWriter, req *http.Request) {

	reqBody, err := ioutil.ReadAll(req.Body)

	var todo model.Todo

	err = json.Unmarshal(reqBody, &todo)

	if err != nil {
		errStr, _ := json.Marshal(model.Error{Stage: "http", Message: "invalid request body", Error: err})
		res.Write(errStr)
	}

	_, err = th.todoSvc.Create(&todo)

	if err != nil {
		fmt.Println(err)
		return
	}
}

func (th TodoHandler) GetHandler(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Hi Get"))
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
