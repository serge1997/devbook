package repository

import (
	"api/src/models"
	"errors"
)

func (app App) PersistFollower(follower *models.Follower) error {
	if err := app.db.Create(follower).Error; err != nil {
		return err
	}
	return nil
}

func (app App) FindFollower(id int) (*models.Follower, error) {
	var ErrFollowerNotFound = errors.New("follower not found")
	var follower models.Follower
	if err := app.db.First(&follower, id).Error; err != nil {
		return nil, err
	}
	if follower.Id == 0 {
		return nil, ErrFollowerNotFound
	}
	return &follower, nil
}

func (app App) DeleteFollower(id int) (bool, error) {
	follower, err := app.FindFollower(id)
	if err != nil {
		return false, err
	}
	if err = app.db.Delete(&follower, follower.Id).Error; err != nil {
		return false, err
	}
	return true, nil
}
