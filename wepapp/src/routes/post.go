package routes

import (
	"net/http"

	"github.com/serge1997/devbook-web-app/src/controller"
)

var postRoutes = []Route{
	{
		Uri:         "/post",
		Method:      http.MethodGet,
		Handle:      controller.GetAllPost,
		RequireAuth: false,
	},
}
