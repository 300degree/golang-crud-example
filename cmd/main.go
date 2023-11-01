package main

import (
	"github.com/300degree/golang-crud-example/pkg/config/db"
	"github.com/300degree/golang-crud-example/pkg/service"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	db.Init()

	app.Route("/product", func(router fiber.Router) {
		router.Post("/create", service.CreateProduct)
	})

	app.Listen(":8000")
}
