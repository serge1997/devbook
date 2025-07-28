package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/serge1997/devbook-web-app/src/config"
	"github.com/serge1997/devbook-web-app/src/response"
	"github.com/serge1997/devbook-web-app/src/utils"
)

func LoginIndex(w http.ResponseWriter, r *http.Request) {
	utils.RenderTemplate(w, "login", nil)
}
func RegisterView(w http.ResponseWriter, r *http.Request) {
	utils.RenderTemplate(w, "register", nil)
}

func Register(w http.ResponseWriter, r *http.Request) {
	var form interface{}
	json.NewDecoder(r.Body).Decode(&form)
	b, err := json.Marshal(form)
	if err != nil {
		response.JSONError(w, err, http.StatusInternalServerError, nil)
		return
	}
	req, err := http.Post(fmt.Sprintf("%s/user", config.API_BASE), "application/json", bytes.NewBuffer(b))
	if err != nil {
		response.JSONError(w, err, http.StatusInternalServerError, nil)
		return
	}
	defer req.Body.Close()
	var res response.Response
	json.NewDecoder(req.Body).Decode(&res)
	response.JSON(w, res)
}
