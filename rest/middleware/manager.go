package middleware

import "net/http"

type Middlware func(http.Handler) http.Handler

type Manager struct{
	globalMiddlewares []Middlware
}

// Constructor
func NewManager() *Manager{
	return &Manager{
		globalMiddlewares: make([]Middlware, 0),
	}
}

func(mngr *Manager) Use(middlewares ...Middlware){
	mngr.globalMiddlewares = append(mngr.globalMiddlewares, middlewares...)
}

func(mngr *Manager) With(next http.Handler, middlewares ...Middlware) http.Handler{
	n := next

	for _, middleware := range middlewares{
		n = middleware(n)
	}
	return n
}

func(mngr *Manager) WrapMux(next http.Handler) http.Handler{
	n := next

	for _, globalMiddleware := range mngr.globalMiddlewares{
		n = globalMiddleware(n)
	}
	return n
}