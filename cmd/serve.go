package cmd

import (
	"practice/config"
	"practice/rest"
	"practice/rest/handlers/product"
	"practice/rest/handlers/user"
	"practice/rest/middleware"
)

func Serve() {

	cnf := config.GetConfig()
	middlewares := middleware.NewMiddleware(cnf)

	productHandler := product.NewHandler(middlewares)
	userHandler := user.NewHandler()

	Server := rest.NewServer(productHandler, userHandler, cnf)
	Server.Start()
}
