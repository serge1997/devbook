package routes

import (
	"api/src/middleware"
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Uri         string
	Method      string
	Handle      func(http.ResponseWriter, *http.Request)
	RequireAuth bool
}

func Configuration(r *mux.Router) *mux.Router {
	routes := userRoutes
	routes = append(routes, authRoutes...)
	for _, ur := range routes {
		if ur.RequireAuth {
			r.HandleFunc(ur.Uri,
				middleware.Logger(middleware.TokenMiddleware(ur.Handle)),
			).Methods(ur.Method)
		} else {
			r.HandleFunc(ur.Uri, middleware.Logger(ur.Handle)).Methods(ur.Method)
		}
	}
	return r
}
