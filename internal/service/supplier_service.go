package service

import (
	"context"
	"errors"

	"github.com/marialobillo/bom_api/internal/entities"
	"github.com/marialobillo/bom_api/internal/repository"
)

type SupplierService interface {
	CreateSupplier(supplier *entities.Supplier) error
	UpdateSupplier(ctx context.Context, supplier *entities.Supplier) error
}

type supplierService struct {
	repo repository.SupplierRepository
}

func NewSupplierService(repo repository.SupplierRepository) SupplierService {
	return &supplierService{
		repo: repo,
	}
}

func (s *supplierService) CreateSupplier(supplier *entities.Supplier) error {
	if supplier.Name == "" {
		return errors.New("supplier name is required")
	}

	return s.repo.CreateSupplier(supplier)
}

func (s *supplierService) UpdateSupplier(ctx context.Context, supplier *entities.Supplier) error {
    existingSupplier, err := s.repo.GetSupplierByID(ctx, supplier.ID)
    if err != nil {
        return err
    }
    if existingSupplier == nil {
        return errors.New("supplier not found")
    }
    return s.repo.UpdateSupplier(ctx, supplier)
}
