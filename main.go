package main

import (
	"log"
	"os"

	"github.com/fahrigunadi/backend-sharing-vision/controllers"
	"github.com/fahrigunadi/backend-sharing-vision/initializers"
	"github.com/gofiber/fiber/v2"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
}

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Sharing Vision API")
	})

	app.Get("/article/:limit/:offset", controllers.GetAllArticles)
	app.Post("/article", controllers.CreateArticle)
	app.Put("/article/:id", controllers.UpdateArticle)
	app.Get("/article/:id", controllers.ShowArticle)
	app.Delete("/article/:id", controllers.DeleteArticle)

	log.Fatal(app.Listen(":" + os.Getenv("APP_PORT")))
}
