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
	database, err := db.NewPostgresConnection()
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	defer database.Close()

	sqlDB := database.DB

	partRepo := repository.NewPartRepository(sqlDB)
	partService := service.NewPartService(partRepo)
	partHandler := handler.NewPartHandler(partService)

	supplierRepo := repository.NewSupplierRepository(sqlDB)
	supplierService := service.NewSupplierService(supplierRepo)
	supplierHandler := handler.NewSupplierHandler(supplierService)

	siteRepo := repository.NewSiteRepository(sqlDB)
	siteService := service.NewSiteService(siteRepo)
	siteHandler := handler.NewSiteHandler(siteService)

	handlers := map[string]interface{}{
		"part":     partHandler,
		"supplier": supplierHandler,
		"site":     siteHandler,
	}

	app := fiber.New()

	routes.Routes(app, handlers)

	if err := app.Listen(":3300"); err != nil {
		log.Fatal(err)
	}
}
