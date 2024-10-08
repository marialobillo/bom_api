package handler

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/marialobillo/bom_api/internal/entities"
	"github.com/marialobillo/bom_api/internal/service"
)

type PartHandler struct {
	partService *service.PartService
}

func NewPartHandler(partService *service.PartService) *PartHandler {
	return &PartHandler{partService: partService}
}

func (h *PartHandler) CreatePart(c * fiber.Ctx) error {
	var part entities.Part 

	if err := c.BodyParser(&part); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if err := h.partService.CreatePart(&part); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(http.StatusCreated).JSON(part)
}
