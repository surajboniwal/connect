package appauth

import (
	"connect-rest-api/internal/util/apperror"
	"fmt"
	"strconv"
	"time"

	"github.com/o1egl/paseto"
)

var secretKey string

func Init(key string) {
	secretKey = key
}

func Generate(userid int64) (string, *apperror.AppError) {
	now := time.Now()
	exp := now.Add(24 * time.Hour)
	nbt := now
	t := paseto.NewV2()

	jsonToken := paseto.JSONToken{
		Issuer:     "connect-rest-api",
		IssuedAt:   now,
		Expiration: exp,
		NotBefore:  nbt,
	}

	jsonToken.Set("user_id", fmt.Sprintf("%v", userid))

	token, err := t.Encrypt([]byte(secretKey), jsonToken, nil)

	if err != nil {
		return "", apperror.Parse(err)
	}

	return token, nil
}

func Validate(token string) (int64, *apperror.AppError) {
	var newJsonToken paseto.JSONToken
	t := paseto.NewV2()
	err := t.Decrypt(token, []byte(secretKey), &newJsonToken, nil)

	if err != nil {
		return 0, apperror.Parse(err)
	}

	i, err := strconv.ParseInt(newJsonToken.Get("user_id"), 10, 64)
	if err != nil {
		return 0, apperror.Parse(err)
	}

	return i, nil
}
