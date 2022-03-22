// Code generated by sqlc. DO NOT EDIT.
// source: element.sql

package db

import (
	"context"
	"database/sql"
)

const createElement = `-- name: CreateElement :one
INSERT INTO elements (
    tariff_id,
    restriction_id
  ) VALUES ($1, $2)
  RETURNING id, tariff_id, restriction_id
`

type CreateElementParams struct {
	TariffID      int64         `db:"tariff_id" json:"tariffID"`
	RestrictionID sql.NullInt64 `db:"restriction_id" json:"restrictionID"`
}

func (q *Queries) CreateElement(ctx context.Context, arg CreateElementParams) (Element, error) {
	row := q.db.QueryRowContext(ctx, createElement, arg.TariffID, arg.RestrictionID)
	var i Element
	err := row.Scan(&i.ID, &i.TariffID, &i.RestrictionID)
	return i, err
}

const deleteElements = `-- name: DeleteElements :exec
DELETE FROM elements
  WHERE tariff_id = $1
`

func (q *Queries) DeleteElements(ctx context.Context, tariffID int64) error {
	_, err := q.db.ExecContext(ctx, deleteElements, tariffID)
	return err
}

const listElements = `-- name: ListElements :many
SELECT id, tariff_id, restriction_id FROM elements
  WHERE tariff_id = $1
  ORDER BY id
`

func (q *Queries) ListElements(ctx context.Context, tariffID int64) ([]Element, error) {
	rows, err := q.db.QueryContext(ctx, listElements, tariffID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Element
	for rows.Next() {
		var i Element
		if err := rows.Scan(&i.ID, &i.TariffID, &i.RestrictionID); err != nil {
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
