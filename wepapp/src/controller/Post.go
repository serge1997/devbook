package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/serge1997/devbook-web-app/src/config"
	"github.com/serge1997/devbook-web-app/src/response"
	"github.com/serge1997/devbook-web-app/src/utils"
)

func GetAllPost(w http.ResponseWriter, r *http.Request) {
	url := fmt.Sprintf("%s/post", config.API_BASE)
	request, err := utils.HttpSend(r, http.MethodGet, url, nil)
	if err != nil {
		response.JSONError(w, err, http.StatusInternalServerError, nil)
		return
	}
	var res response.Response
	defer request.Body.Close()
	json.NewDecoder(request.Body).Decode(&res)
	if res.Code != 200 {
		if res.Code == 0 {
			res.Code = 501
		}
		response.JSONError(w, errors.New(res.Message), res.Code, nil)
		return
	}
	response.JSON(w, res)
}
