package routes

import (
	"net/http"

	"github.com/serge1997/devbook-web-app/src/controller"
)

var authRoutes = []Route{
	{
		Uri:         "/login",
		Method:      http.MethodGet,
		Handle:      controller.LoginIndex,
		RequireAuth: false,
	},
	{
		Uri:         "/login",
		Method:      http.MethodPost,
		Handle:      controller.Login,
		RequireAuth: false,
	},
	{
		Uri:         "/register",
		Method:      http.MethodGet,
		Handle:      controller.RegisterView,
		RequireAuth: false,
	},
	{
		Uri:         "/register",
		Method:      http.MethodPost,
		Handle:      controller.Register,
		RequireAuth: false,
	},
}
