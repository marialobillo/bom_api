package handler_test

import (
	"bytes"
	"context"
	"encoding/json"
	"io"

	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/marialobillo/bom_api/internal/entities"
	"github.com/marialobillo/bom_api/internal/handler"
	"github.com/marialobillo/bom_api/internal/routes"
	"github.com/stretchr/testify/assert"
)

type MockSupplierService struct {
	CreateSupplierFn  func(ctx context.Context, supplier *entities.Supplier) (*entities.Supplier, error)
	UpdateSupplierFn  func(ctx context.Context, id string, supplier *entities.Supplier) (*entities.Supplier, error)
	DeleteSupplierFn  func(ctx context.Context, id string) error
	GetSupplierByIDFn func(ctx context.Context, id string) (*entities.Supplier, error)
	GetAllSuppliersFn func(ctx context.Context) ([]entities.Supplier, error)
}

func (m *MockSupplierService) CreateSupplier(ctx context.Context, supplier *entities.Supplier) (*entities.Supplier, error) {
	if m.CreateSupplierFn != nil {
		return m.CreateSupplierFn(ctx, supplier)
	}
	return nil, nil
}

func (m *MockSupplierService) UpdateSupplier(ctx context.Context, id string, supplier *entities.Supplier) (*entities.Supplier, error) {
	if m.UpdateSupplierFn != nil {
		return m.UpdateSupplierFn(ctx, id, supplier) 
	}
	return nil, nil 
}

func (m *MockSupplierService) DeleteSupplier(ctx context.Context, id string) error {
	if m.DeleteSupplierFn != nil {
		return m.DeleteSupplierFn(ctx, id) 
	}
	return nil 
}

func (m *MockSupplierService) GetSupplierByID(ctx context.Context, id string) (*entities.Supplier, error) {
	if m.GetSupplierByIDFn != nil {
		return m.GetSupplierByIDFn(ctx, id) 
	}
	return nil, nil 
}

func (m *MockSupplierService) GetAllSuppliers(ctx context.Context) ([]entities.Supplier, error) {
	if m.GetAllSuppliersFn != nil {
		return m.GetAllSuppliersFn(ctx) 
	}
	return nil, nil 
}

func setup() *fiber.App {
	app := fiber.New()


	mockService := &MockSupplierService{
		CreateSupplierFn: func(ctx context.Context, supplier *entities.Supplier) (*entities.Supplier, error) {
			return &entities.Supplier{
				ID:      "1",
				Name:    supplier.Name,
				Contact: supplier.Contact,
				Email:   supplier.Email,
				Address: supplier.Address,
			}, nil
		},
		UpdateSupplierFn: func(ctx context.Context, id string, supplier *entities.Supplier) (*entities.Supplier, error) {
			return supplier, nil
		},
		DeleteSupplierFn: func(ctx context.Context, id string) error {
			return nil
		},
		GetSupplierByIDFn: func(ctx context.Context, id string) (*entities.Supplier, error) {
			if id == "1" {
				return &entities.Supplier{
					ID:      "1",
					Name:    "Supplier A",
					Contact: "Contact A",
					Email:   "supplier@mail.com",
					Address: "Address A",
				}, nil
			}
			return nil, nil 
		},
		GetAllSuppliersFn: func(ctx context.Context) ([]entities.Supplier, error) {
			return []entities.Supplier{
				{
					ID:      "1",
					Name:    "Supplier A",
					Contact: "Contact A",
					Email:   "supplier@mail.com",
					Address: "Address A",
				},
				{
					ID:      "2",
					Name:    "Supplier B",
					Contact: "Contact B",
					Email:   "supplierb@mail.com",
					Address: "Address B",
				},
			}, nil
		},
	}

	supplierHandler := handler.NewSupplierHandler(mockService)

	handlers := map[string]interface{}{
		"supplier": supplierHandler,
	}
	routes.Routes(app, handlers)
	return app
}

func TestSupplierHandler(t *testing.T) {
	app := fiber.New()

	mockService := &MockSupplierService{}
	supplierHandler := handler.NewSupplierHandler(mockService)

	app.Post("/suppliers", supplierHandler.CreateSupplier)
	app.Put("/suppliers/:id", supplierHandler.UpdateSupplier)
	app.Delete("/suppliers/:id", supplierHandler.DeleteSupplier)

	t.Run("Create Supplier", func(t *testing.T) {
		supplier := &entities.Supplier{
			Name:    "Supplier A",
			Contact: "Contact A",
			Email:   "supplierA@example.com",
			Address: "Address A",
		}
		body, _ := json.Marshal(supplier)

		req := httptest.NewRequest(http.MethodPost, "/suppliers", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")

		res, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusCreated, res.StatusCode)

		var response map[string]interface{}
		json.NewDecoder(res.Body).Decode(&response)
		assert.Equal(t, "Supplier created successfully", response["message"])
	})

	t.Run("Update Supplier", func(t *testing.T) {
		supplier := &entities.Supplier{
			ID:      "1", // Assuming the ID is "1" for this test
			Name:    "Supplier A Updated",
			Contact: "Contact A Updated",
			Email:   "supplierA_updated@example.com",
			Address: "Address A Updated",
		}
		body, _ := json.Marshal(supplier)

		req := httptest.NewRequest(http.MethodPut, "/suppliers/1", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")

		res, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, res.StatusCode)

		var response map[string]interface{}
		json.NewDecoder(res.Body).Decode(&response)
		assert.Equal(t, "Supplier updated successfully", response["message"])
	})

	t.Run("Delete Supplier", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodDelete, "/suppliers/1", nil) // Assuming the ID is "1"
		res, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, res.StatusCode)

		var response map[string]interface{}
		json.NewDecoder(res.Body).Decode(&response)
		assert.Equal(t, "Supplier deleted successfully", response["message"])
	})
}

func TestGetSupplierByID(t *testing.T) {
	app := setup()

	req := httptest.NewRequest(http.MethodGet, "/api/v1/suppliers/1", nil)
	res, err := app.Test(req)
	if err != nil {
		t.Fatalf("Error testing request: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, res.StatusCode)
	}

	var response map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		t.Fatalf("Error decoding response: %v", err)
	}

	expected := "Supplier A"
	if response["data"].(map[string]interface{})["name"] != expected {
		t.Errorf("Expected supplier name %s, got %s", expected, response["data"].(map[string]interface{})["name"])
	}
}

func TestGetAllSuppliers(t *testing.T) {
    app := setup()

    // Create a sample supplier using a pointer
    supplier := &entities.Supplier{
        Name:    "Supplier A",
        Contact: "Contact A",
        Email:   "supplierA@example.com",
        Address: "Address A",
    }

    supplierJSON, err := json.Marshal(supplier)
    if err != nil {
        t.Fatalf("Error marshalling supplier: %v", err)
    }

    createReq := httptest.NewRequest(http.MethodPost, "/api/v1/suppliers", bytes.NewBuffer(supplierJSON))
	createReq.Header.Set("Content-Type", "application/json")

    createRes, err := app.Test(createReq)
    if err != nil {
        t.Fatalf("Error testing request: %v", err)
    }
    defer createRes.Body.Close()

    body, err := io.ReadAll(createRes.Body)
    if err != nil {
        t.Fatalf("Error reading response body: %v", err)
    }

    if createRes.StatusCode != http.StatusCreated {
        t.Errorf("Expected status code %d, got %d. Response body: %s", http.StatusCreated, createRes.StatusCode, body)
    }

    req := httptest.NewRequest(http.MethodGet, "/api/v1/suppliers", nil)
    res, err := app.Test(req)
    if err != nil {
        t.Fatalf("Error testing request: %v", err)
    }
    defer res.Body.Close()

    if res.StatusCode != http.StatusOK {
        t.Errorf("Expected status code %d, got %d", http.StatusOK, res.StatusCode)
    }

    var response map[string]interface{}
    if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
        t.Fatalf("Error decoding response: %v", err)
    }

    if len(response["data"].([]interface{})) < 1 {
        t.Errorf("Expected at least 1 supplier, got %d", len(response["data"].([]interface{})))
    }
}
