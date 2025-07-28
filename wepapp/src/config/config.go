package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	API_BASE string = ""
)

func Load() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	API_BASE = os.Getenv("API_BASE")

	if API_BASE == "" {
		log.Fatal("API_BASE not informed in .env file")
	}

}
