package port

import "qolibaba/pkg/adapter/storage/types"

type Repo interface {
	CreateWallet(wallet *types.Wallet) (*types.Wallet, error)
}
