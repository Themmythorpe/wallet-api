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

func TestRegisterUser(t *testing.T) {
	// Set up a mock database connection
	db.DB = &MockDB{}

	// Create a new Gin router
	router := gin.Default()
	router.POST("/users/register", controllers.RegisterUser)

	// Create a new HTTP request
	payload := []byte(`{"email": "test@example.com", "password": "password"}`)
	req, err := http.NewRequest("POST", "/users/register", bytes.NewBuffer(payload))
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

func TestLoginUser(t *testing.T) {
	// Set up a mock database connection
	db.DB = &MockDB{}

	// Create a new Gin router
	router := gin.Default()
	router.POST("/users/login", controllers.LoginUser)

	// Create a new HTTP request
	payload := []byte(`{"email": "test@example.com", "password": "password"}`)
	req, err := http.NewRequest("POST", "/users/login", bytes.NewBuffer(payload))
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
		Token string `json:"token"`
	}
	if err := json.Unmarshal(recorder.Body.Bytes(), &response); err != nil {
		t.Errorf("error decoding response body: %v", err)
	}

}
