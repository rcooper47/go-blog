package main

import (
	"go-blog/inits"
	"go-blog/models"
	"go-blog/repository"
	"go-blog/service"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// add prefix to routes
// Timeouts
// Pagination
// Sorting
// Logging
// Can Read Body in from file in my computer
// Load Testing
// Load Balancing
// Component Testing
// Access Log
// Error Log
// Add Caching
// add mocks
// Add Unit Testing.

type configStrs struct{ db *gorm.DB }

var configStrs1 configStrs

func init() {
	inits.LoadEnv()
	inits.ConnectToDb()
	inits.SyncDB()

	dsn := os.Getenv("DB_URL")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	configStrs1.db = db

	if err != nil {
		panic("Failed to connect to DB")
	}

	configStrs1.db.AutoMigrate(&models.Blog{})
}
func SetupRouter() *gin.Engine {
	r := gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"POST", "GET", "PUT", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Authorization", "Accept", "User-Agent", "Cache-Control", "Pragma"}
	config.ExposeHeaders = []string{"Content-Length"}
	config.AllowCredentials = true
	config.MaxAge = 12 * time.Hour

	postgresBlogRepo := repository.NewPostgresBlogRepo(configStrs1.db)
	blogService := service.NewBlogService(postgresBlogRepo)

	r.Use(cors.New(config))

	r.POST("/blogs", blogService.CreateBlog)
	r.GET("/blogs", blogService.GetAllBlogs)
	r.GET("/blogs/:slug", blogService.GetBlog)
	r.DELETE("/blogs/:slug", blogService.DeleteBlog)

	return r
}
func main() {
	r := SetupRouter()
	r.Run()
}
