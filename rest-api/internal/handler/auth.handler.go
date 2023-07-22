package handler

import (
	"connect-rest-api/internal/model"
	"connect-rest-api/internal/params"
	"connect-rest-api/internal/repository"
	"connect-rest-api/internal/util/apphttp"
	"net/http"
)

type AuthHandler struct {
	userRepo              repository.UserRepository
	organizationRepo      repository.OrganizationRepository
	organizationUsersRepo repository.OrganizationUsersRepository
}

func NewAuthHandler(userRepo repository.UserRepository, organizationRepo repository.OrganizationRepository, organizationUsersRepo repository.OrganizationUsersRepository) AuthHandler {
	return AuthHandler{
		userRepo:              userRepo,
		organizationRepo:      organizationRepo,
		organizationUsersRepo: organizationUsersRepo,
	}
}

func (h AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var params params.Register

	if err := apphttp.ParseRequestBody(r, &params); err != nil {
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

	err := h.userRepo.Create(user)

	if err != nil {
		apphttp.WriteJSONResponse(w, err)
		return
	}

	err = h.organizationRepo.Create(organization)

	if err != nil {
		apphttp.WriteJSONResponse(w, err)
		return
	}

	err = h.organizationUsersRepo.AddUser(organization.Id, user.Id)

	if err != nil {
		apphttp.WriteJSONResponse(w, err)
		return
	}

	apphttp.WriteJSONResponse(w, user)
}
