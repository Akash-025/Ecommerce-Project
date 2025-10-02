package cmd

import (
	"fmt"
	"os"
	"practice/config"
	"practice/infra/db"
	"practice/repo"
	"practice/rest"
	"practice/rest/handlers/product"
	"practice/rest/handlers/user"
	"practice/rest/middleware"
)

func Serve() {

	cnf := config.GetConfig()
	middlewares := middleware.NewMiddleware(cnf)

	dbCon, err := db.NewConnection(&cnf.DB)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = db.MigrateDb(dbCon, "./migrations")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	userRepo := repo.NewUserRepo(dbCon)
	productRepo := repo.NewProductRepo(dbCon)

	productHandler := product.NewHandler(middlewares, productRepo)
	userHandler := user.NewHandler(userRepo, cnf)

	Server := rest.NewServer(productHandler, userHandler, cnf)
	Server.Start()
}
