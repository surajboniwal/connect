package idgen

import "connect-rest-api/internal/util/applogger"

type IdGen interface {
	New() int64
}

var logger = applogger.New("idgen")
