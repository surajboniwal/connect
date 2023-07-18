package util

import (
	"encoding/json"
	"net/http"
)

type response struct {
	Status bool        `json:"status"`
	Errors *[]AppError `json:"errors,omitempty"`
	Data   any         `json:"data,omitempty"`
}

func WriteJSONResponse(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")

	res := response{}

	if _, ok := data.(*[]AppError); ok {
		res.Status = false
		res.Data = nil
		res.Errors = data.(*[]AppError)
	} else {
		res.Status = true
		res.Data = data
		res.Errors = nil
	}

	json.NewEncoder(w).Encode(res)
}
