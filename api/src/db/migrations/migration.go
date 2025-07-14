package main

import (
	"api/src/db/migrations/mconfig"
	"api/src/models"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func init() {
	mconfig.Load()
}

func main() {
	runMigrations()
}
func runMigrations() {
	db, err := db()
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Follower{})
	db.AutoMigrate(&models.Post{})
}

func db() (*gorm.DB, error) {
	fmt.Println("Printing cponfig")
	fmt.Println(mconfig.ConnectionStr)
	db, err := gorm.Open(mysql.Open(mconfig.ConnectionStr), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}
	return db, nil
}
