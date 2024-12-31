package bank

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"qolibaba/internal/bank/port"
	"qolibaba/pkg/adapter/storage/types"
)

type service struct {
	bankRepo port.Repo
	validate *validator.Validate
}

func NewService(repo port.Repo) port.Service {
	return &service{
		bankRepo: repo,
		validate: validator.New(),
	}
}

// CreateWallet creates a new wallet using the wallet repository
func (s *service) CreateWallet(wallet *types.Wallet) (*types.Wallet, error) {
	if wallet.Role == "" {
		wallet.Role = "user"
	}

	newWallet, err := s.bankRepo.CreateWallet(wallet)
	if err != nil {
		return nil, fmt.Errorf("error creating wallet: %v", err)
	}

	return newWallet, nil
}

// ChargeWallet charges the wallet and stores the transaction
func (s *service) ChargeWallet(walletID uint, amount float64) (*types.Wallet, *types.Transaction, error) {
	if amount <= 0 {
		return nil, nil, fmt.Errorf("amount must be greater than zero")
	}
	wallet, err := s.bankRepo.UpdateWalletBalance(walletID, amount)
	if err != nil {
		return nil, nil, fmt.Errorf("error charging the wallet: %v", err)
	}
	transaction, err := s.bankRepo.CreateTransaction(walletID, amount, string(types.TransactionTypeDeposit), "Charging the wallet by user.")
	if err != nil {
		return nil, nil, fmt.Errorf("error creating transaction: %v", err)
	}
	return wallet, transaction, nil
}

func (s *service) ProcessUnconfirmedClaim(claim *types.Claim) (*types.Claim, error) {
	if claim.Amount <= 0 {
		return nil, fmt.Errorf("amount must be greater than zero")
	}
	if claim.BuyerUserID == 0 || claim.SellerUserID == 0 {
		return nil, fmt.Errorf("both buyer and seller user IDs must be provided")
	}

	const bankWalletID = 1

	err := s.bankRepo.Withdrawal(claim.BuyerUserID, bankWalletID, claim.Amount)
	if err != nil {
		failedClaim, _ := s.handleClaim(claim, string(types.StatusFailed), fmt.Errorf("withdrawal error: %v", err))
		return failedClaim, nil
	}

	withdrawTransaction, err := s.bankRepo.CreateTransaction(claim.BuyerUserID, claim.Amount, string(types.TransactionTypeWithdrawal), "claim withdrawal to bank")
	if err != nil {
		failedClaim, _ := s.handleClaim(claim, string(types.StatusFailed), fmt.Errorf("withdrawal error: %v", err))
		return failedClaim, nil
	}

	depositTransaction, err := s.bankRepo.CreateTransaction(bankWalletID, claim.Amount, string(types.TransactionTypeDeposit), "claim deposit from buyer")
	if err != nil {
		failedClaim, _ := s.handleClaim(claim, string(types.StatusFailed), fmt.Errorf("withdrawal error: %v", err))
		return failedClaim, nil
	}

	claim.Transactions = []types.Transaction{*withdrawTransaction, *depositTransaction}
	claim.Status = string(types.StatusPending)
	updatedClaim, err := s.handleClaim(claim, string(types.StatusFailed), fmt.Errorf("withdrawal error: %v", err))
	if err != nil {
		return nil, fmt.Errorf("failed to save claim with pending status: %v", err)
	}
	return updatedClaim, nil
}

func (s *service) ProcessConfirmedClaim(claimID uint) (*types.Claim, error) {
	claim, err := s.bankRepo.GetClaimByID(claimID)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve claim: %v", err)
	}

	if claim.Amount <= 0 {
		return s.handleClaim(claim, string(types.StatusFailed), fmt.Errorf("claim amount must be greater than zero"))
	}
	if claim.BuyerUserID == 0 || claim.SellerUserID == 0 {
		return s.handleClaim(claim, string(types.StatusFailed), fmt.Errorf("buyer and seller user IDs must be valid"))
	}

	const bankWalletID = 1

	err = s.bankRepo.Withdrawal(bankWalletID, claim.SellerUserID, claim.Amount)
	if err != nil {
		return s.handleClaim(claim, string(types.StatusFailed), fmt.Errorf("withdrawal error: %v", err))
	}

	withdrawTransaction, err := s.bankRepo.CreateTransaction(bankWalletID, claim.Amount, string(types.TransactionTypeWithdrawal), "claim withdrawal to seller")
	if err != nil {
		return s.handleClaim(claim, string(types.StatusFailed), fmt.Errorf("transaction creation error: %v", err))
	}

	depositTransaction, err := s.bankRepo.CreateTransaction(claim.SellerUserID, claim.Amount, string(types.TransactionTypeDeposit), "claim deposit from bank")
	if err != nil {
		return s.handleClaim(claim, string(types.StatusFailed), fmt.Errorf("transaction creation error: %v", err))
	}

	claim.Transactions = []types.Transaction{*withdrawTransaction, *depositTransaction}
	updatedClaim, err := s.handleClaim(claim, string(types.StatusPaid), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to update claim status: %v", err)
	}

	return updatedClaim, nil
}

func (s *service) handleClaim(claim *types.Claim, status string, failureReason error) (*types.Claim, error) {
	claim.Status = status
	claim, err := s.bankRepo.UpsertClaim(claim, claim.Status)
	if err != nil {
		return nil, fmt.Errorf("failed to save claim")
	}
	return claim, failureReason
}
