package controller

import (
	"api/src/db"
	"api/src/models"
	"api/src/repository"
	"api/src/response"
	"api/src/response/dto"
	"api/src/services/authentication"
	"api/src/services/utils"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func StoreUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		response.JSONError(w, http.StatusBadRequest, err, nil)
		return
	}
	if validateErr := user.Validate(); validateErr != nil {
		response.JSONError(w, http.StatusUnprocessableEntity, validateErr, nil)
		return
	}
	err := utils.ValiateEmail(user.Email)
	if err != nil {
		response.JSONError(w, http.StatusBadRequest, err, nil)
		return
	}
	db, err := db.DB()
	if err != nil {
		response.JSONError(w, http.StatusInternalServerError, err, nil)
		return
	}
	app := repository.New(db)
	repository := repository.NewUserRepository(app, nil)
	saved, err := repository.Persist(&user)
	defer repository.GetApp().Close()
	if err != nil {
		response.JSONError(w, http.StatusInternalServerError, err, nil)
		return
	}
	response.JSONSuccess(w, http.StatusCreated, "User registred successfully", saved)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	db, err := db.DB()
	if err != nil {
		response.JSONError(w, http.StatusInternalServerError, err, nil)
		return
	}
	app := repository.New(db)
	repository := repository.NewUserRepository(app, nil)
	users := repository.FindAll()
	defer repository.GetApp().Close()
	msg := "Listando todos os uarios"
	userCollections := dto.UserCollection(users)
	response.JSONSuccess(w, http.StatusOK, msg, userCollections)
}

func ShowUser(w http.ResponseWriter, r *http.Request) {
	db, err := db.DB()
	if err != nil {
		response.JSONError(w, http.StatusInternalServerError, err, nil)
		return
	}
	param := mux.Vars(r)
	id, err := strconv.Atoi(param["id"])
	if err != nil {
		response.JSONError(w, http.StatusBadRequest, err, nil)
		return
	}
	app := repository.New(db)
	repository := repository.NewUserRepository(app, nil)
	user, err := repository.Find(int(id))
	defer repository.GetApp().Close()
	if err != nil {
		response.JSONError(w, http.StatusNotFound, err, nil)
		return
	}
	msg := "mostrando o usuario"
	userResponse := dto.UserResource(user)
	response.JSONSuccess(w, http.StatusOK, msg, userResponse)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	authId, err := authentication.AuthId(r)
	if err != nil {
		response.JSONError(w, http.StatusUnauthorized, err, nil)
		return
	}
	db, err := db.DB()
	if err != nil {
		response.JSONError(w, http.StatusInternalServerError, err, nil)
		return
	}
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)
	app := repository.New(db)
	repository := repository.NewUserRepository(app, nil)
	defer repository.GetApp().Close()

	if authId != user.Id {
		msg := errors.New("permission denied. You cannot update a user")
		response.JSONError(w, http.StatusForbidden, msg, nil)
		return
	}
	err = utils.ValiateEmail(user.Email)
	if err != nil {
		response.JSONError(w, http.StatusUnauthorized, err, nil)
		return
	}

	result, err := repository.Update(&user)
	if err != nil {
		response.JSONError(w, http.StatusNotFound, err, nil)
		return
	}
	response.JSONSuccess(w, 200, "user updated successfully", result)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		response.JSONError(w, http.StatusNotFound, err, nil)
		return
	}
	db, err := db.DB()
	if err != nil {
		response.JSONError(w, http.StatusInternalServerError, err, nil)
		return
	}
	app := repository.New(db)
	repository := repository.NewUserRepository(app, nil)
	err = repository.Delete(id)
	defer repository.GetApp().Close()
	if err != nil {
		response.JSONError(w, http.StatusInternalServerError, err, nil)
		return
	}
	response.JSONSuccess(w, 200, "user deleted successully", nil)
}

func FolloweUser(w http.ResponseWriter, r *http.Request) {
	var follower models.Follower
	params := mux.Vars(r)
	followedId, _ := strconv.Atoi(params["followed_id"])
	fmt.Println(params)
	if followedId == 0 {
		response.JSONError(w, http.StatusBadRequest, errors.New("invalid parameter"), nil)
		return
	}
	authId, err := authentication.AuthId(r)
	if err != nil {
		response.JSONError(w, http.StatusBadRequest, err, nil)
		return
	}
	follower.FollowerId = uint(followedId)
	follower.UserId = authId
	db, err := db.DB()
	if err != nil {
		response.JSONError(w, http.StatusInternalServerError, err, nil)
		return
	}
	app := repository.New(db)
	followerRepository := repository.NewFollowerRepository(app)
	repository := repository.NewUserRepository(app, followerRepository)
	defer repository.GetApp().Close()
	defer followerRepository.GetApp().Close()
	if err = repository.Follow(&follower); err != nil {
		response.JSONError(w, http.StatusBadRequest, err, nil)
		return
	}

	response.JSONSuccess(w, http.StatusOK, "followed", nil)

}
func UnFolloweUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	followedId, _ := strconv.Atoi(params["follow_id"])
	if followedId == 0 {
		response.JSONError(w, http.StatusBadRequest, errors.New("invalid parameter"), nil)
		return
	}
	db, err := db.DB()
	if err != nil {
		response.JSONError(w, http.StatusInternalServerError, err, nil)
		return
	}
	app := repository.New(db)
	followerRepository := repository.NewFollowerRepository(app)
	repository := repository.NewUserRepository(app, followerRepository)
	defer repository.GetApp().Close()
	defer followerRepository.GetApp().Close()
	if err = repository.Unfollow(followedId); err != nil {
		response.JSONError(w, http.StatusBadRequest, err, nil)
		return
	}

	response.JSONSuccess(w, http.StatusOK, "unfollowed", nil)

}
