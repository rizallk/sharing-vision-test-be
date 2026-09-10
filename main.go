package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/rizallk/sharing-vision-test-be/config"
	"github.com/rizallk/sharing-vision-test-be/controllers"
	"github.com/rizallk/sharing-vision-test-be/repositories"
	"github.com/rizallk/sharing-vision-test-be/routes"
	"github.com/rizallk/sharing-vision-test-be/services"
)

func main() {
	// Load environment & connect database
	config.LoadEnv()
	config.ConnectDB()

	app := fiber.New()

	app.Use(cors.New())

	// Article
	articleRepo := repositories.NewArticleRepository(config.DB)
	articleService := services.NewArticleService(articleRepo)
	articleController := controllers.NewArticleController(articleService)

	// Setup Routes
	routes.Setup(app, articleController)

	// Run Server
	port := config.AppConfig.AppPort
	log.Println("Server is running on port :", port)
	log.Fatal(app.Listen(":" + port))
}
