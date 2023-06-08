package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/rafaelmgr12/go-pay/internal/infra/db"
)

type ChatRepositoryMySQL struct {
	DB      *sql.DB
	Queries *db.Queries
}

func NewChatRepositoryMySQL(dbt *sql.DB) *ChatRepositoryMySQL {
	return &ChatRepositoryMySQL{
		DB:      dbt,
		Queries: db.New(dbt),
	}
}

func (r *ChatRepositoryMySQL) GetAmountById(ctx context.Context, id string) (float64, error) {
	res, err := r.Queries.GetBalanceByAccountID(ctx, id)
	if err != nil {
		return -1, errors.New("error getting balance")
	}
	return res, nil

}

func (r *ChatRepositoryMySQL) Transfer(debtorId string, creditorId string, amount float64, ctx context.Context) error {
	tx, err := r.DB.BeginTx(ctx, nil)
	if err != nil {
		return errors.New("error starting transaction")
	}
	defer tx.Rollback()

	debtorIdAmount, err := r.Queries.GetBalanceByAccountID(ctx, debtorId)
	if err != nil {
		return errors.New("error getting debtor balance")
	}
	creditorIdAmount, err := r.Queries.GetBalanceByAccountID(ctx, creditorId)
	if err != nil {
		return errors.New("error getting creditor balance")
	}

	paramsDebtorID := db.UpdateBalanceParams{
		Amount: debtorIdAmount - amount,
		UserID: debtorId,
	}

	paramsCreditorID := db.UpdateBalanceParams{
		Amount: creditorIdAmount + amount,
		UserID: creditorId,
	}

	err = r.Queries.UpdateBalance(ctx, paramsDebtorID)
	if err != nil {
		return errors.New("error updating debtor balance")
	}

	err = r.Queries.UpdateBalance(ctx, paramsCreditorID)
	if err != nil {
		return errors.New("error updating creditor balance")
	}

	return tx.Commit()
}
