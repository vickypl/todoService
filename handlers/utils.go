package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/src/todoService/model"
)

func ErrorResponseWriter(res http.ResponseWriter, errorMessage model.Error, errorCode int) {
	errStr, _ := json.Marshal(errorMessage)
	res.WriteHeader(errorCode)
	res.Write(errStr)
}
