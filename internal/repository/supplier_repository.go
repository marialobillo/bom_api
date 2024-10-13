package repository

import (
	"context"
	"database/sql"
	"log"

	"github.com/marialobillo/bom_api/internal/entities"
)

type SupplierRepository interface {
	CreateSupplier(supplier *entities.Supplier) error
	GetSupplierByID(ctx context.Context, id string) (*entities.Supplier, error)
	UpdateSupplier(ctx context.Context, supplier *entities.Supplier) error
	DeleteSupplier(ctx context.Context, id string) error
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
	query := "INSERT INTO suppliers (name, contact, email, address) VALUES ($1, $2, $3, $4) RETURNING id"
	err := r.db.QueryRow(query, supplier.Name, supplier.Contact, supplier.Email, supplier.Address).Scan(&supplier.ID)
	if err != nil {
		log.Println("Error creating supplier: ", err)
		return err
	}
	return nil
}

func (r *SupplierRepo) GetSupplierByID(ctx context.Context, id string) (*entities.Supplier, error) {
	query := "SELECT id, name, contact, email, address FROM suppliers WHERE id = $1"
	supplier := &entities.Supplier{}
	row := r.db.QueryRowContext(ctx, query, id)

	err := row.Scan(&supplier.ID, &supplier.Name, &supplier.Contact, &supplier.Email, &supplier.Address)
	if err != nil {
		log.Println("Error getting supplier by id: ", err)
		return nil, err
	}
	return supplier, nil
}

func (r *SupplierRepo) UpdateSupplier(ctx context.Context, supplier *entities.Supplier) error {
	query := "UPDATE suppliers SET name = $1, contact = $2, email = $3, address = $4 WHERE id = $5"
	_, err := r.db.ExecContext(ctx, query, supplier.Name, supplier.Contact, supplier.Email, supplier.Address, supplier.ID)
	if err != nil {
		log.Println("Error updating supplier: ", err)
		return err
	}
	return nil
}

func (r *SupplierRepo) DeleteSupplier(ctx context.Context, id string) error {
	query := "DELETE FROM suppliers WHERE id = $1"
	_, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		log.Println("Error deleting supplier: ", err)
		return err
	}
	return nil
}

func (r *SupplierRepo) GetAllSuppliers(ctx context.Context) ([]entities.Supplier, error) {
	query := "SELECT id, name, contact, email, address FROM suppliers"
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		log.Println("Error getting all suppliers: ", err)
		return nil, err
	}
	defer rows.Close()

	var suppliers []entities.Supplier
	for rows.Next() {
		supplier := entities.Supplier{}
		err = rows.Scan(&supplier.ID, &supplier.Name, &supplier.Contact, &supplier.Email, &supplier.Address)
		if err != nil {
			log.Println("Error scanning suppliers: ", err)
			return nil, err
		}
		suppliers = append(suppliers, supplier)
	}
	return suppliers, nil
}