// api/handlers/http/user.go
package http

import (
	"encoding/json"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// UserHandler defines the HTTP handlers for user operations
type UserHandler struct {
	Service port.UserService
}

// NewUserHandler creates a new UserHandler
func NewUserHandler(service port.UserService) *UserHandler {
	return &UserHandler{
		Service: service,
	}
}

// Register handles user registration
func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	var user domain.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := h.Service.Register(&user); err != nil {
		http.Error(w, "Error registering user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
		return
	}
}

// Login handles user login and JWT token generation
func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var creds struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	token, err := h.Service.Login(creds.Email, creds.Password)
	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(map[string]string{"token": token}); err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
		return
	}
}

// GetProfile handles fetching user profile (protected)
func (h *UserHandler) GetProfile(w http.ResponseWriter, r *http.Request) {
	userIDStr, ok := r.Context().Value("userID").(string)
	if !ok {
		http.Error(w, "User ID not found in context", http.StatusBadRequest)
		return
	}

	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	user, err := h.Service.GetProfile(userID)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
		return
	}
}


type UpdateRoleRequest struct {
	Role string `json:"role"`
}

func UpdateRole(svc *service.UserService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req UpdateRoleRequest 
		if err := c.BodyParser(&req); err != nil {
			return fiber.ErrBadRequest
		}

		userId, err := uuid.Parse(c.Params("id"))
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "missing or invalid user id")
		}

		err = svc.UpdateRole(c.UserContext(), userId, req.Role)
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}

		return nil
	}
}

func TestHandler(ctx *fiber.Ctx) error {
	logger := context.GetLogger(ctx.UserContext())


	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(updatedUser); err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
		return
	}
}

// DeleteUser handles deleting a user (protected)
func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	userIDStr, ok := r.Context().Value("userID").(string)
	if !ok {
		http.Error(w, "User ID not found in context", http.StatusBadRequest)
		return
	}

	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	if err := h.Service.DeleteUser(userID); err != nil {
		http.Error(w, "Error deleting user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
