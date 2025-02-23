// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: price_component_rounding.sql

package db

import (
	"context"
)

const createPriceComponentRounding = `-- name: CreatePriceComponentRounding :one
INSERT INTO price_component_roundings (
    granularity,
    rule
  ) VALUES ($1, $2)
  RETURNING id, granularity, rule
`

type CreatePriceComponentRoundingParams struct {
	Granularity RoundingGranularity `db:"granularity" json:"granularity"`
	Rule        RoundingRule        `db:"rule" json:"rule"`
}

func (q *Queries) CreatePriceComponentRounding(ctx context.Context, arg CreatePriceComponentRoundingParams) (PriceComponentRounding, error) {
	row := q.db.QueryRowContext(ctx, createPriceComponentRounding, arg.Granularity, arg.Rule)
	var i PriceComponentRounding
	err := row.Scan(&i.ID, &i.Granularity, &i.Rule)
	return i, err
}

const deletePriceComponentRoundings = `-- name: DeletePriceComponentRoundings :exec
DELETE FROM price_component_roundings pcr
  USING elements e, price_components pc
  WHERE e.tariff_id = $1 AND e.id = pc.element_id AND 
    (pc.price_rounding_id = pcr.id OR pc.step_rounding_id = pcr.id)
`

func (q *Queries) DeletePriceComponentRoundings(ctx context.Context, tariffID int64) error {
	_, err := q.db.ExecContext(ctx, deletePriceComponentRoundings, tariffID)
	return err
}

const getPriceComponentRounding = `-- name: GetPriceComponentRounding :one
SELECT id, granularity, rule FROM price_component_roundings
  WHERE id = $1
`

func (q *Queries) GetPriceComponentRounding(ctx context.Context, id int64) (PriceComponentRounding, error) {
	row := q.db.QueryRowContext(ctx, getPriceComponentRounding, id)
	var i PriceComponentRounding
	err := row.Scan(&i.ID, &i.Granularity, &i.Rule)
	return i, err
}
