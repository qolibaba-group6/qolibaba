package main

import (
	"bytes"
	"database/sql"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"your_project/src/database"
	"golang.org/x/crypto/bcrypt"
)

func TestRegister(t *testing.T) {
	reqBody := `{"username":"testuser","password":"password"}`
	req := httptest.NewRequest(http.MethodPost, "/register", strings.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	Register(w, req)

	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusCreated {
		t.Errorf("Expected status 201, got %d", res.StatusCode)
	}

	var username string
	err := database.DB.QueryRow("SELECT username FROM users WHERE username = $1", "testuser").Scan(&username)
	if err != nil {
		t.Error("User was not registered in the database")
	}
	if username != "testuser" {
		t.Errorf("Expected username 'testuser', got '%s'", username)
	}
}

func TestLogin(t *testing.T) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
	database.DB.Exec("INSERT INTO users (username, password_hash, role) VALUES ($1, $2, $3)", "testuser", string(hashedPassword), "User")

	reqBody := `{"username":"testuser","password":"password"}`
	req := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	Login(w, req)

	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200, got %d", res.StatusCode)
	}

	cookies := res.Cookies()
	if len(cookies) == 0 {
		t.Error("Expected a token cookie, but got none")
	}

	token := cookies[0].Value
	if token == "" {
		t.Error("Token cookie is empty")
	}
}

func TestChangeUserRole(t *testing.T) {
	hashedPasswordAdmin, _ := bcrypt.GenerateFromPassword([]byte("adminpassword"), bcrypt.DefaultCost)
	database.DB.Exec("INSERT INTO users (username, password_hash, role) VALUES ($1, $2, $3)", "admin", string(hashedPasswordAdmin), "Admin")

	hashedPasswordUser, _ := bcrypt.GenerateFromPassword([]byte("userpassword"), bcrypt.DefaultCost)
	database.DB.Exec("INSERT INTO users (username, password_hash, role) VALUES ($1, $2, $3)", "user", string(hashedPasswordUser), "User")

	reqBody := `{"username":"admin","password":"adminpassword"}`
	loginReq := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(reqBody))
	loginReq.Header.Set("Content-Type", "application/json")

	loginW := httptest.NewRecorder()
	Login(loginW, loginReq)

	loginRes := loginW.Result()
	defer loginRes.Body.Close()

	adminToken := loginRes.Cookies()[0].Value

	changeRoleReqBody := `{"username":"user","new_role":"Operator"}`
	req := httptest.NewRequest(http.MethodPost, "/change-role", strings.NewReader(changeRoleReqBody))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+adminToken)

	w := httptest.NewRecorder()
	ChangeUserRole(w, req)

	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200, got %d", res.StatusCode)
	}

	var role string
	err := database.DB.QueryRow("SELECT role FROM users WHERE username = $1", "user").Scan(&role)
	if err != nil {
		t.Error("Failed to retrieve user role from the database")
	}
	if role != "Operator" {
		t.Errorf("Expected role 'Operator', got '%s'", role)
	}
}
