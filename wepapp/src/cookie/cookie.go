package cookie

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

func Set(w http.ResponseWriter, Id string, token string) error {
	data := map[string]string{
		"Id":    Id,
		"token": token,
	}
	b, err := json.Marshal(data)
	if err != nil {
		return nil
	}
	encoded := base64.StdEncoding.EncodeToString(b)
	http.SetCookie(w, &http.Cookie{
		Name:     "auth",
		Value:    encoded,
		HttpOnly: true,
		Path:     "/",
		Expires:  time.Now().Add(time.Hour * 6),
	})
	return nil
}
func Get(r *http.Request) (map[string]string, error) {
	var (
		ErrDecoding    = errors.New("erro when tryig to decode cookie data")
		ErrCookieParse = errors.New("erro when trying to parse cookie to a map")
	)
	cookie, err := r.Cookie("auth")
	if err != nil {
		return nil, http.ErrNoCookie
	}
	decode, err := base64.StdEncoding.DecodeString(cookie.Value)
	if err != nil {
		return nil, ErrDecoding
	}

	var data map[string]string
	err = json.Unmarshal(decode, &data)
	if err != nil {
		return nil, ErrCookieParse
	}
	return data, nil
}
func Delete(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:   "auth",
		Value:  "",
		MaxAge: -1,
		Path:   "/",
	})
}
