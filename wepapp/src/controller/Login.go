package controller

import (
	"encoding/json"
	"net/http"

	"github.com/serge1997/devbook-web-app/src/utils"
)

func LoginIndex(w http.ResponseWriter, r *http.Request) {
	utils.RenderTemplate(w, "login", nil)
}
func RegisterView(w http.ResponseWriter, r *http.Request) {
	utils.RenderTemplate(w, "register", nil)
}

func Register(w http.ResponseWriter, r *http.Request) {
	request, err := http.Post("http://localhost:3000/user", "", r.Body)
	if err != nil {
		w.WriteHeader(501)
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	defer request.Body.Close()
	w.WriteHeader(request.StatusCode)
	json.NewEncoder(w).Encode(request.Body)
}
