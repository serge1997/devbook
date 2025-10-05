package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/serge1997/devbook-web-app/src/config"
	"github.com/serge1997/devbook-web-app/src/router"
	"github.com/serge1997/devbook-web-app/src/utils"
)

func init() {
	utils.LoadTemplate()
	config.Load()
}
func main() {
	fmt.Println("Wep app running on port: ", config.APP_PORT)
	r := router.Generate()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.APP_PORT), r))
}
