package get_balance

import (
	"github.com/emmanuelviniciusdev/imersao-fullcycle-event-driven-architecture/wallet-balance-api/internal/gateway"
	"time"
)

type GetBalanceOutputDTO struct {
	ID        int       `json:"id"`
	AccountID string    `json:"account_id"`
	Balance   float64   `json:"balance"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type GetBalanceUsecase struct {
	BalanceGateway gateway.BalanceGateway
}

func NewGetBalanceUsecase(balanceGateway gateway.BalanceGateway) *GetBalanceUsecase {
	return &GetBalanceUsecase{BalanceGateway: balanceGateway}
}

func (uc *GetBalanceUsecase) Execute(accountID string) (*GetBalanceOutputDTO, error) {
	balance, err := uc.BalanceGateway.FindByAccountID(accountID)

	if err != nil {
		return nil, err
	}

	output := &GetBalanceOutputDTO{
		ID:        balance.ID,
		AccountID: balance.AccountID,
		Balance:   balance.Balance,
		CreatedAt: balance.CreatedAt,
		UpdatedAt: balance.UpdatedAt,
	}

	return output, nil
}
