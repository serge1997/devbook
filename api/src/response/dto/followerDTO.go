package dto

import "api/src/models"

type FollowerDTO struct {
	Id     uint   `json:"id"`
	UserId uint   `json:"user_id"`
	Name   string `json:"name"`
}

func FollowerResource(follower *models.Follower) FollowerDTO {
	return FollowerDTO{
		Id:     follower.Id,
		UserId: follower.Follower.Id,
		Name:   follower.Follower.Name,
	}
}
