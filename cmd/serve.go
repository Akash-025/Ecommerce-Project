package cmd

import (
	"fmt"
	"os"
	"practice/config"
	"practice/infra/db"
	"practice/product"
	"practice/repo"
	"practice/rest"
	prdctHandler "practice/rest/handlers/product"
	usrHandler "practice/rest/handlers/user"
	"practice/rest/middleware"
	"practice/user"
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

	//repos
	userRepo := repo.NewUserRepo(dbCon)
	productRepo := repo.NewProductRepo(dbCon)

	//domains
	userSvc := user.NewService(userRepo)
	productSvc := product.NewService(productRepo)

	//handlers
	productHandler := prdctHandler.NewHandler(middlewares, productSvc)
	userHandler := usrHandler.NewHandler(cnf, userSvc)

	Server := rest.NewServer(productHandler, userHandler, cnf)
	Server.Start()
}
