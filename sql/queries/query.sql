-- name: GetBalanceByAccountID :one
select amount from balances where id = $1;