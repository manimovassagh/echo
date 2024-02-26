package db

import (
	"github.com/manimovassagh/echo-learn/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Init_DB() *gorm.DB  {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&models.Product{})
	return db
}
