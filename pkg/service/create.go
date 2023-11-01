package service

import (
	"log"

	"github.com/300degree/golang-crud-example/pkg/config/db"
	"github.com/300degree/golang-crud-example/pkg/model"
	"github.com/gofiber/fiber/v2"
)

func CreateProduct(c *fiber.Ctx) error {
	var payload *model.Product
	if err := c.BodyParser(&payload); err != nil {
		log.Fatal(err)
	}

	product := model.Product{
		Name:  payload.Name,
		Price: payload.Price,
	}

	db.Init().Save(&product)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "success", "body": product})
}
