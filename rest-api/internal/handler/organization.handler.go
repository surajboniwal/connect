package handler

import (
	"connect-rest-api/internal/model"
	"connect-rest-api/internal/params"
	"connect-rest-api/internal/repository"
	"connect-rest-api/internal/util/appauth"
	"connect-rest-api/internal/util/apphttp"
	"net/http"
)

type OrganizationHandler struct {
	organizationRepo      repository.OrganizationRepository
	organizationUsersRepo repository.OrganizationUsersRepository
}

func NewOrganizationHandler(organizationRepo repository.OrganizationRepository, organizationUsersRepo repository.OrganizationUsersRepository) OrganizationHandler {
	return OrganizationHandler{
		organizationRepo:      organizationRepo,
		organizationUsersRepo: organizationUsersRepo,
	}
}

func (h OrganizationHandler) Create(w http.ResponseWriter, r *http.Request) {
	var params params.CreateOrganization

	userId := appauth.GetUserIdFromContext(r)

	if err := apphttp.ParseAndValidate(r, &params); err != nil {
		apphttp.WriteJSONResponse(w, err)
		return
	}

	organization := &model.Organization{
		Name: params.Name,
	}

	err := h.organizationRepo.Create(organization)

	if err != nil {
		apphttp.WriteJSONResponse(w, err)
		return
	}

	err = h.organizationUsersRepo.AddUser(organization.Id, *userId)

	if err != nil {
		apphttp.WriteJSONResponse(w, err)
		return
	}

	apphttp.WriteJSONResponse(w, organization)
}

func (h OrganizationHandler) Get(w http.ResponseWriter, r *http.Request) {
	userId := *appauth.GetUserIdFromContext(r)

	organizations, err := h.organizationRepo.GetOrganizationsUsingUserId(userId)

	if err != nil {
		apphttp.WriteJSONResponse(w, err)
		return
	}

	apphttp.WriteJSONResponse(w, organizations)
}
