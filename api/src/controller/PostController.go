package controller

import (
	"api/src/db"
	"api/src/models"
	"api/src/repository"
	"api/src/response"
	"api/src/response/dto"
	"api/src/services/authentication"
	"api/src/services/validation"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func StorePost(w http.ResponseWriter, r *http.Request) {
	var post models.Post
	json.NewDecoder(r.Body).Decode(&post)
	validate := validation.NewFormValidation(post)
	validation.Validate(post, validate)
	if validate.HasFailed() {
		response.JSONError(w, 422, nil, validate.GetErrors())
		return
	}
	db, err := db.DB()
	if err != nil {
		response.JSONError(w, http.StatusInternalServerError, err, nil)
		return
	}
	authId, err := authentication.AuthId(r)
	if err != nil {
		response.JSONError(w, http.StatusUnauthorized, err, "")
		return
	}
	app := repository.New(db)
	repository := repository.NewPostRepository(app)
	defer repository.GetApp().Close()
	post.AuthorId = authId
	result, err := repository.Persist(&post)
	if err != nil {
		response.JSONError(w, http.StatusInternalServerError, err, nil)
		return
	}
	dtoResult := dto.PostResource(result)
	response.JSONSuccess(w, http.StatusCreated, "post created successfully", dtoResult)
}
func ListPostByUser(w http.ResponseWriter, r *http.Request) {

}
func ShowPost(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	id, _ := strconv.Atoi(param["id"])
	db, err := db.DB()
	if err != nil {
		response.JSONError(w, http.StatusInternalServerError, err, nil)
		return
	}
	app := repository.New(db)
	repository := repository.NewPostRepository(app)
	post, err := repository.Find(uint(id))
	if err != nil {
		response.JSONError(w, http.StatusNotFound, err, nil)
		return
	}
	resp := dto.PostResource(post)
	response.JSONSuccess(w, http.StatusCreated, "showing post", resp)
}
func GetAllPost(w http.ResponseWriter, r *http.Request) {
	db, err := db.DB()
	if err != nil {
		response.JSONError(w, http.StatusInternalServerError, err, nil)
		return
	}
	app := repository.New(db)
	repository := repository.NewPostRepository(app)
	posts, err := repository.FindAll()
	if err != nil {
		response.JSONError(w, http.StatusNotFound, err, nil)
		return
	}
	resp := dto.PostCollection(posts)
	response.JSONSuccess(w, http.StatusCreated, "showing post", resp)
}
func UpdatePost(w http.ResponseWriter, r *http.Request) {

}
func DeletePost(w http.ResponseWriter, r *http.Request) {

}
