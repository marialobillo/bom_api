package service

import (
	"github.com/marialobillo/bom_api/internal/entities"
	"github.com/marialobillo/bom_api/internal/repository"
	"errors"
)

type PartService struct {
	partRepo repository.PartRepository
}

func NewPartService(partRepo repository.PartRepository) *PartService {
	return &PartService{
		partRepo: partRepo,
	}
}

func (s *PartService) CreatePart(part *entities.Part) error {
	if part.Name == "" {
		return errors.New("name is required")
	}
	if part.Supplier_id == "" {
		return errors.New("supplier_id is required")
	}
	if part.Price == 0 {
		return errors.New("price is required")
	}
	return s.partRepo.CreatePart(part)
}