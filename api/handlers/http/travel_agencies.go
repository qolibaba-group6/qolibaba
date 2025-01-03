package http

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/ehsansobhani/travel_agencies/internal/travel_agencies/domain"
	"github.com/ehsansobhani/travel_agencies/internal/travel_agencies/service"
	"github.com/ehsansobhani/travel_agencies/pkg/logger"
)

type TravelAgenciesHandler struct {
	CompanyService *service.CompanyService
	TripService    *service.TripService
	Logger         *logger.Logger
}

func NewTravelAgenciesHandler(companyService *service.CompanyService, tripService *service.TripService, logger *logger.Logger) *TravelAgenciesHandler {
	return &TravelAgenciesHandler{
		CompanyService: companyService,
		TripService:    tripService,
		Logger:         logger,
	}
}

// Company Handlers

func (h *TravelAgenciesHandler) CreateCompany(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Name  string `json:"name"`
		Owner string `json:"owner"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.Logger.Error("CreateCompany: invalid request body", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	company, err := h.CompanyService.CreateCompany(req.Name, req.Owner)
	if err != nil {
		h.Logger.Error("CreateCompany: service error", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(company)
}

func (h *TravelAgenciesHandler) GetCompany(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr, ok := vars["id"]
	if !ok {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}

	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	company, err := h.CompanyService.GetCompany(id)
	if err != nil {
		if err == domain.ErrCompanyNotFound {
			http.Error(w, "Company not found", http.StatusNotFound)
			return
		}
		h.Logger.Error("GetCompany: service error", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(company)
}

func (h *TravelAgenciesHandler) UpdateCompany(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr, ok := vars["id"]
	if !ok {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}

	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	var req struct {
		Name  string `json:"name"`
		Owner string `json:"owner"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.Logger.Error("UpdateCompany: invalid request body", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	company, err := h.CompanyService.UpdateCompany(id, req.Name, req.Owner)
	if err != nil {
		if err == domain.ErrCompanyNotFound {
			http.Error(w, "Company not found", http.StatusNotFound)
			return
		}
		h.Logger.Error("UpdateCompany: service error", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(company)
}

func (h *TravelAgenciesHandler) DeleteCompany(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr, ok := vars["id"]
	if !ok {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}

	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	err = h.CompanyService.DeleteCompany(id)
	if err != nil {
		if err == domain.ErrCompanyNotFound {
			http.Error(w, "Company not found", http.StatusNotFound)
			return
		}
		h.Logger.Error("DeleteCompany: service error", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *TravelAgenciesHandler) ListCompanies(w http.ResponseWriter, r *http.Request) {
	companies, err := h.CompanyService.ListCompanies()
	if err != nil {
		h.Logger.Error("ListCompanies: service error", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(companies)
}

// Trip Handlers

func (h *TravelAgenciesHandler) CreateTrip(w http.ResponseWriter, r *http.Request) {
	var req struct {
		CompanyID     string  `json:"company_id"`
		Type          string  `json:"type"`
		Origin        string  `json:"origin"`
		Destination   string  `json:"destination"`
		DepartureTime string  `json:"departure_time"`
		ReleaseDate   string  `json:"release_date"`
		Tariff        float64 `json:"tariff"`
		Status        string  `json:"status"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.Logger.Error("CreateTrip: invalid request body", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	trip, err := h.TripService.CreateTrip(req.CompanyID, req.Type, req.Origin, req.Destination, req.DepartureTime, req.ReleaseDate, req.Tariff, req.Status)
	if err != nil {
		h.Logger.Error("CreateTrip: service error", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(trip)
}

func (h *TravelAgenciesHandler) GetTrip(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr, ok := vars["id"]
	if !ok {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}

	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	trip, err := h.TripService.GetTrip(id)
	if err != nil {
		if err == domain.ErrTripNotFound {
			http.Error(w, "Trip not found", http.StatusNotFound)
			return
		}
		h.Logger.Error("GetTrip: service error", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(trip)
}

func (h *TravelAgenciesHandler) UpdateTrip(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr, ok := vars["id"]
	if !ok {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}

	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	var req struct {
		CompanyID     string  `json:"company_id"`
		Type          string  `json:"type"`
		Origin        string  `json:"origin"`
		Destination   string  `json:"destination"`
		DepartureTime string  `json:"departure_time"`
		ReleaseDate   string  `json:"release_date"`
		Tariff        float64 `json:"tariff"`
		Status        string  `json:"status"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.Logger.Error("UpdateTrip: invalid request body", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	trip, err := h.TripService.UpdateTrip(id, req.CompanyID, req.Type, req.Origin, req.Destination, req.DepartureTime, req.ReleaseDate, req.Tariff, req.Status)
	if err != nil {
		if err == domain.ErrTripNotFound {
			http.Error(w, "Trip not found", http.StatusNotFound)
			return
		}
		h.Logger.Error("UpdateTrip: service error", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(trip)
}

func (h *TravelAgenciesHandler) DeleteTrip(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr, ok := vars["id"]
	if !ok {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}

	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	err = h.TripService.DeleteTrip(id)
	if err != nil {
		if err == domain.ErrTripNotFound {
			http.Error(w, "Trip not found", http.StatusNotFound)
			return
		}
		h.Logger.Error("DeleteTrip: service error", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *TravelAgenciesHandler) ListTrips(w http.ResponseWriter, r *http.Request) {
	companyID := r.URL.Query().Get("company_id")
	trips, err := h.TripService.ListTrips(companyID)
	if err != nil {
		h.Logger.Error("ListTrips: service error", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(trips)
}
