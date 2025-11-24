package controller

import (
	"net/http"

	"github.com/serge1997/devbook-web-app/src/utils"
)

func Index(w http.ResponseWriter, r *http.Request) {
	utils.RenderTemplate(w, "home", struct {
		Title string
	}{
		Title: "Data from controller",
	},
	)
}
