package appenv

import "connect-rest-api/internal/util"

func AppEnv() string {
	return util.Getenv("ENV", "development")
}
