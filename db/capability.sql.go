// Code generated by sqlc. DO NOT EDIT.
// source: capability.sql

package db

import (
	"context"
)

const getCapability = `-- name: GetCapability :one
SELECT id, text, description FROM capabilities
  WHERE id = $1
`

func (q *Queries) GetCapability(ctx context.Context, id int64) (Capability, error) {
	row := q.db.QueryRowContext(ctx, getCapability, id)
	var i Capability
	err := row.Scan(&i.ID, &i.Text, &i.Description)
	return i, err
}

const listCapabilities = `-- name: ListCapabilities :many
SELECT id, text, description FROM capabilities
  ORDER BY id
`

func (q *Queries) ListCapabilities(ctx context.Context) ([]Capability, error) {
	rows, err := q.db.QueryContext(ctx, listCapabilities)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Capability
	for rows.Next() {
		var i Capability
		if err := rows.Scan(&i.ID, &i.Text, &i.Description); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
