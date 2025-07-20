package repository

import (
	"api/src/models"
	"errors"
)

type Follower struct {
	app App
}

func NewFollowerRepository(app *App) *Follower {
	return &Follower{*app}
}
func (follower Follower) GetApp() *App {
	return &follower.app
}

func (repository Follower) Persist(follower *models.Follower) error {
	if err := repository.app.db.Create(follower).Error; err != nil {
		return err
	}
	return nil
}

func (repository Follower) Find(id int) (*models.Follower, error) {
	var ErrFollowerNotFound = errors.New("follower not found")
	var follower models.Follower
	if err := repository.app.db.First(&follower, id).Error; err != nil {
		return nil, err
	}
	if follower.Id == 0 {
		return nil, ErrFollowerNotFound
	}
	return &follower, nil
}

func (repository Follower) Delete(id int) (bool, error) {
	follower, err := repository.Find(id)
	if err != nil {
		return false, err
	}
	if err = repository.app.db.Delete(&follower, follower.Id).Error; err != nil {
		return false, err
	}
	return true, nil
}
