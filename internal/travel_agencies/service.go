package travel_agencies

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"qolibaba/pkg/adapter/storage"
	"qolibaba/pkg/adapter/storage/types"
	"time"
)

type TravelAgencyService struct {
	repository *storage.TravelAgencyRepository
}

func NewTravelAgencyService(repository *storage.TravelAgencyRepository) *TravelAgencyService {
	return &TravelAgencyService{repository: repository}
}

type Vehicle struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

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
	hotelServiceURL := fmt.Sprintf("http://bank-service:8081/api/v1/hotels/get-all")
	req, err := http.NewRequest(http.MethodPost, hotelServiceURL, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request to hotel service: %v", err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending confirmation to hotel service: %v", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("bank service returned status: %v", resp.Status)
	}

	var hotels []types.Hotel
	err = json.NewDecoder(resp.Body).Decode(&hotels)
	if err != nil {
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

// OfferTour handles the logic for offering a new tour, including booking hotels and storing tour details.
func (s *TravelAgencyService) OfferTour(userID uint, roomID uint, startTime time.Time, endTime time.Time, totalPrice float64, goingTransferVehicleID uint, returnTransferVehicleID uint, hotelID uint) (*types.TourBooking, error) {
	hotelServiceURL := "http://localhost:8081/rooms/book-hotel"
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
		log.Printf("Error creating request to hotel service: %v", err)
		return nil, fmt.Errorf("error creating request to hotel service: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Error sending request to hotel service: %v", err)
		return nil, fmt.Errorf("error sending request to hotel service: %v", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		log.Printf("Failed to book hotel: %v", string(body))
		return nil, fmt.Errorf("failed to book hotel: %v", resp.Status)
	}

	var bookingResponse map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&bookingResponse); err != nil {
		log.Printf("Error decoding hotel response: %v", err)
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
		log.Printf("Error saving tour: %v", err)
		return nil, fmt.Errorf("error saving tour: %v", err)
	}

	return savedTour, nil
}
