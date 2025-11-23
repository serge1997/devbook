package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/serge1997/devbook-web-app/src/middlewares"
)

type Route struct {
	Uri         string
	Method      string
	Handle      func(http.ResponseWriter, *http.Request)
	RequireAuth bool
}

func Configuration(r *mux.Router) *mux.Router {
	routes := []Route{}
	routes = append(routes, authRoutes...)
	routes = append(routes, homeRoute)
	routes = append(routes, postRoutes...)
	for _, route := range routes {
		r.HandleFunc(route.Uri, middlewares.LogRequest(middlewares.JWTMiddleware(route.Handle))).Methods(route.Method)
	}
	fileServer := http.FileServer(http.Dir("./assets/"))
	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fileServer))
	return r
}
