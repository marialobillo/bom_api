package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/marialobillo/bom_api/internal/entities"
	"github.com/stretchr/testify/assert"
)

type MockSiteService struct {
	CreateSiteFn func(ctx context.Context, site *entities.Site) (*entities.Site, error)
	UpdateSiteFn func(ctx context.Context, id string, site *entities.Site) (*entities.Site, error)
	DeleteSiteFn func(ctx context.Context, id string) error
}

func (m *MockSiteService) CreateSite(ctx context.Context, site *entities.Site) (*entities.Site, error) {
	return m.CreateSiteFn(ctx, site)
}

func (m *MockSiteService) UpdateSite(ctx context.Context, id string, site *entities.Site) (*entities.Site, error) {
	return m.UpdateSiteFn(ctx, id, site)
}

func (m *MockSiteService) DeleteSite(ctx context.Context, id string) error {
	return m.DeleteSiteFn(ctx, id)
}

func TestCreateSite(t *testing.T) {
	app := fiber.New()
	mockService := &MockSiteService{
		CreateSiteFn: func(ctx context.Context, site *entities.Site) (*entities.Site, error) {
			return &entities.Site{ID: "1", Name: site.Name, Location: site.Location, Address: site.Address}, nil
		},
	}
	handler := NewSiteHandler(mockService)
	app.Post("/sites", handler.CreateSite)

	t.Run("Successful site creation", func(t *testing.T) {
		site := entities.Site{Name: "Test Site", Location: "Test Location", Address: "Test Address"}
		body, _ := json.Marshal(site)
		req := httptest.NewRequest(http.MethodPost, "/sites", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")

		resp, _ := app.Test(req)

		assert.Equal(t, http.StatusCreated, resp.StatusCode)
	})

	t.Run("Failed site creation with invalid input", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/sites", nil)
		resp, _ := app.Test(req)

		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	})
}

func TestUpdateSite(t *testing.T) {
	app := fiber.New()
	mockService := &MockSiteService{
		UpdateSiteFn: func(ctx context.Context, id string, site *entities.Site) (*entities.Site, error) {
			return &entities.Site{ID: id, Name: site.Name, Location: site.Location, Address: site.Address}, nil
		},
	}
	handler := NewSiteHandler(mockService)
	app.Put("/sites/:id", handler.UpdateSite)

	t.Run("Successful site update", func(t *testing.T) {
		site := entities.Site{Name: "Updated Site", Location: "Updated Location", Address: "Updated Address"}
		body, _ := json.Marshal(site)
		req := httptest.NewRequest(http.MethodPut, "/sites/1", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")

		resp, _ := app.Test(req)

		assert.Equal(t, http.StatusOK, resp.StatusCode)
	})

	t.Run("Failed site update with invalid input", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPut, "/sites/1", nil)
		resp, _ := app.Test(req)

		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	})
}

func TestDeleteSite(t *testing.T) {
    app := fiber.New()
    mockService := &MockSiteService{
        DeleteSiteFn: func(ctx context.Context, id string) error {
            if id == "1" {
                return nil
            }
            return errors.New("site not found")
        },
    }
    handler := NewSiteHandler(mockService)
    app.Delete("/sites/:id", handler.DeleteSite)

    t.Run("Successful site deletion", func(t *testing.T) {
        req := httptest.NewRequest(http.MethodDelete, "/sites/1", nil)
        resp, _ := app.Test(req)

        assert.Equal(t, http.StatusOK, resp.StatusCode)

        var body map[string]string
        json.NewDecoder(resp.Body).Decode(&body)
        assert.Equal(t, "Site deleted successfully", body["message"])
    })

    t.Run("Failed site deletion with site not found", func(t *testing.T) {
        req := httptest.NewRequest(http.MethodDelete, "/sites/999", nil)
        resp, _ := app.Test(req)

        assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)

        var body map[string]string
        json.NewDecoder(resp.Body).Decode(&body)
        assert.Equal(t, "site not found", body["error"])
    })
}

