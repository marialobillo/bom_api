package repository

import (
	"database/sql"
	"log"

	"github.com/marialobillo/bom_api/internal/entities"
)

type SupplierRepository interface {
	CreateSupplier(supplier *entities.Supplier) error
}

type SupplierRepo struct {
	db *sql.DB
}

func NewSupplierRepository(db *sql.DB) *SupplierRepo {
	return &SupplierRepo{
		db: db,
	}
}

func (r *SupplierRepo) CreateSupplier(supplier *entities.Supplier) error {
	query := "INSERT INTO suppliers (name, address) VALUES ($1, $2) RETURNING id"
	err := r.db.QueryRow(query, supplier.Name, supplier.Address).Scan(&supplier.ID)
	if err != nil {
		log.Println("Error creating supplier: ", err)
		return err
	}
	return nil
}