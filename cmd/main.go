package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/marialobillo/bom_api/infrastructure/db"
	"github.com/marialobillo/bom_api/internal/handler"
	"github.com/marialobillo/bom_api/internal/repository"
	"github.com/marialobillo/bom_api/internal/routes"
	"github.com/marialobillo/bom_api/internal/service"
)

func main() {
	// Initialize database connection
	database := db.NewPostgresConnection()
	defer database.Close()

	sqlDB := database.DB

	// Setup  repository and service, and handler
	partRepo := repository.NewPartRepository(sqlDB)
	partService := service.NewPartService(partRepo)
	partHandler := handler.NewPartHandler(partService)

	// Setup repository and service, and handler for supplier
	supplierRepo := repository.NewSupplierRepository(sqlDB)
	supplierService := service.NewSupplierService(supplierRepo)
	supplierHandler := handler.NewSupplierHandler(supplierService)

	// Initialize handlers map
	handlers := map[string]interface{}{
		"part":     partHandler,
		"supplier": supplierHandler,
	}

	app := fiber.New()
	//apiv1 := app.Group("/api/v1")

	routes.Routes(app, handlers)

	if err := app.Listen(":3300"); err != nil {
		log.Fatal(err)
	}
}