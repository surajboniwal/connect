package applogger

import (
	"fmt"
	"net/http"
)

type Logger interface {
	I(any, ...string)
	E(any, ...string)
	D(any, ...string)
}

var env string

func Init(_env string) {
	env = _env
}

func New(name string) Logger {
	fmt.Printf("%v %v\n", name, env)
	switch env {
	case "development":
		return newConsoleLogger(name)
	case "production":
		return newSentryLogger(name)
	default:
		return newConsoleLogger(name)
	}
}

func AppLoggerMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if env != "development" {
			h.ServeHTTP(w, r)
		} else {
			logger := New("REQUEST")
			logger.I(r.RequestURI)
			h.ServeHTTP(w, r)
		}
	})
}
