package main

import (
	"go-blog/inits"
	"go-blog/models"
)

// Learn about DB Migrations
func init() {
	inits.LoadEnv()
	inits.ConnectToDb()
}

func main() {
	inits.DB.AutoMigrate(&models.Blog{})
}
