package service

import (
	"context"
	"errors"

	"github.com/marialobillo/bom_api/internal/entities"
	"github.com/marialobillo/bom_api/internal/repository"
)

type SiteService interface {
	CreateSite(ctx context.Context, site *entities.Site) (*entities.Site, error)
	UpdateSite(ctx context.Context, id string, site *entities.Site) (*entities.Site, error)
	
}

type siteService struct {
	repo repository.SiteRepository
}

func NewSiteService(repo repository.SiteRepository) SiteService {
	return &siteService{
		repo: repo,
	}
}

func (s *siteService) CreateSite(ctx context.Context, site *entities.Site) (*entities.Site, error) {
	if site.Name == "" {
		return nil, errors.New("site name is required")
	}

	return s.repo.CreateSite(ctx, site)
}

func (s *siteService) UpdateSite(ctx context.Context, id string, site *entities.Site) (*entities.Site, error) {
	existingSite, err := s.repo.GetSiteByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if existingSite == nil {
		return nil, errors.New("site not found")
	}

	site.ID = existingSite.ID

	updatedSite, err := s.repo.UpdateSite(ctx, site)
	if err != nil {
		return nil, err
	}

	return updatedSite, nil
}