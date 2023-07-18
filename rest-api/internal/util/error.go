package util

type AppError struct {
	OriginalError error  `json:"-"`
	Tag           string `json:"tag"`
	UserMessage   string `json:"message"`
	Code          int    `json:"-"`
}

func (m *AppError) Error() string {
	return m.UserMessage
}
