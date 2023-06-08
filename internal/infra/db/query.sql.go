// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: query.sql

package db

import (
	"context"
	"database/sql"
)

const getBalanceByAccountID = `-- name: GetBalanceByAccountID :one
select amount from balances where id = $1
`

func (q *Queries) GetBalanceByAccountID(ctx context.Context) (sql.NullString, error) {
	row := q.db.QueryRowContext(ctx, getBalanceByAccountID)
	var amount sql.NullString
	err := row.Scan(&amount)
	return amount, err
}