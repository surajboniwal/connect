package handler

import (
	"connect-rest-api/internal/repository"
)

type AuthHandler struct {
	authRepo repository.AuthRepository
}

func NewAuthHandler(authRepo repository.AuthRepository) AuthHandler {
	return AuthHandler{
		authRepo: authRepo,
	}
}
