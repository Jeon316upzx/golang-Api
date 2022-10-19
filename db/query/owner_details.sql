-- name: CreateOwnerDetails :exec
INSERT INTO owner_details (
  name,
  age
) VALUES (
  ?,?
);

-- name: GetOwner :exec
SELECT * FROM owner_details WHERE name = ? LIMIT 1;