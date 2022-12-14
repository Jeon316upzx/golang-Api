// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: owner_details.sql

package bankdb

import (
	"context"
	"database/sql"
)

const createOwnerDetails = `-- name: CreateOwnerDetails :execresult
INSERT INTO owner_details (
  name,
  age
) VALUES (
  ?,?
)
`

type CreateOwnerDetailsParams struct {
	Name string        `json:"name"`
	Age  sql.NullInt32 `json:"age"`
}

func (q *Queries) CreateOwnerDetails(ctx context.Context, arg CreateOwnerDetailsParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, createOwnerDetails, arg.Name, arg.Age)
}

const getOwner = `-- name: GetOwner :one
SELECT id, name, age, image, country FROM owner_details WHERE name = ? LIMIT 1
`

func (q *Queries) GetOwner(ctx context.Context, name string) (OwnerDetail, error) {
	row := q.db.QueryRowContext(ctx, getOwner, name)
	var i OwnerDetail
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Age,
		&i.Image,
		&i.Country,
	)
	return i, err
}
