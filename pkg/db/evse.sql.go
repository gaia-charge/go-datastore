// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: evse.sql

package db

import (
	"context"
	"database/sql"
	"time"

	"github.com/gaia-charge/go-datastore/pkg/geom"
)

const createEvse = `-- name: CreateEvse :one
INSERT INTO evses (
    location_id, 
    uid, 
    evse_id, 
    identifier,
    status, 
    floor_level, 
    geom, 
    geo_location_id, 
    is_remote_capable,
    is_rfid_capable,
    physical_reference, 
    last_updated)
  VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
  RETURNING id, location_id, uid, evse_id, identifier, status, floor_level, geom, geo_location_id, is_remote_capable, is_rfid_capable, physical_reference, last_updated
`

type CreateEvseParams struct {
	LocationID        int64                 `db:"location_id" json:"locationId"`
	Uid               string                `db:"uid" json:"uid"`
	EvseID            sql.NullString        `db:"evse_id" json:"evseId"`
	Identifier        sql.NullString        `db:"identifier" json:"identifier"`
	Status            EvseStatus            `db:"status" json:"status"`
	FloorLevel        sql.NullString        `db:"floor_level" json:"floorLevel"`
	Geom              geom.NullGeometry4326 `db:"geom" json:"geom"`
	GeoLocationID     sql.NullInt64         `db:"geo_location_id" json:"geoLocationId"`
	IsRemoteCapable   bool                  `db:"is_remote_capable" json:"isRemoteCapable"`
	IsRfidCapable     bool                  `db:"is_rfid_capable" json:"isRfidCapable"`
	PhysicalReference sql.NullString        `db:"physical_reference" json:"physicalReference"`
	LastUpdated       time.Time             `db:"last_updated" json:"lastUpdated"`
}

func (q *Queries) CreateEvse(ctx context.Context, arg CreateEvseParams) (Evse, error) {
	row := q.db.QueryRowContext(ctx, createEvse,
		arg.LocationID,
		arg.Uid,
		arg.EvseID,
		arg.Identifier,
		arg.Status,
		arg.FloorLevel,
		arg.Geom,
		arg.GeoLocationID,
		arg.IsRemoteCapable,
		arg.IsRfidCapable,
		arg.PhysicalReference,
		arg.LastUpdated,
	)
	var i Evse
	err := row.Scan(
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
	)
	return i, err
}

const deleteEvse = `-- name: DeleteEvse :exec
DELETE FROM evses
  WHERE id = $1
`

func (q *Queries) DeleteEvse(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteEvse, id)
	return err
}

const deleteEvseByUid = `-- name: DeleteEvseByUid :exec
DELETE FROM evses
  WHERE uid = $1
`

func (q *Queries) DeleteEvseByUid(ctx context.Context, uid string) error {
	_, err := q.db.ExecContext(ctx, deleteEvseByUid, uid)
	return err
}

const getEvse = `-- name: GetEvse :one
SELECT id, location_id, uid, evse_id, identifier, status, floor_level, geom, geo_location_id, is_remote_capable, is_rfid_capable, physical_reference, last_updated FROM evses
  WHERE id = $1
`

func (q *Queries) GetEvse(ctx context.Context, id int64) (Evse, error) {
	row := q.db.QueryRowContext(ctx, getEvse, id)
	var i Evse
	err := row.Scan(
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
	)
	return i, err
}

const getEvseByEvseID = `-- name: GetEvseByEvseID :one
SELECT id, location_id, uid, evse_id, identifier, status, floor_level, geom, geo_location_id, is_remote_capable, is_rfid_capable, physical_reference, last_updated FROM evses
  WHERE evse_id = $1
`

func (q *Queries) GetEvseByEvseID(ctx context.Context, evseID sql.NullString) (Evse, error) {
	row := q.db.QueryRowContext(ctx, getEvseByEvseID, evseID)
	var i Evse
	err := row.Scan(
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
	)
	return i, err
}

const getEvseByIdentifier = `-- name: GetEvseByIdentifier :one
SELECT id, location_id, uid, evse_id, identifier, status, floor_level, geom, geo_location_id, is_remote_capable, is_rfid_capable, physical_reference, last_updated FROM evses
  WHERE identifier = $1
`

func (q *Queries) GetEvseByIdentifier(ctx context.Context, identifier sql.NullString) (Evse, error) {
	row := q.db.QueryRowContext(ctx, getEvseByIdentifier, identifier)
	var i Evse
	err := row.Scan(
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
	)
	return i, err
}

const getEvseByUid = `-- name: GetEvseByUid :one
SELECT id, location_id, uid, evse_id, identifier, status, floor_level, geom, geo_location_id, is_remote_capable, is_rfid_capable, physical_reference, last_updated FROM evses
  WHERE uid = $1
`

func (q *Queries) GetEvseByUid(ctx context.Context, uid string) (Evse, error) {
	row := q.db.QueryRowContext(ctx, getEvseByUid, uid)
	var i Evse
	err := row.Scan(
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
	)
	return i, err
}

const listActiveEvses = `-- name: ListActiveEvses :many
SELECT id, location_id, uid, evse_id, identifier, status, floor_level, geom, geo_location_id, is_remote_capable, is_rfid_capable, physical_reference, last_updated FROM evses
  WHERE location_id = $1 AND status != 'REMOVED'
  ORDER BY id
`

func (q *Queries) ListActiveEvses(ctx context.Context, locationID int64) ([]Evse, error) {
	rows, err := q.db.QueryContext(ctx, listActiveEvses, locationID)
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

const listEvses = `-- name: ListEvses :many
SELECT id, location_id, uid, evse_id, identifier, status, floor_level, geom, geo_location_id, is_remote_capable, is_rfid_capable, physical_reference, last_updated FROM evses
  WHERE location_id = $1
  ORDER BY id
`

func (q *Queries) ListEvses(ctx context.Context, locationID int64) ([]Evse, error) {
	rows, err := q.db.QueryContext(ctx, listEvses, locationID)
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

const listEvsesLikeEvseID = `-- name: ListEvsesLikeEvseID :many
SELECT id, location_id, uid, evse_id, identifier, status, floor_level, geom, geo_location_id, is_remote_capable, is_rfid_capable, physical_reference, last_updated FROM evses
  WHERE evse_id like $1
  ORDER BY id
`

func (q *Queries) ListEvsesLikeEvseID(ctx context.Context, evseID sql.NullString) ([]Evse, error) {
	rows, err := q.db.QueryContext(ctx, listEvsesLikeEvseID, evseID)
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

const updateEvse = `-- name: UpdateEvse :one
UPDATE evses SET (
    evse_id, 
    identifier,
    status, 
    floor_level, 
    geom, 
    geo_location_id, 
    is_remote_capable,
    is_rfid_capable,
    physical_reference, 
    last_updated
  ) = ($2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
  WHERE id = $1
  RETURNING id, location_id, uid, evse_id, identifier, status, floor_level, geom, geo_location_id, is_remote_capable, is_rfid_capable, physical_reference, last_updated
`

type UpdateEvseParams struct {
	ID                int64                 `db:"id" json:"id"`
	EvseID            sql.NullString        `db:"evse_id" json:"evseId"`
	Identifier        sql.NullString        `db:"identifier" json:"identifier"`
	Status            EvseStatus            `db:"status" json:"status"`
	FloorLevel        sql.NullString        `db:"floor_level" json:"floorLevel"`
	Geom              geom.NullGeometry4326 `db:"geom" json:"geom"`
	GeoLocationID     sql.NullInt64         `db:"geo_location_id" json:"geoLocationId"`
	IsRemoteCapable   bool                  `db:"is_remote_capable" json:"isRemoteCapable"`
	IsRfidCapable     bool                  `db:"is_rfid_capable" json:"isRfidCapable"`
	PhysicalReference sql.NullString        `db:"physical_reference" json:"physicalReference"`
	LastUpdated       time.Time             `db:"last_updated" json:"lastUpdated"`
}

func (q *Queries) UpdateEvse(ctx context.Context, arg UpdateEvseParams) (Evse, error) {
	row := q.db.QueryRowContext(ctx, updateEvse,
		arg.ID,
		arg.EvseID,
		arg.Identifier,
		arg.Status,
		arg.FloorLevel,
		arg.Geom,
		arg.GeoLocationID,
		arg.IsRemoteCapable,
		arg.IsRfidCapable,
		arg.PhysicalReference,
		arg.LastUpdated,
	)
	var i Evse
	err := row.Scan(
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
	)
	return i, err
}

const updateEvseByUid = `-- name: UpdateEvseByUid :one
UPDATE evses SET (
    evse_id, 
    identifier,
    status, 
    floor_level, 
    geom, 
    geo_location_id, 
    is_remote_capable,
    is_rfid_capable,
    physical_reference, 
    last_updated
  ) = ($2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
  WHERE uid = $1
  RETURNING id, location_id, uid, evse_id, identifier, status, floor_level, geom, geo_location_id, is_remote_capable, is_rfid_capable, physical_reference, last_updated
`

type UpdateEvseByUidParams struct {
	Uid               string                `db:"uid" json:"uid"`
	EvseID            sql.NullString        `db:"evse_id" json:"evseId"`
	Identifier        sql.NullString        `db:"identifier" json:"identifier"`
	Status            EvseStatus            `db:"status" json:"status"`
	FloorLevel        sql.NullString        `db:"floor_level" json:"floorLevel"`
	Geom              geom.NullGeometry4326 `db:"geom" json:"geom"`
	GeoLocationID     sql.NullInt64         `db:"geo_location_id" json:"geoLocationId"`
	IsRemoteCapable   bool                  `db:"is_remote_capable" json:"isRemoteCapable"`
	IsRfidCapable     bool                  `db:"is_rfid_capable" json:"isRfidCapable"`
	PhysicalReference sql.NullString        `db:"physical_reference" json:"physicalReference"`
	LastUpdated       time.Time             `db:"last_updated" json:"lastUpdated"`
}

func (q *Queries) UpdateEvseByUid(ctx context.Context, arg UpdateEvseByUidParams) (Evse, error) {
	row := q.db.QueryRowContext(ctx, updateEvseByUid,
		arg.Uid,
		arg.EvseID,
		arg.Identifier,
		arg.Status,
		arg.FloorLevel,
		arg.Geom,
		arg.GeoLocationID,
		arg.IsRemoteCapable,
		arg.IsRfidCapable,
		arg.PhysicalReference,
		arg.LastUpdated,
	)
	var i Evse
	err := row.Scan(
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
	)
	return i, err
}

const updateEvseLastUpdated = `-- name: UpdateEvseLastUpdated :exec
UPDATE evses SET last_updated = $2
  WHERE id = $1
`

type UpdateEvseLastUpdatedParams struct {
	ID          int64     `db:"id" json:"id"`
	LastUpdated time.Time `db:"last_updated" json:"lastUpdated"`
}

func (q *Queries) UpdateEvseLastUpdated(ctx context.Context, arg UpdateEvseLastUpdatedParams) error {
	_, err := q.db.ExecContext(ctx, updateEvseLastUpdated, arg.ID, arg.LastUpdated)
	return err
}
