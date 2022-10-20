-- name: CreateAccount :execresult
INSERT INTO accounts (balance,currency, owner) VALUES (?,?,?);

-- name: GetAccount :many
SELECT * FROM accounts WHERE owner = ?;

-- name: UpdateAccount :execresult
UPDATE accounts SET balance = ? WHERE owner = ?;