package models

import "time"

type Post struct {
	Id        uint      `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	AuthorId  uint      `json:"author_id"`
	Likes     uint      `json:"likes"`
	CreatedAt time.Time `json:"created_at"`
	Author    User      `gorm:"foreignKey:AuthorId" json:"author"`
}
