package port

import "qolibaba/pkg/adapter/storage/types"

type Service interface {
	CreateWallet(wallet *types.Wallet) (*types.Wallet, error)
	ChargeWallet(walletID uint, amount float64) (*types.Wallet, *types.Transaction, error)
	ProcessUnconfirmedClaim(claim *types.Claim) (*types.Claim, error)
	ProcessConfirmedClaim(claimID uint) (*types.Claim, error)
}
