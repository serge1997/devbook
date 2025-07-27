package repository

import (
	"api/src/models"
	"errors"

	"gorm.io/gorm"
)

type User struct {
	app                App
	followerRepository *Follower
}

func (user User) GetApp() *App {
	return &user.app
}

func NewUserRepository(app *App, follower *Follower) *User {
	if follower == nil {
		return &User{*app, nil}
	}
	return &User{*app, follower}
}

func (reposiyory User) Persist(user *models.User) (*models.User, error) {
	var ErrNickNameExists = errors.New("user name alwready exists")
	nickname, err := reposiyory.FindByNickName(user.Nick)
	if err != nil {
		return nil, err
	}
	if nickname == nil {
		return nil, ErrNickNameExists
	}
	if err := user.HashPassword(); err != nil {
		return nil, err
	}
	if err := reposiyory.app.db.Create(user).Error; err != nil {
		return nil, err
	}
	if user.Id == 0 {
		err := errors.New("nao foi possivel gravar o usuario")
		return nil, err
	}
	return user, nil
}

func (reposiyory User) FindAll() *[]models.User {
	var users []models.User
	reposiyory.app.db.Find(&users)
	return &users
}

func (reposiyory User) Find(id int) (*models.User, error) {
	var user models.User
	reposiyory.app.db.Preload("Followers.Follower").First(&user, id)
	if user.Id == 0 {
		erro := errors.New("user not found")
		return nil, erro
	}
	return &user, nil
}

func (reposiyory User) Update(user *models.User) (*models.User, error) {
	find, err := reposiyory.Find(int(user.Id))
	if err != nil {
		return nil, err
	}
	if err = reposiyory.app.db.Model(&find).Updates(models.User{Name: user.Name, Email: user.Email, Nick: user.Nick}).Error; err != nil {
		return nil, err
	}
	return find, nil
}

func (reposiyory User) Delete(id int) error {
	find, err := reposiyory.Find(id)
	if err != nil {
		return err
	}
	if err := reposiyory.app.db.Delete(&find, find.Id).Error; err != nil {
		return err
	}
	return nil
}

func (reposiyory User) FindByNickName(nickname string) (*models.User, error) {
	var user models.User
	if err := reposiyory.app.db.Where("nick = ?", nickname).First(&user).Error; !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return &user, nil
}

func (reposiyory User) Follow(follower *models.Follower) error {
	if err := reposiyory.followerRepository.Persist(follower); err != nil {
		return err
	}
	return nil
}

func (reposiyory User) Unfollow(follow_id int) error {
	_, err := reposiyory.followerRepository.Delete(follow_id)
	if err != nil {
		return err
	}
	return nil
}
