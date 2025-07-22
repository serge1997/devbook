package dto

import (
	"api/src/models"
	"time"
)

type PostDTO struct {
	Id        uint      `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	AuthorId  uint      `json:"author_id"`
	Likes     uint      `json:"likes"`
	CreatedAt time.Time `json:"created_at"`
	Author    *UserDTO  `json:"author"`
}

func PostResource(post *models.Post) PostDTO {
	postDto := PostDTO{
		Id:        post.Id,
		Title:     post.Title,
		Content:   post.Content,
		CreatedAt: post.CreatedAt,
		AuthorId:  post.AuthorId,
	}
	if post.Author.Id != 0 {
		postDto.Author = &UserDTO{
			Id:   post.Author.Id,
			Name: post.Author.Name,
			Nick: post.Author.Nick,
		}
	}
	return postDto
}

func PostCollection(posts *[]models.Post) []PostDTO {
	var postsDto []PostDTO
	for _, post := range *posts {
		postsDto = append(postsDto, PostDTO{
			Id:        post.Id,
			Title:     post.Title,
			Content:   post.Content,
			CreatedAt: post.CreatedAt,
			Author: func() *UserDTO {
				if post.Author.Id != 0 {
					return &UserDTO{
						Id:   post.Author.Id,
						Name: post.Author.Name,
						Nick: post.Author.Nick,
					}
				}
				return nil
			}(),
		})
	}
	return postsDto
}
