package routes

import (
	"api/src/controller"
	"net/http"
)

var authRoutes = []Route{
	{
		Uri:         "/auth",
		Method:      http.MethodPost,
		RequireAuth: false,
		Handle:      controller.Auth,
	},
	{
		Uri:         "/logout",
		Method:      http.MethodPost,
		RequireAuth: false,
		Handle:      controller.LogOut,
	},
}
