package repository

import (
	"context"
	"database/sql"

	"github.com/marialobillo/bom_api/internal/entities"
)

type SiteRepository interface {
	CreateSite(ctx context.Context, site *entities.Site) (*entities.Site, error)
	GetSiteByID(ctx context.Context, id string) (*entities.Site, error)
	UpdateSite(ctx context.Context, site *entities.Site) (*entities.Site, error)
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

func (r *SiteRepo) GetSiteByID(ctx context.Context, id string) (*entities.Site, error) {
	query := "SELECT id, name, address, location FROM sites WHERE id = $1"
	site := &entities.Site{}
	row := r.db.QueryRowContext(ctx, query, id)

	err := row.Scan(&site.ID, &site.Name, &site.Address, &site.Location)
	if err != nil {
		return nil, &RepositoryError{
			Message: "failed to get site by id",
			Err:     err,
		}
	}
	return site, nil
}

func (r *SiteRepo) UpdateSite(ctx context.Context, site *entities.Site) (*entities.Site, error) {
	query := "UPDATE sites SET name = $1, address = $2, location = $3 where id = $4"
	_, err := r.db.ExecContext(ctx, query, site.Name, site.Address, site.Location, site.ID)
	if err != nil {
		return nil, &RepositoryError{
			Message: "failed to update site",
			Err:     err,
		}
	}
	return site, nil
}