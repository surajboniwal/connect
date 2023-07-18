package util

import (
	"regexp"
	"strings"

	"go.mongodb.org/mongo-driver/mongo"
)

type AppError struct {
	OriginalError error  `json:"-"`
	Tag           string `json:"tag"`
	UserMessage   string `json:"message"`
	Code          int    `json:"-"`
}

func (m *AppError) Error() string {
	return m.UserMessage
}

var r = regexp.MustCompile(`index:\s(\S+)`)

func ParseMongoError(err error) *AppError {
	var mongoError mongo.WriteException

	if e, ok := err.(mongo.WriteException); !ok {
		return &AppError{
			OriginalError: err,
		}
	} else {
		mongoError = e
	}

	match := r.FindStringSubmatch(mongoError.WriteErrors[0].Message)

	if len(match) <= 1 {
		return &AppError{
			OriginalError: err,
		}
	}

	index := match[1]

	splitIndex := strings.Split(index, "_")

	if len(splitIndex) < 2 {
		return &AppError{
			OriginalError: err,
		}
	}

	tag := splitIndex[0]
	errorType := splitIndex[1]

	return &AppError{
		OriginalError: err,
		Tag:           tag,
		UserMessage:   DBErrorMap[tag][errorType],
		Code:          400,
	}
}
