package database

import (
	"connect-rest-api/internal/config"
	"connect-rest-api/internal/util/applogger"
)

type Database interface {
	Connect(config.Config)
}

var logger applogger.Logger = applogger.New("database")
