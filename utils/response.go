package utils

import (
	"encoding/json"
	"myapp/model"
	"net/http"
)

func Return(w http.ResponseWriter, status bool, code int, err error, data interface{}) {

	json.NewEncoder(w).Encode(model.Response{
		Status: status,
		Code:   code,
		Error:  err.Error(),
		Data:   data,
	})
}
