package rest

import (
	"net/http"
	"practice/rest/handlers"
	"practice/rest/middleware"
)

func InitRoutes(mux *http.ServeMux, manager *middleware.Manager) {
	mux.Handle(
		"GET /products/",
		manager.With(http.HandlerFunc(handlers.GetProducts)),
	)
	mux.Handle("GET /products/{id}", manager.With(http.HandlerFunc(handlers.GetProduct)))
	mux.Handle("POST /products/", manager.With(http.HandlerFunc(handlers.CreateProduct)))
	mux.Handle("PUT /products/{id}", manager.With(http.HandlerFunc(handlers.UpdateProduct)))
	mux.Handle("DELETE /products/{id}", manager.With(http.HandlerFunc(handlers.DeleteProduct)))

	mux.Handle("POST /users/", manager.With(http.HandlerFunc(handlers.CreateUser)))
	mux.Handle("POST /users/login", manager.With(http.HandlerFunc(handlers.LoginUser)))

}
