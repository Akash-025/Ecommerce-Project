package rest

import (
	"fmt"
	"net/http"
	"practice/config"
	"practice/rest/handlers/product"
	"practice/rest/handlers/user"
	"practice/rest/middleware"
	"strconv"
)

type Server struct{
	productHandler *product.Handler
	userHandler *user.Handler
	cnf *config.Config
}

func NewServer(
	productHandler *product.Handler,
	userHandler *user.Handler,
	cnf *config.Config,
	) *Server{
	return &Server{
		productHandler: productHandler,
		userHandler: userHandler,
		cnf: cnf,
	}
}

func (server *Server) Start() {

	manager := middleware.NewManager()
	manager.Use(middleware.Preflight, middleware.Cors, middleware.Logger)

	mux := http.NewServeMux()

	//InitRoutes(mux, manager)
	server.productHandler.RegisterRoutes(mux, manager)
	server.userHandler.RegisterRoutes(mux, manager)

	wrappedMux := manager.WrapMux(mux)

	//mux.Handle("GET /about", utils.ApplyMiddleware(http.HandlerFunc(aboutHandler), middleware.Hudai, middleware.Logger))

	//secureMux := utils.ApplyMiddleware(mux, middleware.Logger, middleware.Hudai)

	addr := ":" + strconv.Itoa(server.cnf.HttpPort)

	fmt.Println("Server running on port", addr)

	http.ListenAndServe(addr, wrappedMux)
}
