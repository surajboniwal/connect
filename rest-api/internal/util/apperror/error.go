package apperror

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

func ParseError(err error) *AppError {
	if e, ok := err.(mongo.WriteException); !ok {
		return parseMongoError(e)
	}

	return &AppError{
		OriginalError: err,
	}
}

func parseMongoError(err mongo.WriteException) *AppError {

	match := r.FindStringSubmatch(err.WriteErrors[0].Message)

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

	return &AppError{
		OriginalError: err,
		Tag:           tag,
		UserMessage:   DBErrorMap[index],
		Code:          400,
	}
}
