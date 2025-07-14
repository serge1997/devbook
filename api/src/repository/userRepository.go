package repository

import (
	"api/src/models"
	"errors"
)

func (app App) PersistUser(user *models.User) (*models.User, error) {
	if err := user.HashPassword(); err != nil {
		return nil, err
	}
	if err := app.db.Create(user).Error; err != nil {
		return nil, err
	}
	if user.Id == 0 {
		err := errors.New("nao foi possivel gravar o usuario")
		return nil, err
	}
	return user, nil
}

func (app App) FindAllUser() *[]models.User {
	var users []models.User
	app.db.Find(&users)
	return &users
}

func (app App) FindUser(id int) (*models.User, error) {
	var user models.User
	app.db.Preload("Followers.Follower").First(&user, id)
	if user.Id == 0 {
		erro := errors.New("user not found")
		return nil, erro
	}
	return &user, nil
}

func (app App) UpdateUser(user *models.User) (*models.User, error) {
	find, err := app.FindUser(int(user.Id))
	if err != nil {
		return nil, err
	}
	if err = app.db.Model(&find).Updates(models.User{Name: user.Name, Email: user.Email, Nick: user.Nick}).Error; err != nil {
		return nil, err
	}
	return find, nil
}

func (app App) DeleteUser(id int) error {
	find, err := app.FindUser(id)
	if err != nil {
		return err
	}
	if err := app.db.Delete(&find, find.Id).Error; err != nil {
		return err
	}
	return nil
}

func (model App) FindUserByNickName(nickname string) (*models.User, error) {
	var user models.User
	if err := model.db.Where("nick = ?", nickname).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (app App) FollowUser(follower *models.Follower) error {
	if err := app.PersistFollower(follower); err != nil {
		return err
	}
	return nil
}

func (app App) UnfollowUser(follow_id int) error {
	_, err := app.DeleteFollower(follow_id)
	if err != nil {
		return err
	}
	return nil
}
