-- name: CreateOwnerDetails :execresult
INSERT INTO owner_details (
  name,
  age
) VALUES (
  ?,?
);

-- name: GetOwner :execresult
SELECT * FROM owner_details WHERE name = ? LIMIT 1;