package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/marialobillo/bom_api/internal/entities"
)

type SupplierRepository interface {
	CreateSupplier(ctx context.Context, supplier *entities.Supplier) error
	GetSupplierByID(ctx context.Context, id string) (*entities.Supplier, error)
	UpdateSupplier(ctx context.Context, supplier *entities.Supplier) error
	DeleteSupplier(ctx context.Context, id string) error
	GetAllSuppliers(ctx context.Context) ([]entities.Supplier, error)
}

type SupplierRepo struct {
	db *sql.DB
}

type RepositoryError struct {
	Message string
	Err     error
}

func (e *RepositoryError) Error() string {
	return e.Message + ": " + e.Err.Error()
}

func NewSupplierRepository(db *sql.DB) *SupplierRepo {
	return &SupplierRepo{
		db: db,
	}
}

func (r *SupplierRepo) CreateSupplier(ctx context.Context, supplier *entities.Supplier) error {
	query := "INSERT INTO suppliers (name, contact, email, address) VALUES ($1, $2, $3, $4) RETURNING id"
	err := r.db.QueryRowContext(ctx, query, supplier.Name, supplier.Contact, supplier.Email, supplier.Address).Scan(&supplier.ID)
	if err != nil {
		return &RepositoryError{
			Message: "failed to create supplier",
			Err:     err,
		}
	}
	return nil
}

func (r *SupplierRepo) GetSupplierByID(ctx context.Context, id string) (*entities.Supplier, error) {
	query := "SELECT id, name, contact, email, address FROM suppliers WHERE id = $1"
	supplier := &entities.Supplier{}
	row := r.db.QueryRowContext(ctx, query, id)

	err := row.Scan(&supplier.ID, &supplier.Name, &supplier.Contact, &supplier.Email, &supplier.Address)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("supplier with id %s not found: %w", id, err)
		}
		return nil, fmt.Errorf("failed to get supplier by id %s: %w", id, err)
	}
	return supplier, nil
}

func (r *SupplierRepo) UpdateSupplier(ctx context.Context, supplier *entities.Supplier) error {
	query := "UPDATE suppliers SET name = $1, contact = $2, email = $3, address = $4 WHERE id = $5"
	_, err := r.db.ExecContext(ctx, query, supplier.Name, supplier.Contact, supplier.Email, supplier.Address, supplier.ID)
	if err != nil {
		return &RepositoryError{
			Message: "failed to update supplier",
			Err:     err,
		}
	}
	return nil
}

func (r *SupplierRepo) DeleteSupplier(ctx context.Context, id string) error {
	query := "DELETE FROM suppliers WHERE id = $1"
	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("error executing delete query: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error getting rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no supplier found with id %s", id)
	}

	return nil
}

func (r *SupplierRepo) GetAllSuppliers(ctx context.Context) ([]entities.Supplier, error) {
	query := "SELECT id, name, contact, email, address FROM suppliers"
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, &RepositoryError{
			Message: "failed to get all suppliers",
			Err:     err,
		}
	}
	defer rows.Close()

	var suppliers []entities.Supplier
	for rows.Next() {
		supplier := entities.Supplier{}
		err = rows.Scan(&supplier.ID, &supplier.Name, &supplier.Contact, &supplier.Email, &supplier.Address)
		if err != nil {
			return nil, &RepositoryError{
				Message: "failed to scan supplier",
				Err:     err,
			}
		}
		suppliers = append(suppliers, supplier)
	}
	return suppliers, nil
}
