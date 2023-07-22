package main

import (
	"connect-rest-api/internal/config"
	"connect-rest-api/internal/database"
	"connect-rest-api/internal/handler"
	"connect-rest-api/internal/repository"
	"connect-rest-api/internal/router"
	"connect-rest-api/internal/util/applogger"
	"connect-rest-api/internal/util/idgen"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	config := config.Load()

	database := database.NewPgDatabase(&config)
	database.Connect()

	r := chi.NewRouter()

	r.Use(applogger.AppLoggerMiddleware)

	idGen := idgen.NewSnowflakeIdGen()

	organizationRepo := repository.NewOrganizationRepositoryPg(database.DB, idGen)
	userRepo := repository.NewUserRepositoryPg(database.DB, idGen)

	organizationHandler := handler.NewOrganizationHandler(&organizationRepo)
	userHandler := handler.NewUserHandler(&userRepo)
	authHandler := handler.NewAuthHandler(&userRepo, &organizationRepo)

	router.OrganizationRoutes(r, &organizationHandler)
	router.UserRoutes(r, &userHandler)
	router.AuthRoutes(r, &authHandler)

	http.ListenAndServe(fmt.Sprintf(":%v", config.PORT), r)
}
