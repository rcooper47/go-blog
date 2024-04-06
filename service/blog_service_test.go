package service

import (
	"bytes"
	"encoding/json"
	"go-blog/models"
	"go-blog/repository"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type testBlogRepository struct {
	db *gorm.DB
}

func NewTestBlogRepo(db *gorm.DB) *testBlogRepository {
	return &testBlogRepository{
		db: db,
	}
}

// GetBlogs implements BlogRepository.
func (pr *testBlogRepository) GetBlogs(c *gin.Context) (*[]models.Blog, error) {
	var blogs []models.Blog
	if err := pr.db.Find(&blogs).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return nil, err
	}

	return &blogs, nil
}

type configStrs struct{ db *gorm.DB }

var configStrs1 configStrs

func initTests() {
	err := godotenv.Load("../.env")

	if err != nil {
		panic("Failed to load env")
	}

	dbStr := os.Getenv("DB_URL")
	db, err := gorm.Open(postgres.Open(dbStr), &gorm.Config{})
	db.AutoMigrate(&models.Blog{})
	configStrs1.db = db

	if err != nil {
		panic("Failed to connect to DB")
	}

	configStrs1.db.AutoMigrate(&models.Blog{})
}

func SetupRouter() *gin.Engine {

	initTests()
	r := gin.Default()
	postgresBlogRepo := repository.NewPostgresBlogRepo(configStrs1.db)
	blogService := NewBlogService(postgresBlogRepo)
	r.GET("/blogs", blogService.GetAllBlogs)
	r.GET("/blogs/:slug", blogService.GetBlog)
	r.POST("/blogs", blogService.CreateBlog)

	return r
}

func TestGetBlogs(t *testing.T) {

	router := SetupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/blogs", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

}

func TestCreateBlog(t *testing.T) {

	testBlog := models.Blog{Title: "Tester", Slug: "Tester", Body: "Empty"}

	router := SetupRouter()
	w := httptest.NewRecorder()
	body, _ := json.Marshal(testBlog)
	req, _ := http.NewRequest("POST", "/blogs", bytes.NewReader(body))
	router.ServeHTTP(w, req)
	assert.Equal(t, 201, w.Code)

}
func TestGetBlog(t *testing.T) {

	//testBlog := models.Blog{Title: "Tester", Slug: "Tester", Body: "Empty"}

	router := SetupRouter()
	w := httptest.NewRecorder()
	//body, _ := json.Marshal(testBlog)
	req, _ := http.NewRequest("GET", "/blogs/Tester", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

}
func TestDeleteBlog(t *testing.T) {

	router := SetupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/blogs/Tester", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

}

// func TestGetBlog(t *testing.T) {
// 	router := SetupRouter()
// 	w := httptest.NewRecorder()
// 	req, _ := http.NewRequest("GET", "/blogs/user1", nil)
// 	router.ServeHTTP(w, req)
// 	assert.Equal(t, 200, w.Code)

// }
