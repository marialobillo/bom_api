package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/marialobillo/bom_api/infrastructure/db"
	"github.com/marialobillo/bom_api/internal/entities"
	"github.com/marialobillo/bom_api/internal/handler"
	"github.com/marialobillo/bom_api/internal/service"
)

func main() {
	
	// Initialize database connection
	dbConn := db.NewPostgresConnection()
	defer dbConn.Close()
	
	// Setup  repository and service
	partRepo := repository.NewPartRepository()
	partService := service.NewPartService(partRepo)

	// Setup handler 
	partHandler := handler.NewPartHandler(partService)

	app := fiber.New()
	apiv1 := app.Group("/api/v1")

	
	routes.SetupRoutes(apiv1, partHandler)

	if err := app.Listen(":3300"); err != nil {
		log.Fatal(err)
	}
}