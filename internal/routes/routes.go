package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/marialobillo/bom_api/internal/handler"
)

func SetupRoutes(app fiber.Router, partHandler *handler.PartHandler) {
	app.Post("/parts", partHandler.CreatePart)
}