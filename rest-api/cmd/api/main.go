package main

import (
	"connect-rest-api/internal/config"
	"connect-rest-api/internal/database"
	"connect-rest-api/internal/handler"
	"connect-rest-api/internal/repository"
	"connect-rest-api/internal/router"
	"connect-rest-api/internal/util/applogger"
	"flag"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

var env *string = flag.String("env", "development", "App environment - |development|production|")

func main() {

	flag.Parse()

	config := config.Load(*env)
	db := database.Connect(&config)
	r := chi.NewRouter()
	fmt.Println(*env)
	applogger.Init(*env)

	r.Use(applogger.AppLoggerMiddleware)

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
