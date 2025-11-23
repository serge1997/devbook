package middlewares

import (
	"fmt"
	"net/http"

	"github.com/serge1997/devbook-web-app/src/cookie"
	"github.com/serge1997/devbook-web-app/src/utils"
)

func JWTMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c, err := cookie.Get(r)
		if err != nil && r.URL.Path != "/login" {
			fmt.Println("error: ", err)
			utils.RenderTemplate(w, "login", nil)
			return
		}
		var token string = fmt.Sprintf("Bearer %s", c["token"])
		r.Header.Set("Authorization", token)
		next(w, r)
	}
}
