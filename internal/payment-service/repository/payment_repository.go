package repository

import (
	"travel-booking-app/internal/database"
	"travel-booking-app/internal/payment-service/model"
)

type PaymentRepository interface {
	CreatePayment(payment *model.Payment) error
	GetPaymentByID(id string) (*model.Payment, error)
	GetPaymentsByUserID(userID string) ([]model.Payment, error)
	UpdatePayment(payment *model.Payment) error
	DeletePayment(id string) error
}
type paymentRepository struct{}

func NewPaymentRepository() PaymentRepository {
	return &paymentRepository{}
}
func (r *paymentRepository) CreatePayment(payment *model.Payment) error {
	return database.DB.Create(payment).Error
}
func (r *paymentRepository) GetPaymentByID(id string) (*model.Payment, error) {
	var payment model.Payment
	if err := database.DB.First(&payment, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &payment, nil
}
func (r *paymentRepository) GetPaymentsByUserID(userID string) ([]model.Payment, error) {
	var payments []model.Payment
	if err := database.DB.Where("user_id = ?", userID).Find(&payments).Error; err != nil {
		return nil, err
	}
	return payments, nil
}
func (r *paymentRepository) UpdatePayment(payment *model.Payment) error {
	return database.DB.Save(payment).Error
}
func (r *paymentRepository) DeletePayment(id string) error {
	return database.DB.Delete(&model.Payment{}, "id = ?", id).Error
}
