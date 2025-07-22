package repository

import (
	"api/src/models"
	"errors"
	"time"
)

type Post struct {
	app App
}

func (post Post) GetApp() *App {
	return &post.app
}
func NewPostRepository(app *App) *Post {
	return &Post{*app}
}

func (repository Post) Persist(post *models.Post) (*models.Post, error) {
	var ErrUnknowPersitUser error = errors.New("some erro ocurred, please try leater")
	post.CreatedAt = time.Now()
	if err := repository.app.db.Create(post).Error; err != nil {
		return nil, err
	}
	if post.Id == 0 {
		return nil, ErrUnknowPersitUser
	}
	return post, nil
}
func (repository Post) FindAllByAuthor(authorId int) (*[]models.Post, error) {
	var posts []models.Post
	if err := repository.app.db.Where("AuthorId = ?", authorId).Find(&posts).Error; err != nil {
		return nil, err
	}
	return &posts, nil
}
func (repository Post) Find(id uint) (*models.Post, error) {
	var post models.Post
	var ErrNotFound = errors.New("post not found")
	if err := repository.app.db.Preload("Author").First(&post, id).Error; err != nil {
		return nil, err
	}
	if post.Id == 0 {
		return nil, ErrNotFound
	}
	return &post, nil
}

func (repository Post) Update(post *models.Post) (*models.Post, error) {
	find, err := repository.Find(post.Id)
	if err != nil {
		return nil, err
	}
	err = repository.app.db.Model(find).Updates(models.Post{Title: post.Title, Content: post.Content}).Error
	if err != nil {
		return nil, err
	}
	return find, nil
}

func (repository Post) Delete(id uint) error {
	find, err := repository.Find(id)
	if err != nil {
		return err
	}
	err = repository.app.db.Delete(find, find.Id).Error
	if err != nil {
		return err
	}
	return nil
}
