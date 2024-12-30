package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"travel-booking-app/internal/payment-service/repository" // مسیر را متناسب با پروژه خود اصلاح کنید
	"travel-booking-app/internal/payment-service/service"
)

type PaymentHandler struct {
	paymentService service.PaymentService
}

func NewPaymentHandler() *PaymentHandler {
	repo := repository.NewPaymentRepository()
	svc := service.NewPaymentService(repo)
	return &PaymentHandler{paymentService: svc}
}

type CreatePaymentRequest struct {
	UserID string  `json:"user_id"`
	Amount float64 `json:"amount"`
}

type PaymentResponse struct {
	ID            string  `json:"id"`
	UserID        string  `json:"user_id"`
	Amount        float64 `json:"amount"`
	Status        string  `json:"status"`
	TransactionID string  `json:"transaction_id"`
	CreatedAt     string  `json:"created_at"`
	UpdatedAt     string  `json:"updated_at"`
}

func (h *PaymentHandler) CreatePayment(w http.ResponseWriter, r *http.Request) {
	var req CreatePaymentRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	paymentID, err := h.paymentService.CreatePayment(req.UserID, req.Amount)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response := map[string]string{
		"message":    "Payment created successfully",
		"payment_id": paymentID,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *PaymentHandler) GetPayment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	paymentID := vars["id"]

	payment, err := h.paymentService.GetPaymentByID(paymentID)
	if err != nil {
		http.Error(w, "Payment not found", http.StatusNotFound)
		return
	}

	response := PaymentResponse{
		ID:            payment.ID.String(),
		UserID:        payment.UserID.String(),
		Amount:        payment.Amount,
		Status:        payment.Status,
		TransactionID: payment.TransactionID,
		CreatedAt:     payment.CreatedAt.Format(time.RFC3339),
		UpdatedAt:     payment.UpdatedAt.Format(time.RFC3339),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *PaymentHandler) GetPaymentsByUser(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("user_id")
	if userID == "" {
		http.Error(w, "user_id is required", http.StatusBadRequest)
		return
	}

	payments, err := h.paymentService.GetPaymentsByUserID(userID)
	if err != nil {
		http.Error(w, "Error fetching payments", http.StatusInternalServerError)
		return
	}

	var response []PaymentResponse
	for _, payment := range payments {
		response = append(response, PaymentResponse{
			ID:            payment.ID.String(),
			UserID:        payment.UserID.String(),
			Amount:        payment.Amount,
			Status:        payment.Status,
			TransactionID: payment.TransactionID,
			CreatedAt:     payment.CreatedAt.Format(time.RFC3339),
			UpdatedAt:     payment.UpdatedAt.Format(time.RFC3339),
		})
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

type UpdatePaymentStatusRequest struct {
	Status string `json:"status"`
}

func (h *PaymentHandler) UpdatePaymentStatus(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	paymentID := vars["id"]

	var req UpdatePaymentStatusRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	err := h.paymentService.UpdatePaymentStatus(paymentID, req.Status)
	if err != nil {
		http.Error(w, "Error updating payment status", http.StatusInternalServerError)
		return
	}

	response := map[string]string{
		"message": "Payment status updated successfully",
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *PaymentHandler) DeletePayment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	paymentID := vars["id"]

	err := h.paymentService.DeletePayment(paymentID)
	if err != nil {
		http.Error(w, "Error deleting payment", http.StatusInternalServerError)
		return
	}

	response := map[string]string{
		"message": "Payment deleted successfully",
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
