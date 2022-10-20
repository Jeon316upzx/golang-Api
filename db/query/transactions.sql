-- name: CreateTransaction :execresult
INSERT INTO transactions (type, amount, account) VALUES (?,?,?);

-- name: GetTransactions :many
SELECT * FROM transactions;

-- name: GetTransactionBy :many
SELECT * FROM transactions WHERE account = ?;

-- name: GetTransaction :execresult
SELECT * FROM transactions WHERE id = ? LIMIT 1;