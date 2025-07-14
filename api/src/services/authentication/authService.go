package authentication

import (
	"api/src/config"
	"api/src/models"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func GenerateToken(user models.User) (string, error) {
	permission := jwt.MapClaims{}
	permission["authorized"] = true
	permission["exp"] = time.Now().Add(time.Hour * 2).Unix()
	permission["authId"] = user.Id
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permission)
	return token.SignedString([]byte(config.APPkey))
}

func ValidateToken(r *http.Request) error {
	var ErrEmptyToken error = errors.New("unauthorized, missing token")
	var ErrInvalidToken error = errors.New("unauthorized, invalid token")
	tokenStr, err := extractToken(r)
	if err != nil {
		return ErrEmptyToken
	}
	token, err := jwt.Parse(tokenStr, checkKey)
	if err != nil {
		return err
	}
	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}
	return ErrInvalidToken
}

func extractToken(r *http.Request) (string, error) {
	token := r.Header.Get("Authorization")
	var ErrEmptyToken error = errors.New("token missing")
	if strings.TrimSpace(token) == "" {
		return "", ErrEmptyToken
	}
	t := strings.Replace(token, "Bearer ", "", 1)
	return t, nil
}

func checkKey(t *jwt.Token) (interface{}, error) {
	if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("token assign metodo not valid %v", t.Header["alg"])
	}
	return []byte(config.APPkey), nil
}

func AuthId(r *http.Request) (uint, error) {
	var ErrImpossibleParseToken error = errors.New("impossible to parse token")
	var ErrWhenExtractAuthId error = errors.New("error when trying to acess auth id from token")
	tokenStr, err := extractToken(r)
	if err != nil {
		return 0, err
	}
	token, err := jwt.Parse(tokenStr, checkKey)
	if err != nil {
		return 0, err
	}
	tokenData, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, ErrImpossibleParseToken
	}
	id, ok := tokenData["authId"].(float64)
	if !ok || id < 1 {
		return 0, ErrWhenExtractAuthId
	}
	result := uint(id)
	return result, nil
}
