package service

import (
	"context"
	"errors"

	"github.com/marialobillo/bom_api/internal/entities"
	"github.com/marialobillo/bom_api/internal/repository"
)

type SupplierService interface {
	CreateSupplier(ctx context.Context, supplier *entities.Supplier) (*entities.Supplier, error) 
	UpdateSupplier(ctx context.Context, id string, supplier *entities.Supplier) (*entities.Supplier, error)
	DeleteSupplier(ctx context.Context, id string) error
	GetSupplierByID(ctx context.Context, id string) (*entities.Supplier, error)
	GetAllSuppliers(ctx context.Context) ([]entities.Supplier, error)
}

type supplierService struct {
	repo repository.SupplierRepository
}

type ServiceError struct {
	Message string
	Err     error
}

func (e *ServiceError) Error() string {
	return e.Message + ": " + e.Err.Error()
}

func NewSupplierService(repo repository.SupplierRepository) SupplierService {
	return &supplierService{
		repo: repo,
	}
}

func (s *supplierService) CreateSupplier(ctx context.Context, supplier *entities.Supplier) (*entities.Supplier, error) {
	if supplier.Name == "" {
		return nil, errors.New("supplier name is required")
	}

	return s.repo.CreateSupplier(ctx, supplier)
}

func (s *supplierService) UpdateSupplier(ctx context.Context, id string, supplier *entities.Supplier) (*entities.Supplier, error) {
	existingSupplier, err := s.repo.GetSupplierByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if existingSupplier == nil {
		return nil, errors.New("supplier not found")
	}

	supplier.ID = existingSupplier.ID 
	updatedSupplier, err := s.repo.UpdateSupplier(ctx, supplier) 
	if err != nil {
		return nil, err
	}
	return updatedSupplier, nil
}

func (s *supplierService) DeleteSupplier(ctx context.Context, id string) error {
	existingSupplier, err := s.repo.GetSupplierByID(ctx, id)
	if err != nil {
		return err
	}
	if existingSupplier == nil {
		return errors.New("supplier not found")
	}
	return s.repo.DeleteSupplier(ctx, id)
}

func (s *supplierService) GetSupplierByID(ctx context.Context, id string) (*entities.Supplier, error) {
	supplier, err := s.repo.GetSupplierByID(ctx, id)
	if err != nil {
		return nil, &ServiceError{
			Message: "failed to get supplier by id",
			Err:     err,
		}
	}
	return supplier, nil
}

func (s *supplierService) GetAllSuppliers(ctx context.Context) ([]entities.Supplier, error) {
	suppliers, err := s.repo.GetAllSuppliers(ctx)
	if err != nil {
		return nil, &ServiceError{
			Message: "failed to get all suppliers",
			Err:     err,
		}
	}
	return suppliers, nil
}
