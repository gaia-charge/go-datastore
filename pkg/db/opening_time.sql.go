// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: opening_time.sql

package db

import (
	"context"
)

const createOpeningTime = `-- name: CreateOpeningTime :one
INSERT INTO opening_times (
    twentyfourseven
  ) VALUES ($1)
  RETURNING id, twentyfourseven
`

func (q *Queries) CreateOpeningTime(ctx context.Context, twentyfourseven bool) (OpeningTime, error) {
	row := q.db.QueryRowContext(ctx, createOpeningTime, twentyfourseven)
	var i OpeningTime
	err := row.Scan(&i.ID, &i.Twentyfourseven)
	return i, err
}

const deleteOpeningTime = `-- name: DeleteOpeningTime :exec
DELETE FROM opening_times
  WHERE id = $1
`

func (q *Queries) DeleteOpeningTime(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteOpeningTime, id)
	return err
}

const getOpeningTime = `-- name: GetOpeningTime :one
SELECT id, twentyfourseven FROM opening_times
  WHERE id = $1
`

func (q *Queries) GetOpeningTime(ctx context.Context, id int64) (OpeningTime, error) {
	row := q.db.QueryRowContext(ctx, getOpeningTime, id)
	var i OpeningTime
	err := row.Scan(&i.ID, &i.Twentyfourseven)
	return i, err
}

const updateOpeningTime = `-- name: UpdateOpeningTime :one
UPDATE opening_times 
  SET twentyfourseven = $2
  WHERE id = $1
  RETURNING id, twentyfourseven
`

type UpdateOpeningTimeParams struct {
	ID              int64 `db:"id" json:"id"`
	Twentyfourseven bool  `db:"twentyfourseven" json:"twentyfourseven"`
}

func (q *Queries) UpdateOpeningTime(ctx context.Context, arg UpdateOpeningTimeParams) (OpeningTime, error) {
	row := q.db.QueryRowContext(ctx, updateOpeningTime, arg.ID, arg.Twentyfourseven)
	var i OpeningTime
	err := row.Scan(&i.ID, &i.Twentyfourseven)
	return i, err
}
