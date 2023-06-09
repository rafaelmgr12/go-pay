package transactions

import (
	"context"
	"errors"

	"github.com/rafaelmgr12/go-pay/internal/domain/gateway"
)

type TransactionsUseCase struct {
	BalanceGateway gateway.BalanceGateway
}

type TransferInputDTO struct {
	DebtorId   string  `json:"debtor_id"`
	CreditorId string  `json:"creditor_id"`
	Amount     float64 `json:"amount"`
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

func (uc *TransactionsUseCase) Transfer(debtorId string, creditorId string, amount float64, ctx context.Context) error {
	errValidate := uc.validate(creditorId, debtorId, amount, ctx)
	if errValidate != nil {
		return errValidate
	}
	err := uc.BalanceGateway.Transfer(debtorId, creditorId, amount, ctx)
	if err != nil {
		return err
	}

	return nil
}

func (uc *TransactionsUseCase) validate(creditorId string, debtorId string, amount float64, ctx context.Context) error {
	if creditorId == debtorId {
		return errors.New("creditor_id and debtor_id must be different")
	}

	if amount <= 0 {
		return errors.New("amount must be greater than 0")
	}

	debtorAmount, err := uc.BalanceGateway.GetAmountById(ctx, debtorId)
	if err != nil {
		return err
	}

	if debtorAmount < amount {
		return errors.New("debtor does not have enough balance")
	}

	return nil
}
