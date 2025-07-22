package controller

import (
	"api/src/models"
	"api/src/response"
	"api/src/services/validation"
	"encoding/json"
	"net/http"
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
}
func ListPostByUser(w http.ResponseWriter, r *http.Request) {

}
func ShowPost(w http.ResponseWriter, r *http.Request) {

}
func UpdatePost(w http.ResponseWriter, r *http.Request) {

}
func DeletePost(w http.ResponseWriter, r *http.Request) {

}
