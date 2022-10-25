-- name: CreateOwner :execresult
INSERT INTO owner (
    owner_details
) VALUES (
  ?
);

-- name: GetProfile :one
SELECT * FROM owner WHERE owner_details = ?;


