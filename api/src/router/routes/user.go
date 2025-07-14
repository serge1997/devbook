package routes

import (
	"api/src/controller"
	"net/http"
)

var userRoutes = []Route{
	{
		Uri:         "/user",
		Method:      http.MethodPost,
		Handle:      controller.StoreUser,
		RequireAuth: false,
	},
	{
		Uri:         "/user",
		Method:      http.MethodGet,
		Handle:      controller.GetUser,
		RequireAuth: true,
	},
	{
		Uri:         "/user/{id}",
		Method:      http.MethodGet,
		Handle:      controller.ShowUser,
		RequireAuth: false,
	},
	{
		Uri:         "/user",
		Method:      http.MethodPut,
		Handle:      controller.UpdateUser,
		RequireAuth: false,
	},
	{
		Uri:         "/user/{id}",
		Method:      http.MethodDelete,
		Handle:      controller.DeleteUser,
		RequireAuth: false,
	},
	{
		Uri:         "/user/{followed_id}/follow",
		Handle:      controller.FolloweUser,
		Method:      http.MethodPost,
		RequireAuth: true,
	},
	{
		Uri:         "/user/{follow_id}/unfollow",
		Handle:      controller.UnFolloweUser,
		Method:      http.MethodDelete,
		RequireAuth: true,
	},
}
