package repository

import (
	"gorm.io/gorm"
)

type App struct {
	db *gorm.DB
}

func New(db *gorm.DB) *App {
	return &App{db}
}

func (orm *App) Close() {
	db, _ := orm.db.DB()
	db.Close()
}
