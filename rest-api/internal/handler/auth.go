package handler

import (
	"connect-rest-api/internal/model"
	"connect-rest-api/internal/params"
	"connect-rest-api/internal/repository"
	"connect-rest-api/internal/util"
	"net/http"
)

type AuthHandler struct {
	userRepo         repository.UserRepository
	organizationRepo repository.OrganizationRepository
}

func NewAuthHandler(userRepo repository.UserRepository, organizationRepository repository.OrganizationRepository) AuthHandler {
	return AuthHandler{
		userRepo:         userRepo,
		organizationRepo: organizationRepository,
	}
}

func (h AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var params params.Register

	if err := util.ParseRequestBody(r, &params); err != nil {
		util.WriteJSONResponse(w, err)
		return
	}

	user, err := h.userRepo.Create(model.User{
		Name:     params.Name,
		Email:    params.Email,
		Password: params.Password,
		Phone:    params.Phone,
	})

	if err != nil {
		util.WriteJSONResponse(w, err)
		return
	}

	_, err = h.organizationRepo.Create(model.Organization{
		Name: params.Organization_Name,
	})

	if err != nil {
		util.WriteJSONResponse(w, err)
		return
	}

	util.WriteJSONResponse(w, user)
}
