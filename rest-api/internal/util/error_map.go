package util

var DBErrorMap = map[string]map[string]string{
	"email": {
		"unique": "Oops! That email is already in use. Please try a different one.",
	},
}
