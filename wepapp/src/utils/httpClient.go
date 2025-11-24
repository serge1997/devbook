package utils

import (
	"fmt"
	"io"
	"net/http"

	"github.com/serge1997/devbook-web-app/src/cookie"
)

func HttpSend(r *http.Request, method, url string, payload io.Reader) (*http.Response, error) {
	request, err := http.NewRequest(method, url, payload)
	if err != nil {
		return nil, err
	}
	c, err := cookie.Get(r)
	if err != nil {
		return nil, err
	}
	token := fmt.Sprintf("Bearer %s", c["token"])
	request.Header.Set("Authorization", token)
	client := &http.Client{}
	response, erro := client.Do(request)
	if erro != nil {
		return nil, erro
	}
	return response, nil
}
