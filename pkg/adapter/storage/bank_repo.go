package storage

import (
	"fmt"
	"gorm.io/gorm"
	"qolibaba/internal/bank/port"
	"qolibaba/pkg/adapter/storage/types"
)

// bankRepo implements the WalletRepository interface
type bankRepo struct {
	db *gorm.DB
}

// NewBankRepo creates a new instance of walletRepo
func NewBankRepo(db *gorm.DB) port.Repo {
	return &bankRepo{
		db: db,
	}
}

// CreateWallet creates a new wallet in the database
func (r *bankRepo) CreateWallet(wallet *types.Wallet) (*types.Wallet, error) {

	if wallet.UserID == nil {
		return nil, fmt.Errorf("user ID must be provided")
	}
	if wallet.Role == "" {
		wallet.Role = "user"
	}
	var existingWallet types.Wallet
	if err := r.db.Where("user_id = ?", wallet.UserID).First(&existingWallet).Error; err == nil {
		return nil, fmt.Errorf("a wallet already exists for the user")
	}

	if err := r.db.Create(wallet).Error; err != nil {
		return nil, fmt.Errorf("error creating wallet: %v", err)
	}

	return wallet, nil
}
