// Code generated by sqlc. DO NOT EDIT.
// source: cdr_charging_period.sql

package db

import (
	"context"
)

const deleteCdrChargingPeriods = `-- name: DeleteCdrChargingPeriods :exec
DELETE FROM charging_periods cp
  USING cdr_charging_periods scp
  WHERE scp.charging_period_id == cp.id AND scp.cdr_id == $1
`

func (q *Queries) DeleteCdrChargingPeriods(ctx context.Context, cdrID int64) error {
	_, err := q.db.ExecContext(ctx, deleteCdrChargingPeriods, cdrID)
	return err
}

const listCdrChargingPeriods = `-- name: ListCdrChargingPeriods :many
SELECT cp.id, cp.start_date_time FROM charging_periods cp
  INNER JOIN cdr_charging_periods scp ON scp.charging_period_id == cp.id
  WHERE scp.cdr_id == $1
  ORDER BY cp.id
`

func (q *Queries) ListCdrChargingPeriods(ctx context.Context, cdrID int64) ([]ChargingPeriod, error) {
	rows, err := q.db.QueryContext(ctx, listCdrChargingPeriods, cdrID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ChargingPeriod
	for rows.Next() {
		var i ChargingPeriod
		if err := rows.Scan(&i.ID, &i.StartDateTime); err != nil {
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

const setCdrChargingPeriod = `-- name: SetCdrChargingPeriod :exec
INSERT INTO cdr_charging_periods (
    cdr_id, 
    charging_period_id
  ) VALUES ($1, $2)
`

type SetCdrChargingPeriodParams struct {
	CdrID            int64 `db:"cdr_id" json:"cdrID"`
	ChargingPeriodID int64 `db:"charging_period_id" json:"chargingPeriodID"`
}

func (q *Queries) SetCdrChargingPeriod(ctx context.Context, arg SetCdrChargingPeriodParams) error {
	_, err := q.db.ExecContext(ctx, setCdrChargingPeriod, arg.CdrID, arg.ChargingPeriodID)
	return err
}
