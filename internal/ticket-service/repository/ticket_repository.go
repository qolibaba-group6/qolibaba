// internal/ticket-service/repository/ticket_repository.go
package repository

import (
	"travel-booking-app/internal/ticket-service/model"

	"travel-booking-app/internal/database"
)

type TicketRepository interface {
	CreateTicket(ticket *model.Ticket) error
	GetTicketByID(id string) (*model.Ticket, error)
	GetTicketsByUserID(userID string) ([]model.Ticket, error)
	UpdateTicket(ticket *model.Ticket) error
	DeleteTicket(id string) error
}

type ticketRepository struct{}

func NewTicketRepository() TicketRepository {
	return &ticketRepository{}
}

func (r *ticketRepository) CreateTicket(ticket *model.Ticket) error {
	return database.DB.Create(ticket).Error
}

func (r *ticketRepository) GetTicketByID(id string) (*model.Ticket, error) {
	var ticket model.Ticket
	err := database.DB.First(&ticket, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &ticket, nil
}

func (r *ticketRepository) GetTicketsByUserID(userID string) ([]model.Ticket, error) {
	var tickets []model.Ticket
	err := database.DB.Where("user_id = ?", userID).Find(&tickets).Error
	if err != nil {
		return nil, err
	}
	return tickets, nil
}

func (r *ticketRepository) UpdateTicket(ticket *model.Ticket) error {
	return database.DB.Save(ticket).Error
}

func (r *ticketRepository) DeleteTicket(id string) error {
	return database.DB.Delete(&model.Ticket{}, "id = ?", id).Error
}
