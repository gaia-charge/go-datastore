// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: charging_period.sql

package db

import (
	"context"
	"time"
)

const createChargingPeriod = `-- name: CreateChargingPeriod :one
INSERT INTO charging_periods (
    start_date_time
  ) VALUES ($1)
  RETURNING id, start_date_time
`

func (q *Queries) CreateChargingPeriod(ctx context.Context, startDateTime time.Time) (ChargingPeriod, error) {
	row := q.db.QueryRowContext(ctx, createChargingPeriod, startDateTime)
	var i ChargingPeriod
	err := row.Scan(&i.ID, &i.StartDateTime)
	return i, err
}
