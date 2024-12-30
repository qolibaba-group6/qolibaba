
package service
import (
	"errors"
	"travel-booking-app/internal/ticket-service/model"
	"travel-booking-app/internal/ticket-service/repository"
	"github.com/google/uuid"
)
type TicketService interface {
	CreateTicket(userID, ticketType string, price float64, returnPolicy string) (string, error)
	GetTicketByID(id string) (*model.Ticket, error)
	GetTicketsByUserID(userID string) ([]model.Ticket, error)
	UpdateTicket(id, ticketType string, price float64, status, returnPolicy string) error
	DeleteTicket(id string) error
}
type ticketService struct {
	repo repository.TicketRepository
}
func NewTicketService(repo repository.TicketRepository) TicketService {
	return &ticketService{
		repo: repo,
	}
}
func (s *ticketService) CreateTicket(userID, ticketType string, price float64, returnPolicy string) (string, error) {
	if userID == "" || ticketType == "" || price <= 0 {
		return "", errors.New("invalid input data")
	}
	ticket := &model.Ticket{
		UserID:       uuid.MustParse(userID),
		Type:         ticketType,
		Price:        price,
		Status:       "active",
		ReturnPolicy: returnPolicy,
	}
	err := s.repo.CreateTicket(ticket)
	if err != nil {
		return "", err
	}
	return ticket.ID.String(), nil
}
func (s *ticketService) GetTicketByID(id string) (*model.Ticket, error) {
	return s.repo.GetTicketByID(id)
}
func (s *ticketService) GetTicketsByUserID(userID string) ([]model.Ticket, error) {
	return s.repo.GetTicketsByUserID(userID)
}
func (s *ticketService) UpdateTicket(id, ticketType string, price float64, status, returnPolicy string) error {
	ticket, err := s.repo.GetTicketByID(id)
	if err != nil {
		return err
	}
	if ticketType != "" {
		ticket.Type = ticketType
	}
	if price > 0 {
		ticket.Price = price
	}
	if status != "" {
		ticket.Status = status
	}
	if returnPolicy != "" {
		ticket.ReturnPolicy = returnPolicy
	}
	return s.repo.UpdateTicket(ticket)
}
func (s *ticketService) DeleteTicket(id string) error {
	return s.repo.DeleteTicket(id)
}
