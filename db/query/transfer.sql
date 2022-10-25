-- name: CreateTransfer :execresult
INSERT INTO transfer (amount, to_account, from_account) VALUES (?,?,?);

-- name: GetTransfers :many
SELECT * FROM transfer;

-- name: GetTransferTo :many
SELECT * FROM transfer WHERE to_account = ?;

-- name: GetTransferFrom :many
SELECT * FROM transfer WHERE from_account = ?;

-- name: GetTransfer :one
SELECT * FROM transfer WHERE id = ? LIMIT 1; 