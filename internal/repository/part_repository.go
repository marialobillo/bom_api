package repository

import (
	"fmt"
	"log"

	"github.com/marialobillo/bom_api/infrastructure/db"
	"github.com/marialobillo/bom_api/internal/entities"
)

func CreatePart(part *entities.Part) error {
	query := "INSERT INTO parts (id, name, supplier_id, price, available, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id"
	err := db.DB.QueryRow(query, part.ID, part.Name, part.Supplier_id, part.Price, part.Available, part.Created_at, part.Updated_at).Scan(&part.ID)
	if err != nil {
		log.Println("Error creating part: ", err)
		return fmt.Errorf("failed to create part: %w", err)
	}
	return nil
}
