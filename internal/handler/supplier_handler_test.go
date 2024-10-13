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
	"github.com/marialobillo/bom_api/internal/routes"
	"github.com/stretchr/testify/assert"
)

type MockSupplierService struct {
	CreateSupplierFn  *entities.Supplier
	UpdateSupplierFn  *entities.Supplier
	DeleteSupplierFn  string
	GetSupplierByIDFn *entities.Supplier
	GetAllSuppliersFn []entities.Supplier
}

func (m *MockSupplierService) CreateSupplier(ctx context.Context, supplier *entities.Supplier) error {
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

func (m *MockSupplierService) GetSupplierByID(ctx context.Context, id string) (*entities.Supplier, error) {
	if m.GetSupplierByIDFn != nil && m.GetSupplierByIDFn.ID == id {
		return m.GetSupplierByIDFn, nil
	}
	return nil, nil
}

func (m *MockSupplierService) GetAllSuppliers(ctx context.Context) ([]entities.Supplier, error) {
	return m.GetAllSuppliersFn, nil
}

func setup() *fiber.App {
	app := fiber.New()

	// Mock database connection and repository
	mockService := &MockSupplierService{
		GetSupplierByIDFn: &entities.Supplier{
			ID:      "1",
			Name:    "Supplier A",
			Contact: "Contact A",
			Email:   "supplier@mail.com",
			Address: "Address A",
		},
	}
	supplierHandler := handler.NewSupplierHandler(mockService)

	// Initialize routes
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

	// t.Run("Get Supplier By ID", func(t *testing.T) {
	// 	mockService.GetSupplierByIDFn = &entities.Supplier{
	// 		ID:      "1",
	// 		Name:    "Supplier A",
	// 		Contact: "Contact A",
	// 		Email:   "supplier@mail.com",
	// 		Address: "Address A",
	// 	}

	// 	req := httptest.NewRequest(http.MethodGet, "/suppliers/1", nil)

	// 	res, err := app.Test(req)

	// 	assert.NoError(t, err)

	// 	assert.Equal(t, http.StatusOK, res.StatusCode)

	// 	var response map[string]interface{}
	// 	err = json.NewDecoder(res.Body).Decode(&response)

	// 	assert.NoError(t, err)

	// 	assert.Equal(t, "Supplier A", response["data"].(map[string]interface{})["name"])
	// 	assert.Equal(t, "Contact A", response["data"].(map[string]interface{})["contact"])
	// 	assert.Equal(t, "supplier@mail.com", response["data"].(map[string]interface{})["email"])
	// 	assert.Equal(t, "Address A", response["data"].(map[string]interface{})["address"])
	// })
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
