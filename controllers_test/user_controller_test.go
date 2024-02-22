package controllers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
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

func (m *MockDB) Where(query interface{}, args ...interface{}) *MockDB {
	return m
}

func (m *MockDB) First(interface{}, ...interface{}) *MockDB {
	return m
}

// SetupRouter initializes and returns a Gin router with the provided routes
func SetupRouter() *gin.Engine {
	router := gin.Default()
	controllers.SetupRouter(router)
	return router
}

// PerformRequest sends an HTTP request to the provided router and returns the response recorder
func PerformRequest(router *gin.Engine, method, path string, body []byte) *httptest.ResponseRecorder {
	req, err := http.NewRequest(method, path, bytes.NewBuffer(body))
	if err != nil {
		panic(err)
	}
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)
	return recorder
}

func TestRegisterUser(t *testing.T) {
	// Set up a mock database connection
	db.DB = &MockDB{}

	// Create a new Gin router
	router := SetupRouter()

	// Create a new HTTP request
	payload := []byte(`{"email": "test@example.com", "password": "password"}`)
	recorder := PerformRequest(router, "POST", "/users/register", payload)

	// Check the response status code
	if status := recorder.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}
}

func TestLoginUser(t *testing.T) {
	// Set up a mock database connection
	db.DB = &MockDB{}

	// Create a new Gin router
	router := SetupRouter()

	// Create a new HTTP request
	payload := []byte(`{"email": "test@example.com", "password": "password"}`)
	recorder := PerformRequest(router, "POST", "/users/login", payload)

	// Check the response status code
	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check the response body
	var response struct {
		Token string `json:"token"`
	}
	if err := json.Unmarshal(recorder.Body.Bytes(), &response); err != nil {
		t.Errorf("error decoding response body: %v", err)
	}
}
