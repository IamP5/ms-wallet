package gateway

import "github.com/IamP5/ms-wallet/wallet-core/internal/entity"

type TransactionGateway interface {
	Create(transaction *entity.Transaction) error
}
