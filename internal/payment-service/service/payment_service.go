package service
import (
	"errors"
	"github.com/google/uuid"
	"travel-booking-app/internal/payment-service/model"
	"travel-booking-app/internal/payment-service/repository"
)
type PaymentService interface {
	CreatePayment(userID string, amount float64) (string, error)
	GetPaymentByID(id string) (*model.Payment, error)
	GetPaymentsByUserID(userID string) ([]model.Payment, error)
	UpdatePaymentStatus(id, status string) error
	DeletePayment(id string) error
}
type paymentService struct {
	repo repository.PaymentRepository
}
func NewPaymentService(repo repository.PaymentRepository) PaymentService {
	return &paymentService{repo: repo}
}
func (s *paymentService) CreatePayment(userID string, amount float64) (string, error) {
	if userID == "" || amount <= 0 {
		return "", errors.New("invalid input data")
	}
	transactionID := uuid.New().String()
	payment := &model.Payment{
		UserID:        uuid.MustParse(userID),
		Amount:        amount,
		Status:        "pending",
		TransactionID: transactionID,
	}
	err := s.repo.CreatePayment(payment)
	if err != nil {
		return "", err
	}
	return payment.ID.String(), nil
}
func (s *paymentService) GetPaymentByID(id string) (*model.Payment, error) {
	return s.repo.GetPaymentByID(id)
}
func (s *paymentService) GetPaymentsByUserID(userID string) ([]model.Payment, error) {
	return s.repo.GetPaymentsByUserID(userID)
}
func (s *paymentService) UpdatePaymentStatus(id, status string) error {
	payment, err := s.repo.GetPaymentByID(id)
	if err != nil {
		return err
	}
	payment.Status = status
	return s.repo.UpdatePayment(payment)
}
func (s *paymentService) DeletePayment(id string) error {
	return s.repo.DeletePayment(id)
}
