package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rizallk/sharing-vision-test-be/controllers"
)

func Setup(app *fiber.App, articleController *controllers.ArticleController) {
	// Grouping URL endpoints
	articleRoute := app.Group("/article")

	articleRoute.Post("/", articleController.CreateArticle)
	articleRoute.Get("/:limit/:offset", articleController.GetAllArticles)
	articleRoute.Get("/:id", articleController.GetArticleByID)

	// Mendukung PUT/PATCH/POST untuk update sesuai requirement file PDF
	articleRoute.Put("/:id", articleController.UpdateArticle)
	articleRoute.Patch("/:id", articleController.UpdateArticle)
	articleRoute.Post("/:id", articleController.UpdateArticle)

	// Mendukung DELETE/POST untuk delete sesuai requirement file PDF
	articleRoute.Delete("/:id", articleController.DeleteArticle)
}
