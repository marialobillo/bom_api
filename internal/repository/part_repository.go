package repository

import (
	"database/sql"
	"log"
	"github.com/marialobillo/bom_api/internal/entities"
)

type PartRepository interface {
	CreatePart(part *entities.Part) error
}

type PartRepo struct {
	db *sql.DB
}

func NewPartRepository(database *sql.DB) PartRepository {
	return &PartRepo{
		db: database,
	}
}

func (r *PartRepo) CreatePart(part *entities.Part) error {
	query := "INSERT INTO parts (name, supplier_id, price, available, description, quantity, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id"
	err := r.db.QueryRow(query, part.Name, part.Supplier_id, part.Price, part.Available, part.Description, part.Quantity, part.Created_at, part.Updated_at).Scan(&part.ID)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
