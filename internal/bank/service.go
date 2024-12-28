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
