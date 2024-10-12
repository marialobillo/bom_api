package handler

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/marialobillo/bom_api/internal/entities"
	"github.com/marialobillo/bom_api/internal/service"
)

type SupplierHandler struct {
	service service.SupplierService
}

func NewSupplierHandler(service service.SupplierService) *SupplierHandler {
	return &SupplierHandler{
		service: service,
	}
}

func (h *SupplierHandler) CreateSupplier(c *fiber.Ctx) error {
	var supplier entities.Supplier

	if err := c.BodyParser(&supplier); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if err := h.service.CreateSupplier(&supplier); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"message": "Supplier created successfully",
		"data":    supplier,
	})
}

func (h *SupplierHandler) UpdateSupplier(c *fiber.Ctx) error {
	id := c.Params("id")
	supplier := new(entities.Supplier)

	if err := c.BodyParser(supplier); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	supplier.ID = id

	if err := h.service.UpdateSupplier(c.Context(), supplier); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Supplier updated successfully",
		"data":    supplier,
	})
}