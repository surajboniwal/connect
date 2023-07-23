package router

import (
	"connect-rest-api/internal/handler"

	"github.com/go-chi/chi/v5"
)

func OrganizationRoutes(r *chi.Mux, organizationHandler *handler.OrganizationHandler) {
	r.Route("/organization", func(router chi.Router) {
		router.Post("/", organizationHandler.Create)
	})
}
