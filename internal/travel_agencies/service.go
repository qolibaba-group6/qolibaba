package travel_agencies

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"qolibaba/pkg/adapter/storage"
	"qolibaba/pkg/adapter/storage/types"
	"qolibaba/pkg/messaging"
	"strconv"
	"time"
)

type TravelAgencyService struct {
	repository *storage.TravelAgencyRepository
	messaging  *messaging.Messaging
}

func NewTravelAgencyService(repository *storage.TravelAgencyRepository, messaging *messaging.Messaging) *TravelAgencyService {
	return &TravelAgencyService{
		repository: repository,
		messaging:  messaging,
	}
}

type Vehicle struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

// RegisterNewAgency registers a new travel agency, ensuring unique emails
func (s *TravelAgencyService) RegisterNewAgency(agency *types.TravelAgency) (*types.TravelAgency, error) {
	existingAgency, _ := s.repository.FindByEmail(agency.Email)
	if existingAgency != nil {
		return nil, errors.New("email already in use")
	}

	if err := s.repository.Create(agency); err != nil {
		return nil, err
	}

	return agency, nil
}

// GetAllHotelsAndVehicles fetches all hotels and vehicles for offering a new tour
func (s *TravelAgencyService) GetAllHotelsAndVehicles() (map[string]interface{}, error) {
	hotelServiceURL := fmt.Sprintf("%s/api/v1/hotels/get-all", os.Getenv("HOTEL_SERVICE_URL"))

	req, err := http.NewRequest(http.MethodGet, hotelServiceURL, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request to hotel service: %v", err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error connecting to hotel service: %v", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("hotel service returned status %v: %v", resp.StatusCode, string(body))
	}

	var hotels []types.Hotel
	if err := json.NewDecoder(resp.Body).Decode(&hotels); err != nil {
		return nil, fmt.Errorf("failed to decode hotel response: %v", err)
	}

	vehicles := []Vehicle{
		{ID: 1, Name: "Vehicle A", Type: "Bus"},
		{ID: 2, Name: "Vehicle B", Type: "Train"},
	}

	return map[string]interface{}{
		"hotels":   hotels,
		"vehicles": vehicles,
	}, nil
}

// OfferTour handles the logic for offering a new tour
func (s *TravelAgencyService) OfferTour(userID uint, roomID uint, startTime, endTime time.Time, totalPrice float64, goingTransferVehicleID, returnTransferVehicleID, hotelID uint) (*types.TourBooking, error) {
	if startTime.After(endTime) {
		return nil, errors.New("start time cannot be after end time")
	}

	hotelServiceURL := fmt.Sprintf("%s/api/v1/rooms/book-hotel", os.Getenv("HOTEL_SERVICE_URL"))
	bookingData := map[string]interface{}{
		"room_id":     roomID,
		"user_id":     userID,
		"start_time":  startTime,
		"end_time":    endTime,
		"total_price": totalPrice,
		"status":      "pending",
	}

	bookingJSON, _ := json.Marshal(bookingData)

	req, err := http.NewRequest(http.MethodPost, hotelServiceURL, bytes.NewBuffer(bookingJSON))
	if err != nil {
		return nil, fmt.Errorf("error creating request to hotel service: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error connecting to hotel service: %v", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("failed to book hotel: %v", string(body))
	}

	var bookingResponse map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&bookingResponse); err != nil {
		return nil, fmt.Errorf("error decoding hotel response: %v", err)
	}

	tour := &types.TourBooking{
		UserID:                  userID,
		RoomID:                  roomID,
		GoingTransferVehicleID:  goingTransferVehicleID,
		ReturnTransferVehicleID: returnTransferVehicleID,
		HotelID:                 hotelID,
		BookingDate:             time.Now(),
		TotalPrice:              totalPrice,
		BookingStatus:           "pending",
		Confirmed:               false,
	}

	savedTour, err := s.repository.SaveTour(tour)
	if err != nil {
		return nil, fmt.Errorf("error saving tour: %v", err)
	}

	return savedTour, nil
}

func (s *TravelAgencyService) CreateTourBooking(booking *types.TourBooking) (*types.TourBooking, error) {
	// Validate input
	if booking.StartTime.After(booking.EndTime) {
		return nil, fmt.Errorf("start time must be before end time")
	}

	// Calculate total price
	totalDays := booking.EndTime.Sub(booking.StartTime).Hours() / 24
	booking.TotalPrice = booking.PerDayPrice * totalDays

	// Create a new tour booking
	newBooking, err := s.repository.CreateTourBooking(booking)
	if err != nil {
		return nil, fmt.Errorf("error creating tour booking: %v", err)
	}

	// Generate claim
	claim := types.Claim{
		BuyerUserID:  booking.UserID,
		SellerUserID: booking.TourID,
		Amount:       booking.TotalPrice,
		ClaimType:    "tour",
		ClaimDetails: fmt.Sprintf("Tour booking from %s to %s", booking.StartTime.Format("2006-01-02"), booking.EndTime.Format("2006-01-02")),
		Status:       "pending",
	}

	// Serialize claim to JSON
	claimData, err := json.Marshal(claim)
	if err != nil {
		return nil, fmt.Errorf("error marshalling claim: %v", err)
	}

	claimID, err := s.messaging.PublishClaimToBank(claimData)
	if err != nil {
		return nil, fmt.Errorf("error sending claim to bank: %v", err)
	}

	// Convert claim ID to uint
	claimIDUint, err := strconv.ParseUint(claimID, 10, 32)
	if err != nil {
		return nil, fmt.Errorf("error converting claimID to uint: %v", err)
	}

	claimIDPointer := uint(claimIDUint)
	newBooking.ClaimID = &claimIDPointer

	// Update booking with claim ID
	updatedBooking, err := s.repository.UpdateTourBooking(newBooking)
	if err != nil {
		return nil, fmt.Errorf("error updating booking with claim ID: %v", err)
	}

	return updatedBooking, nil
}
