package controllers

import (
	"github.com/fahrigunadi/backend-sharing-vision/entities"
	"github.com/fahrigunadi/backend-sharing-vision/initializers"
	"github.com/fahrigunadi/backend-sharing-vision/libraries"
	"github.com/fahrigunadi/backend-sharing-vision/models"
	"github.com/gofiber/fiber/v2"
)

var validation = libraries.NewValidation()
var stauses = []string{"draft", "publish", "trash"}

func GetAllArticles(c *fiber.Ctx) error {
	articles := []models.Article{}

	limit, err := c.ParamsInt("limit")
	if err != nil || limit < 1 {
		limit = 10
	}

	offset, err := c.ParamsInt("offset")
	if err != nil || offset < 1 {
		offset = 0
	}

	filterStatus := c.Query("status")
	if !libraries.ContainsString(stauses, filterStatus) {
		filterStatus = "publish"
	}

	initializers.DB.Limit(limit).Offset(offset).Where("status = ?", filterStatus).Find(&articles)

	return c.JSON(articles)
}

func CreateArticle(c *fiber.Ctx) error {
	eArticle := new(entities.Article)

	if err := c.BodyParser(&eArticle); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"status":  "error",
			"message": "Review your input",
			"errors":  err,
		})
	}

	vErr := validation.Struct(eArticle)

	if vErr != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"status":  "error",
			"message": "Review your input",
			"errors":  vErr,
		})
	}

	initializers.DB.Create(&eArticle)

	return c.JSON(fiber.Map{})
}

func UpdateArticle(c *fiber.Ctx) error {
	article := models.Article{}

	initializers.DB.Find(&article, c.Params("id"))

	if article.Id == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "error",
			"message": "No article found with ID",
		})
	}

	eArticle := new(entities.Article)

	if err := c.BodyParser(&eArticle); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"status":  "error",
			"message": "Review your input",
			"errors":  err,
		})
	}

	vErr := validation.Struct(eArticle)

	if vErr != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"status":  "error",
			"message": "Review your input",
			"errors":  vErr,
		})
	}

	initializers.DB.Model(&article).Updates(&eArticle)

	return c.JSON(fiber.Map{})
}

func ShowArticle(c *fiber.Ctx) error {
	article := models.Article{}

	initializers.DB.Find(&article, c.Params("id"))

	if article.Id == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "error",
			"message": "No article found with ID",
		})
	}

	return c.JSON(article)
}

func DeleteArticle(c *fiber.Ctx) error {
	article := models.Article{}

	initializers.DB.Find(&article, c.Params("id"))

	if article.Id == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "error",
			"message": "No article found with ID",
		})
	}

	initializers.DB.Delete(&article)

	return c.JSON(fiber.Map{})
}
