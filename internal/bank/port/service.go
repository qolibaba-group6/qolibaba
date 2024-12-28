package port

import "qolibaba/pkg/adapter/storage/types"

type Service interface {
	CreateWallet(wallet *types.Wallet) (*types.Wallet, error)
}
