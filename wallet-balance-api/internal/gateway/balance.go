package gateway

import "github.com/emmanuelviniciusdev/imersao-fullcycle-event-driven-architecture/wallet-balance-api/internal/entity"

type BalanceGateway interface {
	Save(balance *entity.Balance) error
	FindByAccountID(accountID string) (*entity.Balance, error)
}
