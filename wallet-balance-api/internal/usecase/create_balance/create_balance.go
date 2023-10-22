package create_balance

import (
	"github.com/emmanuelviniciusdev/imersao-fullcycle-event-driven-architecture/wallet-balance-api/internal/entity"
	"github.com/emmanuelviniciusdev/imersao-fullcycle-event-driven-architecture/wallet-balance-api/internal/gateway"
)

type PayloadCreateBalanceEventDTO struct {
	AccountIDFrom        string  `json:"account_id_from"`
	AccountIDTo          string  `json:"account_id_to"`
	BalanceAccountIDFrom float64 `json:"balance_account_id_from"`
	BalanceAccountIDTo   float64 `json:"balance_account_id_to"`
}

type CreateBalanceEventDTO struct {
	Name    string                       `json:"Name"`
	Payload PayloadCreateBalanceEventDTO `json:"Payload"`
}

type CreateBalanceInputDTO struct {
	AccountID string  `json:"account_id"`
	Balance   float64 `json:"balance"`
}

type CreateBalanceUsecase struct {
	BalanceGateway gateway.BalanceGateway
}

func NewCreateBalanceUsecase(balanceGateway gateway.BalanceGateway) *CreateBalanceUsecase {
	return &CreateBalanceUsecase{
		BalanceGateway: balanceGateway,
	}
}

func (uc *CreateBalanceUsecase) Execute(input CreateBalanceInputDTO) {
	balance := entity.NewBalance(input.AccountID, input.Balance)

	err := uc.BalanceGateway.Save(balance)

	if err != nil {
		panic(err)
	}
}
