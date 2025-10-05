package routes

import (
	"api/src/controller"
	"net/http"
)

var PostRoutes = []Route{
	{
		Uri:         "/post",
		Handle:      controller.StorePost,
		Method:      http.MethodPost,
		RequireAuth: true,
	},
	{
		Uri:         "/post",
		Method:      http.MethodGet,
		Handle:      controller.GetAllPost,
		RequireAuth: true,
	},
	{
		Uri:         "/post/{id}",
		Handle:      controller.ShowPost,
		Method:      http.MethodGet,
		RequireAuth: true,
	},
	{
		Uri:         "/post/list-by-user/{user_id}",
		Handle:      controller.ListPostByUser,
		Method:      http.MethodGet,
		RequireAuth: true,
	},
	{
		Uri:         "/post",
		Handle:      controller.UpdatePost,
		Method:      http.MethodPut,
		RequireAuth: true,
	},
	{
		Uri:         "/post/{id}",
		Handle:      controller.DeletePost,
		Method:      http.MethodDelete,
		RequireAuth: true,
	},
}
