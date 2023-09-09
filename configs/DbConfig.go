package configs

import (
	"idp-go/app/practice/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func GetDb() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("todos.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database!")
	}

	// automigrate
	db.AutoMigrate(&models.ToDo{})

	return db
}
