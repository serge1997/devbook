package db

import (
	"api/src/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func DB() (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(config.ConnectionStr), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}
	return db, nil
}
