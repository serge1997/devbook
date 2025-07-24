package routes

import (
	"net/http"

	"github.com/serge1997/devbook-web-app/src/controller"
)

var homeRoute = Route{
	Uri:         "/",
	Method:      http.MethodGet,
	Handle:      controller.Index,
	RequireAuth: false,
}
