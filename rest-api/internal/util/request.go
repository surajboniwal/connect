package util

import (
	"encoding/json"
	"io"
	"net/http"
)

func ParseRequestBody(r *http.Request, data interface{}) *AppError {
	body, err := io.ReadAll(r.Body)

	if err != nil {
		return &AppError{
			OriginalError: err,
			Tag:           "global",
			UserMessage:   "Error while parsing",
			Code:          400,
		}
	}

	if err := r.Body.Close(); err != nil {
		return &AppError{
			OriginalError: err,
			Tag:           "global",
			UserMessage:   "Something went wrong",
			Code:          500,
		}
	}

	if err := json.Unmarshal(body, data); err != nil {
		return &AppError{
			OriginalError: err,
			Tag:           "global",
			UserMessage:   "Error while parsing",
			Code:          400,
		}
	}

	if err := ValidateParam(data); err != nil {
		return err
	}

	return nil
}
