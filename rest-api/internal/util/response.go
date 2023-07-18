package util

import (
	"encoding/json"
	"net/http"
)

type response struct {
	Status bool      `json:"status"`
	Error  *AppError `json:"error,omitempty"`
	Data   any       `json:"data,omitempty"`
}

func WriteJSONResponse(w http.ResponseWriter, data interface{}, params ...int) {
	w.Header().Set("Content-Type", "application/json")

	res := response{}

	if _, ok := data.(*AppError); ok {

		err := data.(*AppError)

		if err.Tag == "" {
			err.Tag = "global"
		}

		if err.UserMessage == "" {
			err.UserMessage = "Something went wrong"
		}

		if err.Code == 0 {
			err.Code = 500
		}

		w.WriteHeader(err.Code)

		res.Status = false
		res.Data = nil
		res.Error = err
	} else {
		statusCode := 200

		if len(params) > 0 {
			statusCode = params[0]
		}

		w.WriteHeader(statusCode)

		res.Status = true
		res.Data = data
		res.Error = nil
	}

	json.NewEncoder(w).Encode(res)

}
