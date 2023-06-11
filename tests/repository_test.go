package tests

import "context"

type BalanceRepositoryMock struct {
	GetAmountByIdFunc func(ctx context.Context, id string) (float64, error)
	TransferFunc      func(debtorId string, creditorId string, amount float64, ctx context.Context) error
}

func (m *BalanceRepositoryMock) GetAmountById(ctx context.Context, id string) (float64, error) {
	if m.GetAmountByIdFunc != nil {
		return m.GetAmountByIdFunc(ctx, id)
	}
	return 0, nil // Implemente um comportamento padrão ou retorne valores de erro personalizados, se necessário
}

func (m *BalanceRepositoryMock) Transfer(debtorId string, creditorId string, amount float64, ctx context.Context) error {
	if m.TransferFunc != nil {
		return m.TransferFunc(debtorId, creditorId, amount, ctx)
	}
	return nil // Implemente um comportamento padrão ou retorne valores de erro personalizados, se necessário
}
