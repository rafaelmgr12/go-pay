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
