package handler

import "connect-rest-api/internal/repository"

type UserHandler struct {
	userRepo repository.UserRepository
}

func NewUserHandler(userRepo repository.UserRepository) UserHandler {
	return UserHandler{
		userRepo: userRepo,
	}
}
