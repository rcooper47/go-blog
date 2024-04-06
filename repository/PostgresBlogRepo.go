package repository

import (
	"go-blog/models"

	"gorm.io/gorm"
)

type PostgresBlogRepo struct {
	db *gorm.DB
}

func NewPostgresBlogRepo(db *gorm.DB) *PostgresBlogRepo {
	return &PostgresBlogRepo{
		db: db,
	}
}

// GetBlogs implements BlogRepository.
func (pr *PostgresBlogRepo) GetBlogs() (*[]models.Blog, error) {
	var blogs []models.Blog
	if err := pr.db.Find(&blogs).Error; err != nil {
		return nil, err
	}

	return &blogs, nil
}

// CreateBlog implements BlogRepository.
func (pr *PostgresBlogRepo) CreateBlog(blog models.Blog) error {
	response := pr.db.Create(&blog)
	if response.Error != nil {
		return response.Error
	}
	return nil
}

func (pr *PostgresBlogRepo) GetBlog(slug string) (*models.Blog, error) {

	var blog models.Blog
	if err := pr.db.Where("slug = ?", slug).First(&blog).Error; err != nil {
		return nil, err
	}
	return &blog, nil

}
func (pr *PostgresBlogRepo) DeleteBlog(slug string) error {
	var blog models.Blog
	if err := pr.db.Where("slug = ?", slug).Delete(&blog).Error; err != nil {
		return err
	}
	return nil
}
