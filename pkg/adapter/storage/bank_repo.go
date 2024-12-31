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

// UpdateWalletBalance Updates Wallet Balance
func (r *bankRepo) UpdateWalletBalance(walletID uint, amount float64) (*types.Wallet, error) {
	var wallet types.Wallet

	// Retrieve the wallet by ID
	if err := r.db.First(&wallet, walletID).Error; err != nil {
		return nil, fmt.Errorf("wallet not found: %v", err)
	}

	// Update the wallet balance
	wallet.Balance += amount

	// Save the updated wallet
	if err := r.db.Save(&wallet).Error; err != nil {
		return nil, fmt.Errorf("error updating wallet balance: %v", err)
	}

	return &wallet, nil
}

// CreateTransaction Saves a Transaction
func (r *bankRepo) CreateTransaction(walletID uint, amount float64, transactionType, description string) (*types.Transaction, error) {

	transaction := &types.Transaction{
		WalletID:        walletID,
		Amount:          amount,
		TransactionType: transactionType,
		Status:          string(types.StatusCompleted),
		Description:     description,
	}
	if err := r.db.Create(transaction).Error; err != nil {
		return nil, fmt.Errorf("error saving transaction: %v", err)
	}

	return transaction, nil
}

// Withdrawal transfers money from one wallet to another after validation.
func (r *bankRepo) Withdrawal(fromWalletID, toWalletID uint, amount float64) error {
	var fromWallet, toWallet types.Wallet
	if err := r.db.First(&fromWallet, fromWalletID).Error; err != nil {
		return fmt.Errorf("source wallet not found: %v", err)
	}

	if fromWallet.Balance < amount {
		return fmt.Errorf("insufficient funds in source wallet")
	}

	if err := r.db.First(&toWallet, toWalletID).Error; err != nil {
		return fmt.Errorf("destination wallet not found: %v", err)
	}

	err := r.db.Transaction(func(tx *gorm.DB) error {
		fromWallet.Balance -= amount
		if err := tx.Save(&fromWallet).Error; err != nil {
			return fmt.Errorf("error updating source wallet balance: %v", err)
		}
		toWallet.Balance += amount
		if err := tx.Save(&toWallet).Error; err != nil {
			return fmt.Errorf("error updating destination wallet balance: %v", err)
		}

		return nil
	})

	if err != nil {
		return fmt.Errorf("error transferring funds: %v", err)
	}

	return nil
}

// UpsertClaim saves or updates a claim based on its ID and sets the provided status.
func (r *bankRepo) UpsertClaim(claim *types.Claim, status string) (*types.Claim, error) {
	claim.Status = status
	if err := r.db.Save(claim).Error; err != nil {
		return nil, fmt.Errorf("error saving or updating claim: %v", err)
	}

	return claim, nil
}

// GetClaimByID retrieves a claim by its ID.
func (r *bankRepo) GetClaimByID(claimID uint) (*types.Claim, error) {
	var claim types.Claim
	if err := r.db.Preload("Transactions").First(&claim, claimID).Error; err != nil {
		return nil, fmt.Errorf("claim not found: %v", err)
	}
	return &claim, nil
}
