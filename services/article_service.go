package services

import (
	"errors"

	"github.com/rizallk/sharing-vision-test-be/models"
	"github.com/rizallk/sharing-vision-test-be/repositories"
)

type ArticleService interface {
	CreateArticle(req models.ArticleRequest) error
	GetAllArticles(limit int, offset int, status string) ([]models.Post, error)
	GetArticleByID(id uint) (*models.Post, error)
	UpdateArticle(id uint, req models.ArticleRequest) error
	DeleteArticle(id uint) error
}

type articleService struct {
	articleRepo repositories.ArticleRepository
}

func NewArticleService(articleRepo repositories.ArticleRepository) ArticleService {
	return &articleService{articleRepo}
}

// Fungsi helper untuk validasi sesuai spesifikasi soal
func (s *articleService) validateRequest(req models.ArticleRequest) error {
	if len(req.Title) < 20 {
		return errors.New("title is required and must be at least 20 characters")
	}
	if len(req.Content) < 200 {
		return errors.New("content is required and must be at least 200 characters")
	}
	if len(req.Category) < 3 {
		return errors.New("category is required and must be at least 3 characters")
	}
	if req.Status != "publish" && req.Status != "draft" && req.Status != "thrash" {
		return errors.New("status must be either publish, draft, or thrash")
	}
	return nil
}

func (s *articleService) CreateArticle(req models.ArticleRequest) error {
	if err := s.validateRequest(req); err != nil {
		return err
	}

	post := models.Post{
		Title:    req.Title,
		Content:  req.Content,
		Category: req.Category,
		Status:   req.Status,
	}

	return s.articleRepo.Create(&post)
}

func (s *articleService) GetAllArticles(limit int, offset int, status string) ([]models.Post, error) {
	return s.articleRepo.FindAll(limit, offset, status)
}

func (s *articleService) GetArticleByID(id uint) (*models.Post, error) {
	return s.articleRepo.FindByID(id)
}

func (s *articleService) UpdateArticle(id uint, req models.ArticleRequest) error {
	if err := s.validateRequest(req); err != nil {
		return err
	}

	post, err := s.articleRepo.FindByID(id)
	if err != nil {
		return errors.New("article not found")
	}

	post.Title = req.Title
	post.Content = req.Content
	post.Category = req.Category
	post.Status = req.Status

	return s.articleRepo.Update(post)
}

func (s *articleService) DeleteArticle(id uint) error {
	_, err := s.articleRepo.FindByID(id)
	if err != nil {
		return errors.New("article not found")
	}
	return s.articleRepo.Delete(id)
}
