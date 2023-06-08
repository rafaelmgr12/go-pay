package gateway

import "context"

type BalanceGateway interface {
	GetAmountById(ctx context.Context, id string) (float64, error)
}
