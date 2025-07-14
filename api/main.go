package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"net/http"
)

func init() {
	config.Load()
	/*key := make([]byte, 64)
	rand.Read(key)
	strKey := base64.StdEncoding.EncodeToString(key)
	fmt.Println(strKey) */
}
func main() {
	fmt.Println("Running package API")
	r := router.Generate()
	port := fmt.Sprintf(":%d", config.Port)
	http.ListenAndServe(port, r)
}
