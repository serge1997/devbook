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
	fmt.Println("Wep app running ")
	r := router.Generate()
	log.Fatal(http.ListenAndServe(":4000", r))
}
