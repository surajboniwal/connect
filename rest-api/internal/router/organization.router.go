package router

import (
	"connect-rest-api/internal/handler"
	"connect-rest-api/internal/util/appauth"

	"github.com/go-chi/chi/v5"
)

func OrganizationRoutes(r *chi.Mux, organizationHandler *handler.OrganizationHandler) {
	r.Route("/organization", func(router chi.Router) {
		router.Use(appauth.AuthMiddleware)
		router.Post("/", organizationHandler.Create)
		router.Get("/", organizationHandler.Get)
	})
}
