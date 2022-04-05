// Code generated by sqlc. DO NOT EDIT.
// source: evse_parking_restriction.sql

package db

import (
	"context"
)

const listEvseParkingRestrictions = `-- name: ListEvseParkingRestrictions :many
SELECT pr.id, pr.text, pr.description FROM parking_restrictions pr
  INNER JOIN evse_parking_restrictions ep ON ep.parking_restriction_id = pr.id
  WHERE ep.evse_id = $1
  ORDER BY pr.id
`

func (q *Queries) ListEvseParkingRestrictions(ctx context.Context, evseID int64) ([]ParkingRestriction, error) {
	rows, err := q.db.QueryContext(ctx, listEvseParkingRestrictions, evseID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ParkingRestriction
	for rows.Next() {
		var i ParkingRestriction
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

const setEvseParkingRestriction = `-- name: SetEvseParkingRestriction :exec
INSERT INTO evse_parking_restrictions (evse_id, parking_restriction_id)
  VALUES ($1, $2)
`

type SetEvseParkingRestrictionParams struct {
	EvseID               int64 `db:"evse_id" json:"evseID"`
	ParkingRestrictionID int64 `db:"parking_restriction_id" json:"parkingRestrictionID"`
}

func (q *Queries) SetEvseParkingRestriction(ctx context.Context, arg SetEvseParkingRestrictionParams) error {
	_, err := q.db.ExecContext(ctx, setEvseParkingRestriction, arg.EvseID, arg.ParkingRestrictionID)
	return err
}

const unsetEvseParkingRestrictions = `-- name: UnsetEvseParkingRestrictions :exec
DELETE FROM evse_parking_restrictions ep
  WHERE ep.evse_id = $1
`

func (q *Queries) UnsetEvseParkingRestrictions(ctx context.Context, evseID int64) error {
	_, err := q.db.ExecContext(ctx, unsetEvseParkingRestrictions, evseID)
	return err
}
