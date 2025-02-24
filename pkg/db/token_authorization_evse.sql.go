// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: token_authorization_evse.sql

package db

import (
	"context"
)

const listTokenAuthorizationEvses = `-- name: ListTokenAuthorizationEvses :many
SELECT e.id, e.location_id, e.uid, e.evse_id, e.identifier, e.status, e.floor_level, e.geom, e.geo_location_id, e.is_remote_capable, e.is_rfid_capable, e.physical_reference, e.last_updated FROM evses e
  INNER JOIN token_authorization_evses tae ON tae.evse_id = e.id
  WHERE tae.token_authorization_id = $1
  ORDER BY e.id
`

func (q *Queries) ListTokenAuthorizationEvses(ctx context.Context, tokenAuthorizationID int64) ([]Evse, error) {
	rows, err := q.db.QueryContext(ctx, listTokenAuthorizationEvses, tokenAuthorizationID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Evse
	for rows.Next() {
		var i Evse
		if err := rows.Scan(
			&i.ID,
			&i.LocationID,
			&i.Uid,
			&i.EvseID,
			&i.Identifier,
			&i.Status,
			&i.FloorLevel,
			&i.Geom,
			&i.GeoLocationID,
			&i.IsRemoteCapable,
			&i.IsRfidCapable,
			&i.PhysicalReference,
			&i.LastUpdated,
		); err != nil {
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

const setTokenAuthorizationEvse = `-- name: SetTokenAuthorizationEvse :exec
INSERT INTO token_authorization_evses (token_authorization_id, evse_id)
  VALUES ($1, $2)
`

type SetTokenAuthorizationEvseParams struct {
	TokenAuthorizationID int64 `db:"token_authorization_id" json:"tokenAuthorizationId"`
	EvseID               int64 `db:"evse_id" json:"evseId"`
}

func (q *Queries) SetTokenAuthorizationEvse(ctx context.Context, arg SetTokenAuthorizationEvseParams) error {
	_, err := q.db.ExecContext(ctx, setTokenAuthorizationEvse, arg.TokenAuthorizationID, arg.EvseID)
	return err
}
