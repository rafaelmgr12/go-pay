package transactions

import (
	"context"

	"github.com/rafaelmgr12/go-pay/internal/domain/gateway"
)

type TransactionsUseCase struct {
	BalanceGateway gateway.BalanceGateway
}

func NewTransactionsUseCase(bg gateway.BalanceGateway) *TransactionsUseCase {
	return &TransactionsUseCase{
		BalanceGateway: bg,
	}
}

func (uc *TransactionsUseCase) GetAmountById(id string, ctx context.Context) (float64, error) {
	amount, err := uc.BalanceGateway.GetAmountById(ctx, id)
	if err != nil {
		return 0, err
	}

	return amount, nil
}
