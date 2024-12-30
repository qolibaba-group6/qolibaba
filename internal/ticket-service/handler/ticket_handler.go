// internal/ticket-service/handler/ticket_handler.go
package handler

import (
	"encoding/json"
	"net/http"
	"travel-booking-app/internal/ticket-service/repository"
	"travel-booking-app/internal/ticket-service/service"

	"time"

	"github.com/gorilla/mux"
)

type TicketHandler struct {
	ticketService service.TicketService
}

func NewTicketHandler() *TicketHandler {
	repo := repository.NewTicketRepository()
	service := service.NewTicketService(repo)
	return &TicketHandler{
		ticketService: service,
	}
}

type CreateTicketRequest struct {
	UserID       string  `json:"user_id"`
	Type         string  `json:"type"`
	Price        float64 `json:"price"`
	ReturnPolicy string  `json:"return_policy"`
}

type UpdateTicketRequest struct {
	Type         string  `json:"type,omitempty"`
	Price        float64 `json:"price,omitempty"`
	Status       string  `json:"status,omitempty"`
	ReturnPolicy string  `json:"return_policy,omitempty"`
}

type TicketResponse struct {
	ID           string  `json:"id"`
	UserID       string  `json:"user_id"`
	Type         string  `json:"type"`
	Price        float64 `json:"price"`
	Status       string  `json:"status"`
	ReturnPolicy string  `json:"return_policy"`
	CreatedAt    string  `json:"created_at"`
	UpdatedAt    string  `json:"updated_at"`
}

func (h *TicketHandler) CreateTicket(w http.ResponseWriter, r *http.Request) {
	var req CreateTicketRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	ticketID, err := h.ticketService.CreateTicket(req.UserID, req.Type, req.Price, req.ReturnPolicy)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response := map[string]string{
		"message":   "Ticket created successfully",
		"ticket_id": ticketID,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *TicketHandler) GetTicket(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ticketID := vars["id"]

	ticket, err := h.ticketService.GetTicketByID(ticketID)
	if err != nil {
		http.Error(w, "Ticket not found", http.StatusNotFound)
		return
	}

	response := TicketResponse{
		ID:           ticket.ID.String(),
		UserID:       ticket.UserID.String(),
		Type:         ticket.Type,
		Price:        ticket.Price,
		Status:       ticket.Status,
		ReturnPolicy: ticket.ReturnPolicy,
		CreatedAt:    ticket.CreatedAt.Format(time.RFC3339),
		UpdatedAt:    ticket.UpdatedAt.Format(time.RFC3339),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *TicketHandler) GetTicketsByUser(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("user_id")
	if userID == "" {
		http.Error(w, "user_id is required", http.StatusBadRequest)
		return
	}

	tickets, err := h.ticketService.GetTicketsByUserID(userID)
	if err != nil {
		http.Error(w, "Error fetching tickets", http.StatusInternalServerError)
		return
	}

	var response []TicketResponse
	for _, ticket := range tickets {
		response = append(response, TicketResponse{
			ID:           ticket.ID.String(),
			UserID:       ticket.UserID.String(),
			Type:         ticket.Type,
			Price:        ticket.Price,
			Status:       ticket.Status,
			ReturnPolicy: ticket.ReturnPolicy,
			CreatedAt:    ticket.CreatedAt.Format(time.RFC3339),
			UpdatedAt:    ticket.UpdatedAt.Format(time.RFC3339),
		})
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *TicketHandler) UpdateTicket(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ticketID := vars["id"]

	var req UpdateTicketRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	err := h.ticketService.UpdateTicket(ticketID, req.Type, req.Price, req.Status, req.ReturnPolicy)
	if err != nil {
		http.Error(w, "Error updating ticket", http.StatusInternalServerError)
		return
	}

	response := map[string]string{
		"message": "Ticket updated successfully",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *TicketHandler) DeleteTicket(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ticketID := vars["id"]

	err := h.ticketService.DeleteTicket(ticketID)
	if err != nil {
		http.Error(w, "Error deleting ticket", http.StatusInternalServerError)
		return
	}

	response := map[string]string{
		"message": "Ticket deleted successfully",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
