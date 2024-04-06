package service

import (
	"go-blog/models"
	"go-blog/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BlogService struct {
	repo repository.BlogRepository
}

func NewBlogService(repo repository.BlogRepository) *BlogService {
	return &BlogService{repo: repo}
}

func (blogService *BlogService) GetAllBlogs(c *gin.Context) {
	var blogs *[]models.Blog
	blogs, err := blogService.repo.GetBlogs()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err})
		return
	}
	// Add handler next to complete the spaghetti

	c.JSON(http.StatusOK, gin.H{"data": blogs})

}

func (blogService *BlogService) CreateBlog(c *gin.Context) {

	request := models.BlogRequest{}
	//var blogs []models.Blog
	//result := inits.DB.Find(&blogs)

	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	blog := models.Blog{Title: request.Title, Slug: request.Slug, Body: request.Body}

	err := blogService.repo.CreateBlog(blog)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Post successful"})

}

// func BlogHandler(c *gin.Context) gin.HandlerFunc {
// 	return func(c *gin.Context) {

// 	}
// }

func (blogService *BlogService) GetBlog(c *gin.Context) {
	var blog *models.Blog
	slug := c.Param("slug")
	blog, err := blogService.repo.GetBlog(slug)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": blog})

}

// func (blogService *BlogService) UpdateBlog(c *gin.Context) {
// 	var blog *models.Blog
// 	c.BindJSON(&blog)
// 	slug := c.Param("slug")

// 	blog, err := blogService.repo.GetBlog(slug)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"data": blog})

// }

// type errorResponse struct {
// 	name string
// 	code int
// }

func (blogService *BlogService) DeleteBlog(c *gin.Context) {
	slug := c.Param("slug")
	err := blogService.repo.DeleteBlog(slug)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Delete Successful"})
}
