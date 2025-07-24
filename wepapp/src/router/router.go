package router

import (
	"github.com/gorilla/mux"
	"github.com/serge1997/devbook-web-app/src/routes"
)

func Generate() *mux.Router {
	r := mux.NewRouter()
	return routes.Configuration(r)
}
