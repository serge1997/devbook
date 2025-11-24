package controller

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/serge1997/devbook-web-app/src/config"
	"github.com/serge1997/devbook-web-app/src/cookie"
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
	url := fmt.Sprintf("%s/user", config.API_BASE)
	req, err := http.Post(url, "application/json", bytes.NewBuffer(b))
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
	url := fmt.Sprintf("%s/auth", config.API_BASE)
	request, err := http.Post(url, "application/json", bytes.NewBuffer(loginFormToByte))
	if err != nil {
		response.JSONError(w, err, http.StatusInternalServerError, nil)
		return
	}
	defer request.Body.Close()
	var resp response.Response
	json.NewDecoder(request.Body).Decode(&resp)
	if resp.Code != 200 {
		response.JSONError(w, errors.New(resp.Message), http.StatusInternalServerError, nil)
		return
	}
	responseMapa, _ := resp.Data.(map[string]any)
	id, assertedId := responseMapa["id"].(string)
	token, assertedToken := responseMapa["token"].(string)
	if !assertedId || !assertedToken {
		response.JSONError(w, errors.New("error when tried to set cookie"), http.StatusInternalServerError, nil)
		return
	}
	cookie.Set(w, id, token)
	response.JSON(w, resp)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	cookie.Delete(w)
	var res response.Response
	res.Message = "Logout realizado"
	res.Code = 200
	response.JSON(w, res)
}
