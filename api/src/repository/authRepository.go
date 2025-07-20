package repository

import (
	"api/src/models"
	"api/src/services/authentication"
	"api/src/services/security"
	"errors"
)

type Auth struct {
	app App
}

func (auth Auth) GetApp() *App {
	return &auth.app
}
func NewAuthRepository(app *App) *Auth {
	return &Auth{app: *app}
}

func (auth Auth) Login(user models.User, userRepository User) (*models.User, *string, error) {
	find, err := userRepository.FindByNickName(user.Nick)
	if err != nil {
		return nil, nil, errors.New("usuario  incorreto")
	}
	err = security.Check(find.Password, user.Password)
	if err != nil {
		return nil, nil, errors.New("senha invalida")
	}
	token, err := authentication.GenerateToken(*find)
	if err != nil {
		return nil, nil, errors.New("erro ao gerar o token")
	}
	return find, &token, nil
}
