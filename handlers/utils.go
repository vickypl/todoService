package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/src/todoService/model"
)

func JWTMiddlewareAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
		fmt.Print("JWT Auth  will be added")
	})
}

func ErrorResponseWriter(res http.ResponseWriter, errorMessage model.Error, errorCode int) {
	errStr, _ := json.Marshal(errorMessage)
	res.WriteHeader(errorCode)
	res.Write(errStr)
}

func ResponseWrapper(res http.ResponseWriter, data []*model.Todo) {
	type resp struct {
		Data []*model.Todo `json:"data"`
	}

	result, err := json.Marshal(resp{Data: data})
	if err != nil {
		ErrorResponseWriter(res, model.Error{Stage: "http", Error: err, Message: "Marshalling error"}, http.StatusInternalServerError)
		return
	}

	res.Write(result)
	return
}

func ValidateFilters(queryParams url.Values) (model.Filter, error) {
	filter := model.Filter{}

	if queryParams.Get("userid") != "" {
		filter["userid"] = queryParams.Get("userid")
	} else if queryParams.Get("title") != "" {
		filter["title"] = queryParams.Get("title")
	} else if queryParams.Get("discription") != "" {
		filter["discription"] = queryParams.Get("discription")
	} else if queryParams.Get("priority") != "" {
		filter["priority"] = queryParams.Get("priority")
	} else if queryParams.Get("status") != "" {
		filter["status"] = queryParams.Get("status")
	} else {
		return nil, errors.New("Invalid Filter")
	}

	return filter, nil
}
