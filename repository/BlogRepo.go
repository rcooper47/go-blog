package repository

import (
	"go-blog/models"
)

type BlogRepository interface {
	GetBlogs() (*[]models.Blog, error)
	CreateBlog(blog models.Blog) error
	GetBlog(slug string) (*models.Blog, error)
	DeleteBlog(slug string) error // maybe rename to less specific names
}
