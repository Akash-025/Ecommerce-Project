package rest

import (
	"fmt"
	"net/http"
	"practice/config"
	"practice/rest/middleware"
	"strconv"
)

func Start(cnf config.Config) {

	manager := middleware.NewManager()
	manager.Use(middleware.Preflight, middleware.Cors, middleware.Logger)

	mux := http.NewServeMux()

	InitRoutes(mux, manager)

	wrappedMux := manager.WrapMux(mux)

	//mux.Handle("GET /about", utils.ApplyMiddleware(http.HandlerFunc(aboutHandler), middleware.Hudai, middleware.Logger))

	//secureMux := utils.ApplyMiddleware(mux, middleware.Logger, middleware.Hudai)

	addr := ":" + strconv.Itoa(cnf.HttpPort)

	fmt.Println("Server running on port", addr)

	http.ListenAndServe(addr, wrappedMux)
}
