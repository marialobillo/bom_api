package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"github.com/marialobillo/bom_api/infrastructure/db"
	"github.com/marialobillo/bom_api/internal/entities"
	"github.com/marialobillo/bom_api/internal/service"
)

func main() {
	app := fiber.New()
	apiv1 := app.Group("/api/v1")

	db.Connect()

	apiv1.Post("/parts", func(c *fiber.Ctx) error {
		var part entities.Part

		if err := c.BodyParser(&part); err != nil {
			return c.Status(400).SendString("Invalid Request")
		}

		if err := service.CreatePart(&part); err != nil {
			return c.Status(500).SendString("Error creating part")
		}

		return c.Status(201).JSON(part)
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	if err := app.Listen(":3300"); err != nil {
		log.Fatal(err)
	}
}