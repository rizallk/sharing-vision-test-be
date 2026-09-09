package repositories

import (
	"github.com/rizallk/sharing-vision-test-be/models"
	"gorm.io/gorm"
)

type ArticleRepository interface {
	Create(post *models.Post) error
	FindAll(limit int, offset int) ([]models.Post, error)
	FindByID(id uint) (*models.Post, error)
	Update(post *models.Post) error
	Delete(id uint) error
}

type articleRepository struct {
	db *gorm.DB
}

func NewArticleRepository(db *gorm.DB) ArticleRepository {
	return &articleRepository{db}
}

func (r *articleRepository) Create(post *models.Post) error {
	return r.db.Create(post).Error
}

func (r *articleRepository) FindAll(limit int, offset int) ([]models.Post, error) {
	var posts []models.Post
	err := r.db.Limit(limit).Offset(offset).Find(&posts).Error
	return posts, err
}

func (r *articleRepository) FindByID(id uint) (*models.Post, error) {
	var post models.Post
	err := r.db.First(&post, id).Error
	return &post, err
}

func (r *articleRepository) Update(post *models.Post) error {
	return r.db.Save(post).Error
}

func (r *articleRepository) Delete(id uint) error {
	return r.db.Delete(&models.Post{}, id).Error
}
