package repository

import (
	"log"

	"github.com/marialobillo/bom_api/internal/entities"
	"github.com/marialobillo/bom_api/infrastructure/db"
)

func CreatePart(part *entities.Part) error {
	query := "INSERT INTO parts (id, name, supplier_id, price, available, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id"
	err := db.DB.QueryRow(query, part.ID, part.Name, part.Supplier_id, part.Price, part.Available, part.Created_at, part.Updated_at).Scan(&part.ID)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
