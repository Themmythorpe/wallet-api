package controllers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
	"wallet-api-go/controllers"
	"wallet-api-go/db"
	"wallet-api-go/models"

	"github.com/gin-gonic/gin"
)

// MockDB provides a mock implementation of the database functions
type MockDB struct{}

// Mocking the DB functions
func (m *MockDB) Create(interface{}) *MockDB {
	return m
}

func (m *MockDB) Save(interface{}) *MockDB {
	return m
}

func (m *MockDB) Where(query interface{}, args ...interface{}) *MockDB {
	return m
}

func (m *MockDB) First(interface{}, ...interface{}) *MockDB {
	return m
}

func (m *MockDB) Begin() *MockDB {
	return m
}

func (m *MockDB) Rollback() *MockDB {
	return m
}

func (m *MockDB) Commit() *MockDB {
	return m
}

func TestCreateWallet(t *testing.T) {
	// Set up a mock database connection
	db.DB = &MockDB{}

	// Create a new Gin router
	router := gin.Default()
	router.POST("/create", controllers.CreateWallet)

	// Create a new HTTP request
	payload := []byte(`{"UserID": 1, "Currency": "USD"}`)
	req, err := http.NewRequest("POST", "/create", bytes.NewBuffer(payload))
	if err != nil {
		t.Fatal(err)
	}

	// Create a response recorder to record the response
	recorder := httptest.NewRecorder()

	// Serve the HTTP request to the router
	router.ServeHTTP(recorder, req)

	// Check the response status code
	if status := recorder.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}
}

func TestCreditWallet(t *testing.T) {
	// Set up a mock database connection
	db.DB = &MockDB{}

	// Create a new Gin router
	router := gin.Default()
	router.POST("/:wallet_id/credit", controllers.CreditWallet)

	// Create a new HTTP request
	payload := []byte(`{"amount": 100, "description": "Test credit"}`)
	req, err := http.NewRequest("POST", "/1/credit", bytes.NewBuffer(payload))
	if err != nil {
		t.Fatal(err)
	}

	// Create a response recorder to record the response
	recorder := httptest.NewRecorder()

	// Serve the HTTP request to the router
	router.ServeHTTP(recorder, req)

	// Check the response status code
	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestDebitWallet(t *testing.T) {
	// Set up a mock database connection
	db.DB = &MockDB{}

	// Create a new Gin router
	router := gin.Default()
	router.POST("/:wallet_id/debit", controllers.DebitWallet)

	// Create a new HTTP request
	payload := []byte(`{"amount": 50, "description": "Test debit"}`)
	req, err := http.NewRequest("POST", "/1/debit", bytes.NewBuffer(payload))
	if err != nil {
		t.Fatal(err)
	}

	// Create a response recorder to record the response
	recorder := httptest.NewRecorder()

	// Serve the HTTP request to the router
	router.ServeHTTP(recorder, req)

	// Check the response status code
	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestGetWalletTransactions(t *testing.T) {
	// Set up a mock database connection
	db.DB = &MockDB{}

	// Create a new Gin router
	router := gin.Default()
	router.GET("/:wallet_id/transactions", controllers.GetWalletTransactions)

	// Create a new HTTP request
	req, err := http.NewRequest("GET", "/1/transactions", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a response recorder to record the response
	recorder := httptest.NewRecorder()

	// Serve the HTTP request to the router
	router.ServeHTTP(recorder, req)

	// Check the response status code
	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body
	var response struct {
		Transactions []models.Transaction `json:"transactions"`
	}
	if err := json.Unmarshal(recorder.Body.Bytes(), &response); err != nil {
		t.Errorf("error decoding response body: %v", err)
	}

}