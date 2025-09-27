package utils

import "net/http"

type Middlware func(http.Handler) http.Handler

func ApplyMiddleware(mainHandler http.Handler, middlewares ...Middlware)http.Handler{
	for _, middleware := range middlewares{
		mainHandler = middleware(mainHandler)
	}
	return mainHandler
}
