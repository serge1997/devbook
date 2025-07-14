package repository

import (
	"api/src/models"
	"api/src/services/authentication"
	"api/src/services/security"
	"errors"
)

func (app App) Login(user models.User) (*models.User, *string, error) {
	find, err := app.FindUserByNickName(user.Nick)
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
