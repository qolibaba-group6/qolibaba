package travel_agencies

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/redis/go-redis/v9"
	"io"
	"log"
	"net/http"
	"os"
	"qolibaba/internal/travel_agencies/port"
	"qolibaba/pkg/adapter/storage/types"
	"qolibaba/pkg/messaging"
	"time"
)

type TravelAgencyService struct {
	repository port.Repo
	messaging  *messaging.Messaging
	redis      *redis.Client
	validator  *validator.Validate
}

func NewTravelAgencyService(repository port.Repo, messaging *messaging.Messaging, redis *redis.Client) *TravelAgencyService {
	return &TravelAgencyService{
		repository: repository,
		messaging:  messaging,
		redis:      redis,
		validator:  validator.New(),
	}
}

type ServiceEvent struct {
	EventType    string  `json:"event_type"`
	ServiceID    uint    `json:"service_id"`
	Name         string  `json:"name"`
	GeneralPrice float64 `json:"general_price"`
	TourPrice    float64 `json:"tour_price"`
	Timestamp    string  `json:"timestamp"`
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
	cacheKey := "hotels_and_vehicles_data"
	cachedData, err := s.redis.Get(context.Background(), cacheKey).Result()
	if err == nil && cachedData != "" {
		var cachedResponse map[string]interface{}
		err := json.Unmarshal([]byte(cachedData), &cachedResponse)
		if err != nil {
			return nil, fmt.Errorf("failed to decode cached data: %v", err)
		}
		return cachedResponse, nil
	}
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
			fmt.Println("Error closing response body:", err)
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

	responseData := map[string]interface{}{
		"hotels":   hotels,
		"vehicles": vehicles,
	}

	cacheData, err := json.Marshal(responseData)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal data for caching: %v", err)
	}

	startTime := 7 * 24 * time.Hour
	err = s.redis.Set(context.Background(), cacheKey, cacheData, startTime).Err()
	if err != nil {
		return nil, fmt.Errorf("failed to cache data in Redis: %v", err)
	}

	return responseData, nil
}

// OfferTour handles the logic for offering a new tour
func (s *TravelAgencyService) OfferTour(tour *types.Tour) (*types.Tour, error) {

	if err := s.validator.Struct(tour); err != nil {
		log.Printf("Validation failed: %v", err)
		return nil, fmt.Errorf("validation error: %v", err)
	}

	hotelKey := fmt.Sprintf("hotel:%d", tour.HotelID)
	vehicleGoingKey := fmt.Sprintf("vehicle:%d", tour.GoingVehicleID)
	vehicleReturnKey := fmt.Sprintf("vehicle:%d", tour.ReturnVehicleID)

	hotel, err := s.redis.Get(context.Background(), hotelKey).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil, fmt.Errorf("hotel not found in Redis cache")
		}
		return nil, fmt.Errorf("error checking hotel in Redis: %v", err)
	}
	log.Printf("Hotel found in Redis: %s", hotel)

	vehicleGoing, err := s.redis.Get(context.Background(), vehicleGoingKey).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil, fmt.Errorf("going vehicle not found in Redis cache")
		}
		return nil, fmt.Errorf("error checking going vehicle in Redis: %v", err)
	}
	log.Printf("Going vehicle found in Redis: %s", vehicleGoing)

	vehicleReturn, err := s.redis.Get(context.Background(), vehicleReturnKey).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil, fmt.Errorf("return vehicle not found in Redis cache")
		}
		return nil, fmt.Errorf("error checking return vehicle in Redis: %v", err)
	}
	log.Printf("Return vehicle found in Redis: %s", vehicleReturn)

	tour.StartDate = time.Now().Add(1 * time.Hour)
	tour.EndDate = time.Now().Add(1 * time.Hour).Add(2 * time.Hour)
	if _, err := s.repository.SaveTour(tour); err != nil {
		log.Printf("Error saving tour: %v", err)
		return nil, err
	}

	return tour, nil
}

func (s *TravelAgencyService) CreateTourBooking(booking *types.TourBooking) (*types.TourBooking, error) {
	tour, err := s.repository.GetTourByID(booking.TourID)
	if err != nil {
		return nil, fmt.Errorf("error fetching tour: %v", err)
	}

	claim := types.Claim{
		BuyerUserID:  booking.UserID,
		SellerUserID: tour.AgencyID,
		Amount:       tour.TotalPrice,
		ClaimType:    "tour",
		ClaimDetails: fmt.Sprintf("Booking for tour %d from %s to %s", booking.TourID, tour.StartDate, tour.EndDate),
		Status:       "pending",
	}

	_, err = json.Marshal(claim)
	if err != nil {
		return nil, fmt.Errorf("error marshalling claim: %v", err)
	}

	/*claimID, err := s.messaging.PublishMessage(messaging.,claimData)
	if err != nil {
		return nil, fmt.Errorf("error sending claim to bank: %v", err)
	}*/

	var claimNum uint = 10
	booking.ClaimID = &claimNum
	booking.BookingStatus = "pending"
	booking.Confirmed = false

	newBooking, err := s.repository.CreateTourBooking(booking)
	if err != nil {
		return nil, fmt.Errorf("error saving tour booking: %v", err)
	}

	return newBooking, nil
}

func (s *TravelAgencyService) ConfirmTourBooking(bookingID uint) (*types.TourBooking, error) {
	booking, err := s.repository.ConfirmBooking(bookingID)
	if err != nil {
		return nil, fmt.Errorf("error confirming booking: %v", err)
	}

	if booking.ClaimID == nil {
		return nil, fmt.Errorf("no claimId associated with this booking")
	}

	bankServiceURL := fmt.Sprintf("http://bank-service:7070/api/v1/bank/process-confirmed-claim/%d", booking.ClaimID)
	req, err := http.NewRequest(http.MethodPost, bankServiceURL, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request to bank service: %v", err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending confirmation to bank service: %v", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Printf("error closing response body: %v", err)
		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("bank service returned status: %v", resp.Status)
	}

	log.Printf("Successfully confirmed claim with ID: %d in Bank Service", *booking.ClaimID)

	return booking, nil
}
