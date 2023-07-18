package main

import (
	"connect-rest-api/internal/config"
	"connect-rest-api/internal/database"
	"connect-rest-api/internal/handler"
	"connect-rest-api/internal/repository"
	"connect-rest-api/internal/router"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	config := config.Load()
	db := database.Connect(&config)
	r := chi.NewRouter()

	organizationRepo := repository.NewOrganizationRepositoryMongo(&db)
	userRepo := repository.NewUserRepositoryMongo(&db)

	organizationHandler := handler.NewOrganizationHandler(&organizationRepo)
	userHandler := handler.NewUserHandler(&userRepo)
	authHandler := handler.NewAuthHandler(&userRepo, &organizationRepo)

	router.OrganizationRoutes(r, &organizationHandler)
	router.UserRoutes(r, &userHandler)
	router.AuthRoutes(r, &authHandler)

	http.ListenAndServe(fmt.Sprintf(":%v", config.PORT), r)
}
