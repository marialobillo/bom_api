package repository

import (
	"context"
	"database/sql"

	"github.com/marialobillo/bom_api/internal/entities"
)

type SiteRepository interface {
	CreateSite(ctx context.Context, site *entities.Site) (*entities.Site, error)
}

type SiteRepo struct {
	db *sql.DB
}


func NewSiteRepository(db *sql.DB) *SiteRepo {
	return &SiteRepo{
		db: db,
	}
}

func (r *SiteRepo) CreateSite(ctx context.Context, site *entities.Site) (*entities.Site, error) {
	query := "INSERT INTO sites (name, address, location) VALUES ($1, $2, $3) RETURNING id"
	err := r.db.QueryRowContext(ctx, query, site.Name, site.Address, site.Location).Scan(&site.ID)
	if err != nil {
		return nil, &RepositoryError{
			Message: "failed to create site",
			Err:     err,
		}
	}
	return site, nil
}