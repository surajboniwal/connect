package router

import (
	"connect-rest-api/internal/handler"
	"connect-rest-api/internal/util/appauth"

	"github.com/go-chi/chi/v5"
)

func UserRoutes(r *chi.Mux, userHandler *handler.UserHandler) {
	r.Route("/user", func(router chi.Router) {
		router.Use(appauth.AuthMiddleware)
		router.Post("/", userHandler.Create)
	})
}
