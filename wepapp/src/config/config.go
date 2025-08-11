package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	API_BASE string = ""
	APP_PORT int
	HashKey  []byte
	BlockKey []byte
)

func Load() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	API_BASE = os.Getenv("API_BASE")

	if API_BASE == "" {
		log.Fatal("API_BASE not informed in .env file")
	}

	HashKey = []byte(os.Getenv("HASH_KEY"))
	BlockKey = []byte(os.Getenv("BLOCK_KEY"))

	APP_PORT, err := strconv.Atoi(os.Getenv("APP_PORT"))
	if err != nil || APP_PORT == 0 {
		log.Fatal("APP PORT not informed")
	}

	fmt.Println(APP_PORT)
}
