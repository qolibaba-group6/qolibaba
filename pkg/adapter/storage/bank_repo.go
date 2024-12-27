package storage

import (
	"errors"
	"gorm.io/gorm"
	"qolibaba/pkg/adapter/storage/types"
)

type BankRepository struct {
	db *gorm.DB
}

func NewWalletRepository(db *gorm.DB) *BankRepository {
	return &BankRepository{
		db: db,
	}
}

// AddCardToWallet adds a new card to the wallet table
func (r *BankRepository) AddCardToWallet(userID uint, cardNumber string) (*types.Wallet, error) {
	var existingWallet types.Wallet
	if err := r.db.Where("card_number = ?", cardNumber).First(&existingWallet).Error; err == nil {
		return nil, errors.New("card number already exists")
	}
	newWallet := &types.Wallet{
		UserID:     &userID,
		CardNumber: cardNumber,
		Balance:    0.0,
		Role:       types.WalletRoleUser,
	}

	if err := r.db.Create(newWallet).Error; err != nil {
		return nil, err
	}

	return newWallet, nil
}
