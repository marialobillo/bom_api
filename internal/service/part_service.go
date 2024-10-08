package service

import (
	"github.com/marialobillo/bom_api/internal/entities"
	"github.com/marialobillo/bom_api/internal/repository"
	"errors"
)

func CreatePart(part *entities.Part) error {
	if part.Name == "" || part.Supplier_id == "" {
		return errors.New("Name and Supplier_id are required")
	}
	return repository.CreatePart(part)
}