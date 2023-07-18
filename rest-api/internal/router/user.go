package router

import (
	"connect-rest-api/internal/handler"

	"github.com/go-chi/chi/v5"
)

func UserRoutes(r *chi.Mux, userHandler *handler.UserHandler) {
	r.Route("/user", func(router chi.Router) {
	})
}
