package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/marialobillo/bom_api/internal/handler"
)

func PartRoutes(app fiber.Router, partHandler *handler.PartHandler) {
	app.Post("/parts", partHandler.CreatePart)
}

func SupplierRoutes(app fiber.Router, supplierHandler *handler.SupplierHandler) {
	app.Post("/suppliers", supplierHandler.CreateSupplier)
	app.Put("/suppliers/:id", supplierHandler.UpdateSupplier)
	app.Delete("/suppliers/:id", supplierHandler.DeleteSupplier)
	app.Get("/suppliers/:id", supplierHandler.GetSupplierByID)
	app.Get("/suppliers", supplierHandler.GetAllSuppliers)
}

func SiteRoutes(app fiber.Router, siteHandler *handler.SiteHandler) {
	app.Post("/sites", siteHandler.CreateSite)
}

func Routes(app *fiber.App, handlers map[string]interface{}) {
	apiv1 := app.Group("/api/v1")

	if partHandler, ok := handlers["part"].(*handler.PartHandler); ok {
		PartRoutes(apiv1, partHandler)
	}

	if supplierHandler, ok := handlers["supplier"].(*handler.SupplierHandler); ok {
		SupplierRoutes(apiv1, supplierHandler)
	}

	if siteHandler, ok := handlers["site"].(*handler.SiteHandler); ok {
		SiteRoutes(apiv1, siteHandler)
	}
}
