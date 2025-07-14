package mconfig

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	ConnectionStr = ""
	Port          = 0
	APPkey        = ""
)

func Load() {
	var erro error

	if erro = godotenv.Load(); erro != nil {
		log.Fatal(erro)
	}
	Port, erro = strconv.Atoi(os.Getenv("APP_PORT"))
	if erro != nil {
		Port = 3000
	}
	APPkey = os.Getenv("APP_KEY")
	ConnectionStr = fmt.Sprintf("%s:@/%s?charset=utf8mb4&parseTime=True&loc=Local", os.Getenv("DB_USER"), os.Getenv("DB_NAME"))
}
