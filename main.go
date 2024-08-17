package main

import (
	"log"
	"os"

	"github.com/fahrigunadi/backend-sharing-vision/controllers"
	"github.com/fahrigunadi/backend-sharing-vision/initializers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
}

func main() {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowHeaders: "Origin,Content-Type,Accept,Content-Length,Accept-Language,Accept-Encoding,Connection,Access-Control-Allow-Origin",
		AllowOrigins: os.Getenv("CORS_ALLOW_ORIGINS"),
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
	}))

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
