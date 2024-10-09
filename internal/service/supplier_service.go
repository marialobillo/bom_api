package service

import (
	"github.com/marialobillo/bom_api/internal/repository"
	"github.com/marialobillo/bom_api/internal/entities"
)

type SupplierService interface {
	repo repository.SupplierService
}

func NewSupplierService(repo repository.SupplierService) *SupplierService {
	return &supplierService{
		repo: repo,
	}
}

func (s *SupplierService) CreateSupplier(supplier *entities.Supplier) error {
	if supplier.Name == "" {
		return errors.New("supplier name is required")
	}

	return s.repo.CreateSupplier(supplier)
}