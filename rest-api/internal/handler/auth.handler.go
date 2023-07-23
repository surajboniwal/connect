package handler

import (
	"connect-rest-api/internal/model"
	"connect-rest-api/internal/params"
	"connect-rest-api/internal/repository"
	"connect-rest-api/internal/util/apphttp"
	"net/http"
)

type AuthHandler struct {
	authRepo repository.AuthRepository
}

func NewAuthHandler(authRepo repository.AuthRepository) AuthHandler {
	return AuthHandler{
		authRepo: authRepo,
	}
}

func (h AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var params params.Register

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

	organization := &model.Organization{
		Name: params.Organization_Name,
	}

	err := h.authRepo.Register(user, organization)

	if err != nil {
		apphttp.WriteJSONResponse(w, err)
		return
	}

	apphttp.WriteJSONResponse(w, user)
}
