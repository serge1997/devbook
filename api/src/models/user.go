package models

import (
	"api/src/services/security"
	"errors"
	"strings"
	"time"
)

type User struct {
	Id        uint       `json:"id" gorm:"primaryKey"`
	Name      string     `json:"name"`
	Nick      string     `json:"nick"`
	Email     string     `json:"email"`
	Password  string     `json:"password"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	Followers []Follower `gorm:"foreignKey:UserId" json:"followers"`
}

func (user *User) Validate() (erro error) {

	if strings.TrimSpace(user.Name) == "" {
		erro = errors.New("nome é obrigatorio")
		return erro
	}
	if strings.TrimSpace(user.Email) == "" {
		erro = errors.New("emal é obrigatorio")
		return erro
	}
	if strings.TrimSpace(user.Nick) == "" {
		erro = errors.New("nome de usuario é obrigatorio")
		return erro
	}
	if strings.TrimSpace(user.Password) == "" {
		erro = errors.New("senha é obrigatorio")
		return erro
	}
	return nil
}

func (user *User) HashPassword() error {
	hashed, err := security.Hash(user.Password)
	if err != nil {
		return nil
	}
	user.Password = string(hashed)
	return nil
}
