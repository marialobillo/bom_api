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
	ctx := c.Context()

	createdSupplier, err := h.service.CreateSupplier(ctx, &supplier)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"message": "Supplier created successfully",
		"data":    createdSupplier,
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

func (h *SupplierHandler) DeleteSupplier(c *fiber.Ctx) error {
	id := c.Params("id")

	if err := h.service.DeleteSupplier(c.Context(), id); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Supplier deleted successfully",
	})
}

func (h *SupplierHandler) GetSupplierByID(c *fiber.Ctx) error {
	id := c.Params("id")

	supplier, err := h.service.GetSupplierByID(c.Context(), id)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if supplier == nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"error": "supplier not found",
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"data": supplier,
	})
}

func (h *SupplierHandler) GetAllSuppliers(c *fiber.Ctx) error {
	suppliers, err := h.service.GetAllSuppliers(c.Context())
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"data": suppliers,
	})
}
