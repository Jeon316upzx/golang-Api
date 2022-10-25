-- name: CreateAccount :execresult
INSERT INTO accounts (balance,currency, owner) VALUES (?,?,?);

-- name: GetAccount :many
SELECT * FROM accounts WHERE owner = ?;

-- name: GetAccountById :one
SELECT * FROM accounts WHERE owner = ? LIMIT 1;

-- name: UpdateAccount :execresult
UPDATE accounts SET balance = ? WHERE owner = ?;