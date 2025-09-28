package cmd

import (
	"practice/config"
	"practice/repo"
	"practice/rest"
	"practice/rest/handlers/product"
	"practice/rest/handlers/user"
	"practice/rest/middleware"
)

func Serve() {

	cnf := config.GetConfig()
	middlewares := middleware.NewMiddleware(cnf)

	userRepo := repo.NewUserRepo()
	productRepo := repo.NewProductRepo()

	productHandler := product.NewHandler(middlewares, productRepo)
	userHandler := user.NewHandler(userRepo, cnf)

	Server := rest.NewServer(productHandler, userHandler, cnf)
	Server.Start()
}
