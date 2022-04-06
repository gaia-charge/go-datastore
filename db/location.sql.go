// Code generated by sqlc. DO NOT EDIT.
// source: location.sql

package db

import (
	"context"
	"database/sql"
	"time"

	"github.com/twpayne/go-geom"
)

const createLocation = `-- name: CreateLocation :one
INSERT INTO locations (
    uid, 
    country_code,
    party_id,
    type, 
    name, 
    address, 
    city, 
    postal_code, 
    country, 
    geom, 
    geo_location_id, 
    operator_id, 
    suboperator_id, 
    owner_id, 
    time_zone, 
    opening_time_id,
    charging_when_closed, 
    energy_mix_id, 
    last_updated
  ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19)
  RETURNING id, uid, country_code, party_id, type, name, address, city, postal_code, country, geom, geo_location_id, operator_id, suboperator_id, owner_id, time_zone, opening_time_id, charging_when_closed, energy_mix_id, last_updated
`

type CreateLocationParams struct {
	Uid                string         `db:"uid" json:"uid"`
	CountryCode        sql.NullString `db:"country_code" json:"countryCode"`
	PartyID            sql.NullString `db:"party_id" json:"partyID"`
	Type               LocationType   `db:"type" json:"type"`
	Name               sql.NullString `db:"name" json:"name"`
	Address            string         `db:"address" json:"address"`
	City               string         `db:"city" json:"city"`
	PostalCode         string         `db:"postal_code" json:"postalCode"`
	Country            string         `db:"country" json:"country"`
	Geom               geom.Point     `db:"geom" json:"geom"`
	GeoLocationID      int64          `db:"geo_location_id" json:"geoLocationID"`
	OperatorID         sql.NullInt64  `db:"operator_id" json:"operatorID"`
	SuboperatorID      sql.NullInt64  `db:"suboperator_id" json:"suboperatorID"`
	OwnerID            sql.NullInt64  `db:"owner_id" json:"ownerID"`
	TimeZone           sql.NullString `db:"time_zone" json:"timeZone"`
	OpeningTimeID      sql.NullInt64  `db:"opening_time_id" json:"openingTimeID"`
	ChargingWhenClosed bool           `db:"charging_when_closed" json:"chargingWhenClosed"`
	EnergyMixID        sql.NullInt64  `db:"energy_mix_id" json:"energyMixID"`
	LastUpdated        time.Time      `db:"last_updated" json:"lastUpdated"`
}

func (q *Queries) CreateLocation(ctx context.Context, arg CreateLocationParams) (Location, error) {
	row := q.db.QueryRowContext(ctx, createLocation,
		arg.Uid,
		arg.CountryCode,
		arg.PartyID,
		arg.Type,
		arg.Name,
		arg.Address,
		arg.City,
		arg.PostalCode,
		arg.Country,
		arg.Geom,
		arg.GeoLocationID,
		arg.OperatorID,
		arg.SuboperatorID,
		arg.OwnerID,
		arg.TimeZone,
		arg.OpeningTimeID,
		arg.ChargingWhenClosed,
		arg.EnergyMixID,
		arg.LastUpdated,
	)
	var i Location
	err := row.Scan(
		&i.ID,
		&i.Uid,
		&i.CountryCode,
		&i.PartyID,
		&i.Type,
		&i.Name,
		&i.Address,
		&i.City,
		&i.PostalCode,
		&i.Country,
		&i.Geom,
		&i.GeoLocationID,
		&i.OperatorID,
		&i.SuboperatorID,
		&i.OwnerID,
		&i.TimeZone,
		&i.OpeningTimeID,
		&i.ChargingWhenClosed,
		&i.EnergyMixID,
		&i.LastUpdated,
	)
	return i, err
}

const deleteLocation = `-- name: DeleteLocation :exec
DELETE FROM locations
  WHERE id = $1
`

func (q *Queries) DeleteLocation(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteLocation, id)
	return err
}

const deleteLocationByUid = `-- name: DeleteLocationByUid :one
DELETE FROM locations
  WHERE uid = $1
  RETURNING id, uid, country_code, party_id, type, name, address, city, postal_code, country, geom, geo_location_id, operator_id, suboperator_id, owner_id, time_zone, opening_time_id, charging_when_closed, energy_mix_id, last_updated
`

func (q *Queries) DeleteLocationByUid(ctx context.Context, uid string) (Location, error) {
	row := q.db.QueryRowContext(ctx, deleteLocationByUid, uid)
	var i Location
	err := row.Scan(
		&i.ID,
		&i.Uid,
		&i.CountryCode,
		&i.PartyID,
		&i.Type,
		&i.Name,
		&i.Address,
		&i.City,
		&i.PostalCode,
		&i.Country,
		&i.Geom,
		&i.GeoLocationID,
		&i.OperatorID,
		&i.SuboperatorID,
		&i.OwnerID,
		&i.TimeZone,
		&i.OpeningTimeID,
		&i.ChargingWhenClosed,
		&i.EnergyMixID,
		&i.LastUpdated,
	)
	return i, err
}

const getLocation = `-- name: GetLocation :one
SELECT id, uid, country_code, party_id, type, name, address, city, postal_code, country, geom, geo_location_id, operator_id, suboperator_id, owner_id, time_zone, opening_time_id, charging_when_closed, energy_mix_id, last_updated FROM locations
  WHERE id = $1
`

func (q *Queries) GetLocation(ctx context.Context, id int64) (Location, error) {
	row := q.db.QueryRowContext(ctx, getLocation, id)
	var i Location
	err := row.Scan(
		&i.ID,
		&i.Uid,
		&i.CountryCode,
		&i.PartyID,
		&i.Type,
		&i.Name,
		&i.Address,
		&i.City,
		&i.PostalCode,
		&i.Country,
		&i.Geom,
		&i.GeoLocationID,
		&i.OperatorID,
		&i.SuboperatorID,
		&i.OwnerID,
		&i.TimeZone,
		&i.OpeningTimeID,
		&i.ChargingWhenClosed,
		&i.EnergyMixID,
		&i.LastUpdated,
	)
	return i, err
}

const getLocationByIdentityOrderByLastUpdated = `-- name: GetLocationByIdentityOrderByLastUpdated :one
SELECT id, uid, country_code, party_id, type, name, address, city, postal_code, country, geom, geo_location_id, operator_id, suboperator_id, owner_id, time_zone, opening_time_id, charging_when_closed, energy_mix_id, last_updated FROM locations
  WHERE country_code = $1 AND party_id = $2
  ORDER BY last_updated DESC
  LIMIT 1
`

type GetLocationByIdentityOrderByLastUpdatedParams struct {
	CountryCode sql.NullString `db:"country_code" json:"countryCode"`
	PartyID     sql.NullString `db:"party_id" json:"partyID"`
}

func (q *Queries) GetLocationByIdentityOrderByLastUpdated(ctx context.Context, arg GetLocationByIdentityOrderByLastUpdatedParams) (Location, error) {
	row := q.db.QueryRowContext(ctx, getLocationByIdentityOrderByLastUpdated, arg.CountryCode, arg.PartyID)
	var i Location
	err := row.Scan(
		&i.ID,
		&i.Uid,
		&i.CountryCode,
		&i.PartyID,
		&i.Type,
		&i.Name,
		&i.Address,
		&i.City,
		&i.PostalCode,
		&i.Country,
		&i.Geom,
		&i.GeoLocationID,
		&i.OperatorID,
		&i.SuboperatorID,
		&i.OwnerID,
		&i.TimeZone,
		&i.OpeningTimeID,
		&i.ChargingWhenClosed,
		&i.EnergyMixID,
		&i.LastUpdated,
	)
	return i, err
}

const getLocationByUid = `-- name: GetLocationByUid :one
SELECT id, uid, country_code, party_id, type, name, address, city, postal_code, country, geom, geo_location_id, operator_id, suboperator_id, owner_id, time_zone, opening_time_id, charging_when_closed, energy_mix_id, last_updated FROM locations
  WHERE uid = $1
`

func (q *Queries) GetLocationByUid(ctx context.Context, uid string) (Location, error) {
	row := q.db.QueryRowContext(ctx, getLocationByUid, uid)
	var i Location
	err := row.Scan(
		&i.ID,
		&i.Uid,
		&i.CountryCode,
		&i.PartyID,
		&i.Type,
		&i.Name,
		&i.Address,
		&i.City,
		&i.PostalCode,
		&i.Country,
		&i.Geom,
		&i.GeoLocationID,
		&i.OperatorID,
		&i.SuboperatorID,
		&i.OwnerID,
		&i.TimeZone,
		&i.OpeningTimeID,
		&i.ChargingWhenClosed,
		&i.EnergyMixID,
		&i.LastUpdated,
	)
	return i, err
}

const listLocations = `-- name: ListLocations :many
SELECT id, uid, country_code, party_id, type, name, address, city, postal_code, country, geom, geo_location_id, operator_id, suboperator_id, owner_id, time_zone, opening_time_id, charging_when_closed, energy_mix_id, last_updated FROM locations
  ORDER BY name
`

func (q *Queries) ListLocations(ctx context.Context) ([]Location, error) {
	rows, err := q.db.QueryContext(ctx, listLocations)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Location
	for rows.Next() {
		var i Location
		if err := rows.Scan(
			&i.ID,
			&i.Uid,
			&i.CountryCode,
			&i.PartyID,
			&i.Type,
			&i.Name,
			&i.Address,
			&i.City,
			&i.PostalCode,
			&i.Country,
			&i.Geom,
			&i.GeoLocationID,
			&i.OperatorID,
			&i.SuboperatorID,
			&i.OwnerID,
			&i.TimeZone,
			&i.OpeningTimeID,
			&i.ChargingWhenClosed,
			&i.EnergyMixID,
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

const updateLocation = `-- name: UpdateLocation :one
UPDATE locations SET (
    country_code,
    party_id,
    type, 
    name, 
    address, 
    city, 
    postal_code, 
    country, 
    geom, 
    geo_location_id, 
    operator_id, 
    suboperator_id, 
    owner_id, 
    time_zone, 
    opening_time_id,
    charging_when_closed,
    energy_mix_id, 
    last_updated
  ) = ($2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19)
  WHERE id = $1
  RETURNING id, uid, country_code, party_id, type, name, address, city, postal_code, country, geom, geo_location_id, operator_id, suboperator_id, owner_id, time_zone, opening_time_id, charging_when_closed, energy_mix_id, last_updated
`

type UpdateLocationParams struct {
	ID                 int64          `db:"id" json:"id"`
	CountryCode        sql.NullString `db:"country_code" json:"countryCode"`
	PartyID            sql.NullString `db:"party_id" json:"partyID"`
	Type               LocationType   `db:"type" json:"type"`
	Name               sql.NullString `db:"name" json:"name"`
	Address            string         `db:"address" json:"address"`
	City               string         `db:"city" json:"city"`
	PostalCode         string         `db:"postal_code" json:"postalCode"`
	Country            string         `db:"country" json:"country"`
	Geom               geom.Point     `db:"geom" json:"geom"`
	GeoLocationID      int64          `db:"geo_location_id" json:"geoLocationID"`
	OperatorID         sql.NullInt64  `db:"operator_id" json:"operatorID"`
	SuboperatorID      sql.NullInt64  `db:"suboperator_id" json:"suboperatorID"`
	OwnerID            sql.NullInt64  `db:"owner_id" json:"ownerID"`
	TimeZone           sql.NullString `db:"time_zone" json:"timeZone"`
	OpeningTimeID      sql.NullInt64  `db:"opening_time_id" json:"openingTimeID"`
	ChargingWhenClosed bool           `db:"charging_when_closed" json:"chargingWhenClosed"`
	EnergyMixID        sql.NullInt64  `db:"energy_mix_id" json:"energyMixID"`
	LastUpdated        time.Time      `db:"last_updated" json:"lastUpdated"`
}

func (q *Queries) UpdateLocation(ctx context.Context, arg UpdateLocationParams) (Location, error) {
	row := q.db.QueryRowContext(ctx, updateLocation,
		arg.ID,
		arg.CountryCode,
		arg.PartyID,
		arg.Type,
		arg.Name,
		arg.Address,
		arg.City,
		arg.PostalCode,
		arg.Country,
		arg.Geom,
		arg.GeoLocationID,
		arg.OperatorID,
		arg.SuboperatorID,
		arg.OwnerID,
		arg.TimeZone,
		arg.OpeningTimeID,
		arg.ChargingWhenClosed,
		arg.EnergyMixID,
		arg.LastUpdated,
	)
	var i Location
	err := row.Scan(
		&i.ID,
		&i.Uid,
		&i.CountryCode,
		&i.PartyID,
		&i.Type,
		&i.Name,
		&i.Address,
		&i.City,
		&i.PostalCode,
		&i.Country,
		&i.Geom,
		&i.GeoLocationID,
		&i.OperatorID,
		&i.SuboperatorID,
		&i.OwnerID,
		&i.TimeZone,
		&i.OpeningTimeID,
		&i.ChargingWhenClosed,
		&i.EnergyMixID,
		&i.LastUpdated,
	)
	return i, err
}

const updateLocationByUid = `-- name: UpdateLocationByUid :one
UPDATE locations SET (
    country_code,
    party_id,
    type, 
    name, 
    address, 
    city, 
    postal_code, 
    country, 
    geom, 
    geo_location_id, 
    operator_id, 
    suboperator_id, 
    owner_id, 
    time_zone, 
    opening_time_id,
    charging_when_closed,
    energy_mix_id, 
    last_updated
  ) = ($2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19)
  WHERE uid = $1
  RETURNING id, uid, country_code, party_id, type, name, address, city, postal_code, country, geom, geo_location_id, operator_id, suboperator_id, owner_id, time_zone, opening_time_id, charging_when_closed, energy_mix_id, last_updated
`

type UpdateLocationByUidParams struct {
	Uid                string         `db:"uid" json:"uid"`
	CountryCode        sql.NullString `db:"country_code" json:"countryCode"`
	PartyID            sql.NullString `db:"party_id" json:"partyID"`
	Type               LocationType   `db:"type" json:"type"`
	Name               sql.NullString `db:"name" json:"name"`
	Address            string         `db:"address" json:"address"`
	City               string         `db:"city" json:"city"`
	PostalCode         string         `db:"postal_code" json:"postalCode"`
	Country            string         `db:"country" json:"country"`
	Geom               geom.Point     `db:"geom" json:"geom"`
	GeoLocationID      int64          `db:"geo_location_id" json:"geoLocationID"`
	OperatorID         sql.NullInt64  `db:"operator_id" json:"operatorID"`
	SuboperatorID      sql.NullInt64  `db:"suboperator_id" json:"suboperatorID"`
	OwnerID            sql.NullInt64  `db:"owner_id" json:"ownerID"`
	TimeZone           sql.NullString `db:"time_zone" json:"timeZone"`
	OpeningTimeID      sql.NullInt64  `db:"opening_time_id" json:"openingTimeID"`
	ChargingWhenClosed bool           `db:"charging_when_closed" json:"chargingWhenClosed"`
	EnergyMixID        sql.NullInt64  `db:"energy_mix_id" json:"energyMixID"`
	LastUpdated        time.Time      `db:"last_updated" json:"lastUpdated"`
}

func (q *Queries) UpdateLocationByUid(ctx context.Context, arg UpdateLocationByUidParams) (Location, error) {
	row := q.db.QueryRowContext(ctx, updateLocationByUid,
		arg.Uid,
		arg.CountryCode,
		arg.PartyID,
		arg.Type,
		arg.Name,
		arg.Address,
		arg.City,
		arg.PostalCode,
		arg.Country,
		arg.Geom,
		arg.GeoLocationID,
		arg.OperatorID,
		arg.SuboperatorID,
		arg.OwnerID,
		arg.TimeZone,
		arg.OpeningTimeID,
		arg.ChargingWhenClosed,
		arg.EnergyMixID,
		arg.LastUpdated,
	)
	var i Location
	err := row.Scan(
		&i.ID,
		&i.Uid,
		&i.CountryCode,
		&i.PartyID,
		&i.Type,
		&i.Name,
		&i.Address,
		&i.City,
		&i.PostalCode,
		&i.Country,
		&i.Geom,
		&i.GeoLocationID,
		&i.OperatorID,
		&i.SuboperatorID,
		&i.OwnerID,
		&i.TimeZone,
		&i.OpeningTimeID,
		&i.ChargingWhenClosed,
		&i.EnergyMixID,
		&i.LastUpdated,
	)
	return i, err
}

const updateLocationLastUpdated = `-- name: UpdateLocationLastUpdated :exec
UPDATE locations SET last_updated = $2
  WHERE id = $1
`

type UpdateLocationLastUpdatedParams struct {
	ID          int64     `db:"id" json:"id"`
	LastUpdated time.Time `db:"last_updated" json:"lastUpdated"`
}

func (q *Queries) UpdateLocationLastUpdated(ctx context.Context, arg UpdateLocationLastUpdatedParams) error {
	_, err := q.db.ExecContext(ctx, updateLocationLastUpdated, arg.ID, arg.LastUpdated)
	return err
}
