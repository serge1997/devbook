package dto

import (
	"api/src/models"
)

type UserDTO struct {
	Id        uint          `json:"id"`
	Name      string        `json:"name"`
	Nick      string        `json:"nick"`
	Email     string        `json:"email"`
	CreatedAt string        `json:"created_at"`
	UpdatedAt string        `json:"updated_at"`
	Followers []FollowerDTO `json:"followers"`
}

func UserResource(user *models.User) UserDTO {
	var followers []FollowerDTO
	usersFollowers := user.Followers
	for _, follower := range usersFollowers {
		followers = append(followers, FollowerResource(&follower))
	}
	return UserDTO{
		Id:        user.Id,
		Name:      user.Name,
		Email:     user.Email,
		Nick:      user.Nick,
		CreatedAt: user.CreatedAt.Format("02/01/2006 15:04:15"),
		UpdatedAt: user.UpdatedAt.Format("02/01/2006 15:04:15"),
		Followers: followers,
	}
}

func UserCollection(users *[]models.User) []UserDTO {
	var collections []UserDTO
	for _, user := range *users {
		collections = append(collections, UserDTO{
			Id:        user.Id,
			Name:      user.Name,
			Email:     user.Email,
			Nick:      user.Nick,
			CreatedAt: user.CreatedAt.Format("02/01/2006 15:04:15"),
			UpdatedAt: user.UpdatedAt.Format("02/01/2006 15:04:15"),
		})
	}
	return collections
}
