package models

import "time"

type Post struct {
	Id        uint      `json:"id"`
	Title     string    `json:"title" validate:"required"`
	Content   string    `json:"content" validate:"required"`
	AuthorId  uint      `json:"author_id"`
	Likes     uint      `json:"likes"`
	CreatedAt time.Time `json:"created_at"`
	Author    User      `gorm:"foreignKey:AuthorId" json:"author"`
}
