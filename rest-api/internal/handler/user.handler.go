package handler

import (
	"connect-rest-api/internal/model"
	"connect-rest-api/internal/params"
	"connect-rest-api/internal/repository"
	"connect-rest-api/internal/util/appauth"
	"connect-rest-api/internal/util/apphttp"
	"net/http"
)

type UserHandler struct {
	userRepo repository.UserRepository
}

func NewUserHandler(userRepo repository.UserRepository) UserHandler {
	return UserHandler{
		userRepo: userRepo,
	}
}

func (h UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var params params.CreateUser

	if err := apphttp.ParseAndValidate(r, &params); err != nil {
		apphttp.WriteJSONResponse(w, err)
		return
	}

	user := &model.User{
		Name:     params.Name,
		Email:    params.Email,
		Password: params.Password,
		Phone:    params.Phone,
	}

	err := h.userRepo.Create(user)

	if err != nil {
		apphttp.WriteJSONResponse(w, err)
		return
	}

	token, err := appauth.Generate(user.Id)

	if err != nil {
		apphttp.WriteJSONResponse(w, err)
		return
	}

	apphttp.WriteJSONResponse(w, map[string]string{
		"token": token,
	})
}
