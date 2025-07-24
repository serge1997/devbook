package middlewares

import (
	"fmt"
	"net/http"
)

func LogRequest(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("%s %s %s \n", r.Method, r.Host, r.URL.Path)
		next(w, r)
	}
}
