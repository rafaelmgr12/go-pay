-- name: GetBalanceByAccountID :one
select amount from balances where id = ?;