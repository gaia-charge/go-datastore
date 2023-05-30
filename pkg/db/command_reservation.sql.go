// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: command_reservation.sql

package db

import (
	"context"
	"database/sql"
	"time"
)

const createCommandReservation = `-- name: CreateCommandReservation :one
INSERT INTO command_reservations (
    status,
    token_id,
    expiry_date,
    location_id,
    evse_uid,
    last_updated
  ) VALUES ($1, $2, $3, $4, $5, $6)
  RETURNING id, status, token_id, expiry_date, reservation_id, location_id, evse_uid, last_updated
`

type CreateCommandReservationParams struct {
	Status      CommandResponseType `db:"status" json:"status"`
	TokenID     int64               `db:"token_id" json:"tokenID"`
	ExpiryDate  time.Time           `db:"expiry_date" json:"expiryDate"`
	LocationID  string              `db:"location_id" json:"locationID"`
	EvseUid     sql.NullString      `db:"evse_uid" json:"evseUid"`
	LastUpdated time.Time           `db:"last_updated" json:"lastUpdated"`
}

func (q *Queries) CreateCommandReservation(ctx context.Context, arg CreateCommandReservationParams) (CommandReservation, error) {
	row := q.db.QueryRowContext(ctx, createCommandReservation,
		arg.Status,
		arg.TokenID,
		arg.ExpiryDate,
		arg.LocationID,
		arg.EvseUid,
		arg.LastUpdated,
	)
	var i CommandReservation
	err := row.Scan(
		&i.ID,
		&i.Status,
		&i.TokenID,
		&i.ExpiryDate,
		&i.ReservationID,
		&i.LocationID,
		&i.EvseUid,
		&i.LastUpdated,
	)
	return i, err
}

const getCommandReservation = `-- name: GetCommandReservation :one
SELECT id, status, token_id, expiry_date, reservation_id, location_id, evse_uid, last_updated FROM command_reservations
  WHERE id = $1
`

func (q *Queries) GetCommandReservation(ctx context.Context, id int64) (CommandReservation, error) {
	row := q.db.QueryRowContext(ctx, getCommandReservation, id)
	var i CommandReservation
	err := row.Scan(
		&i.ID,
		&i.Status,
		&i.TokenID,
		&i.ExpiryDate,
		&i.ReservationID,
		&i.LocationID,
		&i.EvseUid,
		&i.LastUpdated,
	)
	return i, err
}

const updateCommandReservation = `-- name: UpdateCommandReservation :one
UPDATE command_reservations SET (
    status,
    expiry_date,
    evse_uid,
    last_updated
  ) = ($2, $3, $4, $5)
  WHERE id = $1
  RETURNING id, status, token_id, expiry_date, reservation_id, location_id, evse_uid, last_updated
`

type UpdateCommandReservationParams struct {
	ID          int64               `db:"id" json:"id"`
	Status      CommandResponseType `db:"status" json:"status"`
	ExpiryDate  time.Time           `db:"expiry_date" json:"expiryDate"`
	EvseUid     sql.NullString      `db:"evse_uid" json:"evseUid"`
	LastUpdated time.Time           `db:"last_updated" json:"lastUpdated"`
}

func (q *Queries) UpdateCommandReservation(ctx context.Context, arg UpdateCommandReservationParams) (CommandReservation, error) {
	row := q.db.QueryRowContext(ctx, updateCommandReservation,
		arg.ID,
		arg.Status,
		arg.ExpiryDate,
		arg.EvseUid,
		arg.LastUpdated,
	)
	var i CommandReservation
	err := row.Scan(
		&i.ID,
		&i.Status,
		&i.TokenID,
		&i.ExpiryDate,
		&i.ReservationID,
		&i.LocationID,
		&i.EvseUid,
		&i.LastUpdated,
	)
	return i, err
}
