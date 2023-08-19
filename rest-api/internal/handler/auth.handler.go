package handler

import (
	"connect-rest-api/internal/model"
	"connect-rest-api/internal/params"
	"connect-rest-api/internal/repository"
	"connect-rest-api/internal/util/appauth"
	"connect-rest-api/internal/util/apperror"
	"connect-rest-api/internal/util/apphttp"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

type AuthHandler struct {
	userRepo repository.UserRepository
}

func NewAuthHandler(userRepo repository.UserRepository) AuthHandler {
	return AuthHandler{
		userRepo: userRepo,
	}
}

func (h AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var params params.Login

	if err := apphttp.ParseAndValidate(r, &params); err != nil {
		apphttp.WriteJSONResponse(w, err)
		return
	}

	user, err := h.userRepo.GetByEmail(params.Email)

	if err != nil {
		apphttp.WriteJSONResponse(w, &apperror.UnauthorizedError)
		return
	}

	e := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(params.Password))

	if e != nil {
		apphttp.WriteJSONResponse(w, &apperror.UnauthorizedError)
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

	err := h.userRepo.Create(user)

	if err != nil {
		apphttp.WriteJSONResponse(w, err)
		return
	}

	apphttp.WriteJSONResponse(w, "Registration successful", 201)
}
