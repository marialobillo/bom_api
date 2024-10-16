package handler

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/marialobillo/bom_api/internal/entities"
	"github.com/marialobillo/bom_api/internal/service"
)


type SiteHandler struct {
	service service.SiteService
}

func NewSiteHandler(service service.SiteService) *SiteHandler {
	return &SiteHandler{
		service: service,
	}
}

func (h *SiteHandler) CreateSite(c *fiber.Ctx) error {
	var site entities.Site

	if err := c.BodyParser(&site); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	ctx := c.Context()

	createdSite, err := h.service.CreateSite(ctx, &site)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"message": "Site created successfully",
		"data":    createdSite,
	})
}

func (h *SiteHandler) UpdateSite(c *fiber.Ctx) error {
	id := c.Params("id")
	site := new(entities.Site)

	if err := c.BodyParser(&site); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	ctx := c.Context()
	site.ID = id
	updatedSite, err := h.service.UpdateSite(ctx, id, site)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Site updated successfully",
		"data":    updatedSite,
	})
}