-- name: GetBalanceByAccountID :one
select amount from balances where user_id = ?;