package web

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// Mock implementation of UserRouterInterface
type mockUserRouter struct{}

func (m *mockUserRouter) GetUsers(w http.ResponseWriter, r *http.Request) {
	// Mock implementation for GetUsers
}

func (m *mockUserRouter) AddUser(w http.ResponseWriter, r *http.Request) {
	// Mock implementation for AddUser
}

func (m *mockUserRouter) UpdateUser(w http.ResponseWriter, r *http.Request) {
	// Mock implementation for UpdateUser
}

func (m *mockUserRouter) DeleteUser(w http.ResponseWriter, r *http.Request) {
	// Mock implementation for DeleteUser
}

func TestNewAppWithCORS(t *testing.T) {
	mockRouter := &mockUserRouter{}
	app := NewApp(true, mockRouter)

	// Verify that the handlers are set up correctly
	expectedPaths := []string{
		"/api/get-users",
		"/api/add-user",
		"/api/update-user",
		"/api/delete-user",
		"/",
	}

	for _, path := range expectedPaths {
		_, exists := app.handlers[path]
		if !exists {
			t.Errorf("Handler for path %s not found", path)
		}

		// You can further test the handler here if needed
		// For example, you can use httptest to send mock HTTP requests and check responses
	}
}

func TestNewAppWithoutCORS(t *testing.T) {
	mockRouter := &mockUserRouter{}
	app := NewApp(false, mockRouter)

	// Verify that the handlers are set up correctly with CORS disabled
	expectedPaths := []string{
		"/api/get-users",
		"/api/add-user",
		"/api/update-user",
		"/api/delete-user",
		"/",
	}

	for _, path := range expectedPaths {
		_, exists := app.handlers[path]
		if !exists {
			t.Errorf("Handler for path %s not found", path)
		}

		// You can further test the handler here if needed
		// For example, you can use httptest to send mock HTTP requests and check responses
	}
}

func TestServe(t *testing.T) {
	mockRouter := &mockUserRouter{}
	app := NewApp(true, mockRouter)

	// Create a test HTTP server
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// You can simulate requests to your server here
	}))

	defer testServer.Close()

	// Start the server in a goroutine
	go func() {
		err := app.Serve()
		if err != nil {
			t.Errorf("Serve() returned an error: %v", err)
		}
	}()

	// You can use the testServer to send HTTP requests to the running server and test its responses
	// For example:
	// resp, err := http.Get(testServer.URL + "/api/get-users")
	// ...

	// You can add more test cases to check various aspects of server behavior
}

// Add test cases for the disableCors function as needed.
