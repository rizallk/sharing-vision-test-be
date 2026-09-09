package controllers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/rizallk/sharing-vision-test-be/models"
	"github.com/rizallk/sharing-vision-test-be/services"
)

type ArticleController struct {
	articleService services.ArticleService
}

func NewArticleController(articleService services.ArticleService) *ArticleController {
	return &ArticleController{articleService}
}

// POST /article/
func (c *ArticleController) CreateArticle(ctx *fiber.Ctx) error {
	var req models.ArticleRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	if err := c.articleService.CreateArticle(req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Article created successfully"})
}

// GET /article/:limit/:offset
func (c *ArticleController) GetAllArticles(ctx *fiber.Ctx) error {
	limit, err := strconv.Atoi(ctx.Params("limit"))
	if err != nil || limit < 0 {
		limit = 10 // Default fallback limit
	}

	offset, err := strconv.Atoi(ctx.Params("offset"))
	if err != nil || offset < 0 {
		offset = 0 // Default fallback offset
	}

	articles, err := c.articleService.GetAllArticles(limit, offset)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.JSON(articles)
}

// GET /article/:id
func (c *ArticleController) GetArticleByID(ctx *fiber.Ctx) error {
	id, err := strconv.ParseUint(ctx.Params("id"), 10, 32)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID format"})
	}

	article, err := c.articleService.GetArticleByID(uint(id))
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.JSON(article)
}

// POST /article/:id (atau PUT/PATCH sesuai dokumen)
func (c *ArticleController) UpdateArticle(ctx *fiber.Ctx) error {
	id, err := strconv.ParseUint(ctx.Params("id"), 10, 32)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID format"})
	}

	var req models.ArticleRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	if err := c.articleService.UpdateArticle(uint(id), req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.JSON(fiber.Map{"message": "Article updated successfully"})
}

// DELETE /article/:id
func (c *ArticleController) DeleteArticle(ctx *fiber.Ctx) error {
	id, err := strconv.ParseUint(ctx.Params("id"), 10, 32)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID format"})
	}

	if err := c.articleService.DeleteArticle(uint(id)); err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.JSON(fiber.Map{"message": "Article deleted successfully"})
}
