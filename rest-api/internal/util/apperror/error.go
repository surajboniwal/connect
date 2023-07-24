package apperror

import (
	"connect-rest-api/internal/util/applogger"

	"github.com/lib/pq"
)

var logger = applogger.New("apperror")

type AppError struct {
	OriginalError error  `json:"-"`
	Tag           string `json:"tag"`
	UserMessage   string `json:"message"`
	Code          int    `json:"-"`
}

func Parse(err error) *AppError {
	logger.E(err)
	if e, ok := err.(*pq.Error); ok {
		return parsePgError(e)
	}

	return &AppError{
		OriginalError: err,
	}
}

func parsePgError(err *pq.Error) *AppError {

	return &AppError{
		OriginalError: err,
		Tag:           err.Column,
		UserMessage:   DBErrorMap[err.Constraint],
		Code:          400,
	}
}
