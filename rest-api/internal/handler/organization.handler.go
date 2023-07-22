package handler

import "connect-rest-api/internal/repository"

type OrganizationHandler struct {
	organizationRepo repository.OrganizationRepository
}

func NewOrganizationHandler(organizationRepo repository.OrganizationRepository) OrganizationHandler {
	return OrganizationHandler{
		organizationRepo: organizationRepo,
	}
}
