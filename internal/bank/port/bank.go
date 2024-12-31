package port

import "qolibaba/pkg/adapter/storage/types"

type Repo interface {
	CreateWallet(wallet *types.Wallet) (*types.Wallet, error)
	UpdateWalletBalance(walletID uint, amount float64) (*types.Wallet, error)
	CreateTransaction(walletID uint, amount float64, transactionType, description string) (*types.Transaction, error)
	Withdrawal(fromWalletID, toWalletID uint, amount float64) error
	UpsertClaim(claim *types.Claim, status string) (*types.Claim, error)
	GetClaimByID(claimID uint) (*types.Claim, error)
}
