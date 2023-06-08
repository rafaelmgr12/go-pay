-- name: GetBalanceByAccountID :one
select amount from balances where user_id = ?;

-- name: UpdateBalance :exec
UPDATE balances SET amount = ? WHERE user_id = ?;
