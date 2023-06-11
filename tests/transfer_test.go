package tests

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/rafaelmgr12/go-pay/internal/usecase/transactions"
)

func TestGetAmountById(t *testing.T) {
	mock := &BalanceRepositoryMock{
		GetAmountByIdFunc: func(ctx context.Context, id string) (float64, error) {
			if id == "123" {
				return 10.5, nil
			}
			return 0, errors.New("ID não encontrado")
		},
	}
	service := transactions.NewTransactionsUseCase(mock)
	result, err := service.GetAmountById("123", context.Background())
	if err != nil {
		t.Errorf("Erro inesperado: %s", err.Error())
	}
	if result != 10.5 {
		t.Errorf("Resultado incorreto. Esperado: %f, Recebido: %f", 10.5, result)
	}
}

func TestTransfer(t *testing.T) {
	mock := &BalanceRepositoryMock{
		TransferFunc: func(debtorId string, creditorId string, amount float64, ctx context.Context) error {

			if debtorId != "123" {
				return errors.New("ID devedor incorreto")
			}
			if creditorId != "456" {
				return errors.New("ID credor incorreto")
			}
			if amount != 100.0 {
				return errors.New("Valor de transferência incorreto")
			}
			return nil
		},

		GetAmountByIdFunc: func(ctx context.Context, id string) (float64, error) {
			if id == "123" {
				return 500.0, nil
			}
			return 0, errors.New("ID não encontrado")
		},
	}

	service := transactions.NewTransactionsUseCase(mock)

	err := service.Transfer("123", "456", 100.0, context.Background())

	if err != nil {
		t.Errorf("Erro inesperado: %s", err.Error())
	}
}

func TestTransfefWithInsufficientsAmount(t *testing.T) {
	mock := &BalanceRepositoryMock{
		GetAmountByIdFunc: func(ctx context.Context, id string) (float64, error) {
			// Implemente o comportamento do mock para o teste
			// Retorne o saldo do devedor como 500.0
			if id == "123" {
				return 500.0, nil
			}
			return 0, errors.New("ID não encontrado")
		},
	}

	service := transactions.NewTransactionsUseCase(mock)

	err := service.Transfer("123", "456", 600.0, context.Background())
	fmt.Println(err)
	if err == nil {
		t.Errorf("Erro esperado não ocorreu")
	}
}
