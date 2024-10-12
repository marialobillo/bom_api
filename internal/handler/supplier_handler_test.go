package handler_test

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/marialobillo/bom_api/internal/entities"
	"github.com/marialobillo/bom_api/internal/handler"
	"github.com/stretchr/testify/assert"
)

type MockSupplierService struct {
	CreateSupplierFn *entities.Supplier
	UpdateSupplierFn *entities.Supplier
	DeleteSupplierFn string
}

func (m *MockSupplierService) CreateSupplier(supplier *entities.Supplier) error {
	m.CreateSupplierFn = supplier
	return nil
}

func (m *MockSupplierService) UpdateSupplier(ctx context.Context, supplier *entities.Supplier) error {
	m.UpdateSupplierFn = supplier
	return nil
}

func (m *MockSupplierService) DeleteSupplier(ctx context.Context, id string) error {
	m.DeleteSupplierFn = id
	return nil
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