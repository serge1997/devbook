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

func Login(w http.ResponseWriter, r *http.Request) {
	var loginForm interface{}
	json.NewDecoder(r.Body).Decode(&loginForm)
	loginFormToByte, err := json.Marshal(loginForm)
	if err != nil {
		response.JSONError(w, err, http.StatusInternalServerError, nil)
		return
	}
	request, err := http.Post(fmt.Sprintf("%s/auth", config.API_BASE), "application/json", bytes.NewBuffer(loginFormToByte))
	if err != nil {
		response.JSONError(w, err, http.StatusInternalServerError, nil)
		return
	}
	defer request.Body.Close()
	var resp response.Response
	json.NewDecoder(request.Body).Decode(&resp)
	responseMapa, ok := resp.Data.(map[string]interface{})
	if !ok {
		response.JSONError(w, err, http.StatusInternalServerError, nil)
		return
	}
	fmt.Println(responseMapa)
	response.JSON(w, resp)
}
