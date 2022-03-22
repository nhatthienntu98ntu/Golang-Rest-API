package ultils

import (
	"encoding/json"
	"net/http"

	"github.com/nhatthienntu98ntu/Golang-Rest-API/model"
)

func Return(w http.ResponseWriter, status bool, code int, err error, data interface{}) {
	w.Header().Set("Content-type", "application/json")
	response := model.Response{
		Status: status,
		Code:   code,
		Error:  "",
		Data:   data,
	}
	if err != nil {
		response.Error = err.Error()
	}
	json.NewEncoder(w).Encode(response)
}
