package controller

import (
	"api/src/db"
	"api/src/models"
	"api/src/repository"
	"api/src/response"
	"api/src/response/dto"
	"encoding/json"
	"net/http"
	"strconv"
)

func Auth(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)
	db, err := db.DB()
	if err != nil {
		response.JSONError(w, http.StatusInternalServerError, err, nil)
	}
	app := repository.New(db)
	userRepository := repository.NewUserRepository(app, nil)
	repository := repository.NewAuthRepository(app)
	defer repository.GetApp().Close()
	defer userRepository.GetApp().Close()
	finded, token, err := repository.Login(user, *userRepository)
	if err != nil {
		response.JSONError(w, http.StatusUnauthorized, err, nil)
		return
	}
	data := dto.UserResource(finded)
	strId := strconv.Itoa(int(data.Id))
	res := map[string]string{
		"id":    strId,
		"name":  data.Name,
		"email": data.Email,
		"token": *token,
	}
	response.JSONSuccess(w, 200, "authentication completed", res)
}

func LogOut(w http.ResponseWriter, r *http.Request) {
	response.JSONSuccess(w, 501, "Method not implemented", nil)
}
