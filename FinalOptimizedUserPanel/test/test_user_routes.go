package test

import (
	"net/http"
	"net/http/httptest"
	"src/controllers"
	"testing"
)

func TestGetUsers(t *testing.T) {
	req := httptest.NewRequest("GET", "/users", nil)
	w := httptest.NewRecorder()
	controllers.GetUsers(w, req)
	if w.Result().StatusCode != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Result().StatusCode)
	}
	expectedBody := `["User1","User2","User3"]`
	actualBody := w.Body.String()
	if actualBody != expectedBody {
		t.Errorf("Expected body %s, got %s", expectedBody, actualBody)
	}
}
