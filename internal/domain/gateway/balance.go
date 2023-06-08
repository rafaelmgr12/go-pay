package gateway

import "context"

type BalanceGateway interface {
	GetAmountById(ctx context.Context, id string) (float64, error)
	Transfer(debtorId string, creditorId string, amount float64, ctx context.Context) error
}
