package inits

import (
	"go-blog/models"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDb() {
	dsn := os.Getenv("DB_URL")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	DB = db

	if err != nil {
		panic("Failed to connect to DB")
	}
}

func SyncDB() {
	DB.AutoMigrate(&models.Blog{})
}
