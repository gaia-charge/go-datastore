// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: location.sql

package db

import (
	"context"
	"database/sql"
	"time"

	"github.com/gaia-charge/go-datastore/pkg/geom"
)

const createLocation = `-- name: CreateLocation :one
INSERT INTO locations (
    uid, 
    credential_id,
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
    available_evses,
    total_evses,
    is_intermediate_cdr_capable,
    is_remote_capable,
    is_rfid_capable,
    operator_id, 
    suboperator_id, 
    owner_id, 
    time_zone, 
    opening_time_id,
    charging_when_closed, 
    energy_mix_id, 
    is_published,
    last_updated
  ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26)
  RETURNING id, uid, credential_id, country_code, party_id, type, name, address, city, postal_code, country, geom, geo_location_id, available_evses, total_evses, is_remote_capable, is_rfid_capable, operator_id, suboperator_id, owner_id, time_zone, opening_time_id, charging_when_closed, energy_mix_id, last_updated, is_published, added_date, is_intermediate_cdr_capable, is_removed
`

type CreateLocationParams struct {
	Uid                      string            `db:"uid" json:"uid"`
	CredentialID             int64             `db:"credential_id" json:"credentialId"`
	CountryCode              sql.NullString    `db:"country_code" json:"countryCode"`
	PartyID                  sql.NullString    `db:"party_id" json:"partyId"`
	Type                     LocationType      `db:"type" json:"type"`
	Name                     sql.NullString    `db:"name" json:"name"`
	Address                  string            `db:"address" json:"address"`
	City                     string            `db:"city" json:"city"`
	PostalCode               string            `db:"postal_code" json:"postalCode"`
	Country                  string            `db:"country" json:"country"`
	Geom                     geom.Geometry4326 `db:"geom" json:"geom"`
	GeoLocationID            int64             `db:"geo_location_id" json:"geoLocationId"`
	AvailableEvses           int32             `db:"available_evses" json:"availableEvses"`
	TotalEvses               int32             `db:"total_evses" json:"totalEvses"`
	IsIntermediateCdrCapable bool              `db:"is_intermediate_cdr_capable" json:"isIntermediateCdrCapable"`
	IsRemoteCapable          bool              `db:"is_remote_capable" json:"isRemoteCapable"`
	IsRfidCapable            bool              `db:"is_rfid_capable" json:"isRfidCapable"`
	OperatorID               sql.NullInt64     `db:"operator_id" json:"operatorId"`
	SuboperatorID            sql.NullInt64     `db:"suboperator_id" json:"suboperatorId"`
	OwnerID                  sql.NullInt64     `db:"owner_id" json:"ownerId"`
	TimeZone                 sql.NullString    `db:"time_zone" json:"timeZone"`
	OpeningTimeID            sql.NullInt64     `db:"opening_time_id" json:"openingTimeId"`
	ChargingWhenClosed       bool              `db:"charging_when_closed" json:"chargingWhenClosed"`
	EnergyMixID              sql.NullInt64     `db:"energy_mix_id" json:"energyMixId"`
	IsPublished              bool              `db:"is_published" json:"isPublished"`
	LastUpdated              time.Time         `db:"last_updated" json:"lastUpdated"`
}

func (q *Queries) CreateLocation(ctx context.Context, arg CreateLocationParams) (Location, error) {
	row := q.db.QueryRowContext(ctx, createLocation,
		arg.Uid,
		arg.CredentialID,
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
		arg.AvailableEvses,
		arg.TotalEvses,
		arg.IsIntermediateCdrCapable,
		arg.IsRemoteCapable,
		arg.IsRfidCapable,
		arg.OperatorID,
		arg.SuboperatorID,
		arg.OwnerID,
		arg.TimeZone,
		arg.OpeningTimeID,
		arg.ChargingWhenClosed,
		arg.EnergyMixID,
		arg.IsPublished,
		arg.LastUpdated,
	)
	var i Location
	err := row.Scan(
		&i.ID,
		&i.Uid,
		&i.CredentialID,
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
		&i.AvailableEvses,
		&i.TotalEvses,
		&i.IsRemoteCapable,
		&i.IsRfidCapable,
		&i.OperatorID,
		&i.SuboperatorID,
		&i.OwnerID,
		&i.TimeZone,
		&i.OpeningTimeID,
		&i.ChargingWhenClosed,
		&i.EnergyMixID,
		&i.LastUpdated,
		&i.IsPublished,
		&i.AddedDate,
		&i.IsIntermediateCdrCapable,
		&i.IsRemoved,
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
  RETURNING id, uid, credential_id, country_code, party_id, type, name, address, city, postal_code, country, geom, geo_location_id, available_evses, total_evses, is_remote_capable, is_rfid_capable, operator_id, suboperator_id, owner_id, time_zone, opening_time_id, charging_when_closed, energy_mix_id, last_updated, is_published, added_date, is_intermediate_cdr_capable, is_removed
`

func (q *Queries) DeleteLocationByUid(ctx context.Context, uid string) (Location, error) {
	row := q.db.QueryRowContext(ctx, deleteLocationByUid, uid)
	var i Location
	err := row.Scan(
		&i.ID,
		&i.Uid,
		&i.CredentialID,
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
		&i.AvailableEvses,
		&i.TotalEvses,
		&i.IsRemoteCapable,
		&i.IsRfidCapable,
		&i.OperatorID,
		&i.SuboperatorID,
		&i.OwnerID,
		&i.TimeZone,
		&i.OpeningTimeID,
		&i.ChargingWhenClosed,
		&i.EnergyMixID,
		&i.LastUpdated,
		&i.IsPublished,
		&i.AddedDate,
		&i.IsIntermediateCdrCapable,
		&i.IsRemoved,
	)
	return i, err
}

const getLocation = `-- name: GetLocation :one
SELECT id, uid, credential_id, country_code, party_id, type, name, address, city, postal_code, country, geom, geo_location_id, available_evses, total_evses, is_remote_capable, is_rfid_capable, operator_id, suboperator_id, owner_id, time_zone, opening_time_id, charging_when_closed, energy_mix_id, last_updated, is_published, added_date, is_intermediate_cdr_capable, is_removed FROM locations
  WHERE id = $1
`

func (q *Queries) GetLocation(ctx context.Context, id int64) (Location, error) {
	row := q.db.QueryRowContext(ctx, getLocation, id)
	var i Location
	err := row.Scan(
		&i.ID,
		&i.Uid,
		&i.CredentialID,
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
		&i.AvailableEvses,
		&i.TotalEvses,
		&i.IsRemoteCapable,
		&i.IsRfidCapable,
		&i.OperatorID,
		&i.SuboperatorID,
		&i.OwnerID,
		&i.TimeZone,
		&i.OpeningTimeID,
		&i.ChargingWhenClosed,
		&i.EnergyMixID,
		&i.LastUpdated,
		&i.IsPublished,
		&i.AddedDate,
		&i.IsIntermediateCdrCapable,
		&i.IsRemoved,
	)
	return i, err
}

const getLocationByLastUpdated = `-- name: GetLocationByLastUpdated :one
SELECT id, uid, credential_id, country_code, party_id, type, name, address, city, postal_code, country, geom, geo_location_id, available_evses, total_evses, is_remote_capable, is_rfid_capable, operator_id, suboperator_id, owner_id, time_zone, opening_time_id, charging_when_closed, energy_mix_id, last_updated, is_published, added_date, is_intermediate_cdr_capable, is_removed FROM locations
  WHERE ($1::BIGINT = -1 OR $1::BIGINT = credential_id) AND
    ($2::TEXT = '' OR $2::TEXT = country_code) AND
    ($3::TEXT = '' OR $3::TEXT = party_id)
  ORDER BY last_updated DESC
  LIMIT 1
`

type GetLocationByLastUpdatedParams struct {
	CredentialID int64  `db:"credential_id" json:"credentialId"`
	CountryCode  string `db:"country_code" json:"countryCode"`
	PartyID      string `db:"party_id" json:"partyId"`
}

func (q *Queries) GetLocationByLastUpdated(ctx context.Context, arg GetLocationByLastUpdatedParams) (Location, error) {
	row := q.db.QueryRowContext(ctx, getLocationByLastUpdated, arg.CredentialID, arg.CountryCode, arg.PartyID)
	var i Location
	err := row.Scan(
		&i.ID,
		&i.Uid,
		&i.CredentialID,
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
		&i.AvailableEvses,
		&i.TotalEvses,
		&i.IsRemoteCapable,
		&i.IsRfidCapable,
		&i.OperatorID,
		&i.SuboperatorID,
		&i.OwnerID,
		&i.TimeZone,
		&i.OpeningTimeID,
		&i.ChargingWhenClosed,
		&i.EnergyMixID,
		&i.LastUpdated,
		&i.IsPublished,
		&i.AddedDate,
		&i.IsIntermediateCdrCapable,
		&i.IsRemoved,
	)
	return i, err
}

const getLocationByUid = `-- name: GetLocationByUid :one
SELECT id, uid, credential_id, country_code, party_id, type, name, address, city, postal_code, country, geom, geo_location_id, available_evses, total_evses, is_remote_capable, is_rfid_capable, operator_id, suboperator_id, owner_id, time_zone, opening_time_id, charging_when_closed, energy_mix_id, last_updated, is_published, added_date, is_intermediate_cdr_capable, is_removed FROM locations
  WHERE uid = $1
`

func (q *Queries) GetLocationByUid(ctx context.Context, uid string) (Location, error) {
	row := q.db.QueryRowContext(ctx, getLocationByUid, uid)
	var i Location
	err := row.Scan(
		&i.ID,
		&i.Uid,
		&i.CredentialID,
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
		&i.AvailableEvses,
		&i.TotalEvses,
		&i.IsRemoteCapable,
		&i.IsRfidCapable,
		&i.OperatorID,
		&i.SuboperatorID,
		&i.OwnerID,
		&i.TimeZone,
		&i.OpeningTimeID,
		&i.ChargingWhenClosed,
		&i.EnergyMixID,
		&i.LastUpdated,
		&i.IsPublished,
		&i.AddedDate,
		&i.IsIntermediateCdrCapable,
		&i.IsRemoved,
	)
	return i, err
}

const listLocations = `-- name: ListLocations :many
SELECT id, uid, credential_id, country_code, party_id, type, name, address, city, postal_code, country, geom, geo_location_id, available_evses, total_evses, is_remote_capable, is_rfid_capable, operator_id, suboperator_id, owner_id, time_zone, opening_time_id, charging_when_closed, energy_mix_id, last_updated, is_published, added_date, is_intermediate_cdr_capable, is_removed FROM locations
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
			&i.CredentialID,
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
			&i.AvailableEvses,
			&i.TotalEvses,
			&i.IsRemoteCapable,
			&i.IsRfidCapable,
			&i.OperatorID,
			&i.SuboperatorID,
			&i.OwnerID,
			&i.TimeZone,
			&i.OpeningTimeID,
			&i.ChargingWhenClosed,
			&i.EnergyMixID,
			&i.LastUpdated,
			&i.IsPublished,
			&i.AddedDate,
			&i.IsIntermediateCdrCapable,
			&i.IsRemoved,
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

const listLocationsByCountry = `-- name: ListLocationsByCountry :many
SELECT id, uid, credential_id, country_code, party_id, type, name, address, city, postal_code, country, geom, geo_location_id, available_evses, total_evses, is_remote_capable, is_rfid_capable, operator_id, suboperator_id, owner_id, time_zone, opening_time_id, charging_when_closed, energy_mix_id, last_updated, is_published, added_date, is_intermediate_cdr_capable, is_removed FROM locations
  WHERE is_published = true AND is_removed = false AND total_evses > 0 AND country = $1
`

func (q *Queries) ListLocationsByCountry(ctx context.Context, country string) ([]Location, error) {
	rows, err := q.db.QueryContext(ctx, listLocationsByCountry, country)
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
			&i.CredentialID,
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
			&i.AvailableEvses,
			&i.TotalEvses,
			&i.IsRemoteCapable,
			&i.IsRfidCapable,
			&i.OperatorID,
			&i.SuboperatorID,
			&i.OwnerID,
			&i.TimeZone,
			&i.OpeningTimeID,
			&i.ChargingWhenClosed,
			&i.EnergyMixID,
			&i.LastUpdated,
			&i.IsPublished,
			&i.AddedDate,
			&i.IsIntermediateCdrCapable,
			&i.IsRemoved,
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

const listLocationsByGeom = `-- name: ListLocationsByGeom :many
SELECT l.id, l.uid, l.credential_id, l.country_code, l.party_id, l.type, l.name, l.address, l.city, l.postal_code, l.country, l.geom, l.geo_location_id, l.available_evses, l.total_evses, l.is_remote_capable, l.is_rfid_capable, l.operator_id, l.suboperator_id, l.owner_id, l.time_zone, l.opening_time_id, l.charging_when_closed, l.energy_mix_id, l.last_updated, l.is_published, l.added_date, l.is_intermediate_cdr_capable, l.is_removed FROM locations l
  INNER JOIN credentials c ON c.id = l.credential_id
  WHERE c.is_available = true AND l.is_published = true AND l.is_removed = false AND l.total_evses > 0 AND 
    ($1::BOOLEAN = true OR l.is_intermediate_cdr_capable = true) AND 
    (
      ($2::BOOLEAN = true AND l.is_remote_capable = true) OR 
      ($3::BOOLEAN = true AND l.is_rfid_capable = true)
    ) AND
    ST_Intersects(l.geom, ST_MakeEnvelope($4::FLOAT, $5::FLOAT, $6::FLOAT, $7::FLOAT, 4326)) AND
    ($8::INT = 0 OR l.last_updated > NOW() - '1 second'::INTERVAL * $8::INT)
  LIMIT 500
`

type ListLocationsByGeomParams struct {
	IsExperimental  bool    `db:"is_experimental" json:"isExperimental"`
	IsRemoteCapable bool    `db:"is_remote_capable" json:"isRemoteCapable"`
	IsRfidCapable   bool    `db:"is_rfid_capable" json:"isRfidCapable"`
	XMin            float64 `db:"x_min" json:"xMin"`
	YMin            float64 `db:"y_min" json:"yMin"`
	XMax            float64 `db:"x_max" json:"xMax"`
	YMax            float64 `db:"y_max" json:"yMax"`
	Interval        int32   `db:"interval" json:"interval"`
}

func (q *Queries) ListLocationsByGeom(ctx context.Context, arg ListLocationsByGeomParams) ([]Location, error) {
	rows, err := q.db.QueryContext(ctx, listLocationsByGeom,
		arg.IsExperimental,
		arg.IsRemoteCapable,
		arg.IsRfidCapable,
		arg.XMin,
		arg.YMin,
		arg.XMax,
		arg.YMax,
		arg.Interval,
	)
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
			&i.CredentialID,
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
			&i.AvailableEvses,
			&i.TotalEvses,
			&i.IsRemoteCapable,
			&i.IsRfidCapable,
			&i.OperatorID,
			&i.SuboperatorID,
			&i.OwnerID,
			&i.TimeZone,
			&i.OpeningTimeID,
			&i.ChargingWhenClosed,
			&i.EnergyMixID,
			&i.LastUpdated,
			&i.IsPublished,
			&i.AddedDate,
			&i.IsIntermediateCdrCapable,
			&i.IsRemoved,
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
    available_evses,
    total_evses,
    is_intermediate_cdr_capable,
    is_remote_capable,
    is_rfid_capable,
    is_removed,
    operator_id, 
    suboperator_id, 
    owner_id, 
    time_zone, 
    opening_time_id,
    charging_when_closed,
    energy_mix_id, 
    last_updated
  ) = ($2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25)
  WHERE id = $1
  RETURNING id, uid, credential_id, country_code, party_id, type, name, address, city, postal_code, country, geom, geo_location_id, available_evses, total_evses, is_remote_capable, is_rfid_capable, operator_id, suboperator_id, owner_id, time_zone, opening_time_id, charging_when_closed, energy_mix_id, last_updated, is_published, added_date, is_intermediate_cdr_capable, is_removed
`

type UpdateLocationParams struct {
	ID                       int64             `db:"id" json:"id"`
	CountryCode              sql.NullString    `db:"country_code" json:"countryCode"`
	PartyID                  sql.NullString    `db:"party_id" json:"partyId"`
	Type                     LocationType      `db:"type" json:"type"`
	Name                     sql.NullString    `db:"name" json:"name"`
	Address                  string            `db:"address" json:"address"`
	City                     string            `db:"city" json:"city"`
	PostalCode               string            `db:"postal_code" json:"postalCode"`
	Country                  string            `db:"country" json:"country"`
	Geom                     geom.Geometry4326 `db:"geom" json:"geom"`
	GeoLocationID            int64             `db:"geo_location_id" json:"geoLocationId"`
	AvailableEvses           int32             `db:"available_evses" json:"availableEvses"`
	TotalEvses               int32             `db:"total_evses" json:"totalEvses"`
	IsIntermediateCdrCapable bool              `db:"is_intermediate_cdr_capable" json:"isIntermediateCdrCapable"`
	IsRemoteCapable          bool              `db:"is_remote_capable" json:"isRemoteCapable"`
	IsRfidCapable            bool              `db:"is_rfid_capable" json:"isRfidCapable"`
	IsRemoved                bool              `db:"is_removed" json:"isRemoved"`
	OperatorID               sql.NullInt64     `db:"operator_id" json:"operatorId"`
	SuboperatorID            sql.NullInt64     `db:"suboperator_id" json:"suboperatorId"`
	OwnerID                  sql.NullInt64     `db:"owner_id" json:"ownerId"`
	TimeZone                 sql.NullString    `db:"time_zone" json:"timeZone"`
	OpeningTimeID            sql.NullInt64     `db:"opening_time_id" json:"openingTimeId"`
	ChargingWhenClosed       bool              `db:"charging_when_closed" json:"chargingWhenClosed"`
	EnergyMixID              sql.NullInt64     `db:"energy_mix_id" json:"energyMixId"`
	LastUpdated              time.Time         `db:"last_updated" json:"lastUpdated"`
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
		arg.AvailableEvses,
		arg.TotalEvses,
		arg.IsIntermediateCdrCapable,
		arg.IsRemoteCapable,
		arg.IsRfidCapable,
		arg.IsRemoved,
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
		&i.CredentialID,
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
		&i.AvailableEvses,
		&i.TotalEvses,
		&i.IsRemoteCapable,
		&i.IsRfidCapable,
		&i.OperatorID,
		&i.SuboperatorID,
		&i.OwnerID,
		&i.TimeZone,
		&i.OpeningTimeID,
		&i.ChargingWhenClosed,
		&i.EnergyMixID,
		&i.LastUpdated,
		&i.IsPublished,
		&i.AddedDate,
		&i.IsIntermediateCdrCapable,
		&i.IsRemoved,
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
    available_evses,
    total_evses,
    is_intermediate_cdr_capable,
    is_remote_capable,
    is_rfid_capable,
    is_removed,
    operator_id, 
    suboperator_id, 
    owner_id, 
    time_zone, 
    opening_time_id,
    charging_when_closed,
    energy_mix_id, 
    last_updated
  ) = ($2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25)
  WHERE uid = $1
  RETURNING id, uid, credential_id, country_code, party_id, type, name, address, city, postal_code, country, geom, geo_location_id, available_evses, total_evses, is_remote_capable, is_rfid_capable, operator_id, suboperator_id, owner_id, time_zone, opening_time_id, charging_when_closed, energy_mix_id, last_updated, is_published, added_date, is_intermediate_cdr_capable, is_removed
`

type UpdateLocationByUidParams struct {
	Uid                      string            `db:"uid" json:"uid"`
	CountryCode              sql.NullString    `db:"country_code" json:"countryCode"`
	PartyID                  sql.NullString    `db:"party_id" json:"partyId"`
	Type                     LocationType      `db:"type" json:"type"`
	Name                     sql.NullString    `db:"name" json:"name"`
	Address                  string            `db:"address" json:"address"`
	City                     string            `db:"city" json:"city"`
	PostalCode               string            `db:"postal_code" json:"postalCode"`
	Country                  string            `db:"country" json:"country"`
	Geom                     geom.Geometry4326 `db:"geom" json:"geom"`
	GeoLocationID            int64             `db:"geo_location_id" json:"geoLocationId"`
	AvailableEvses           int32             `db:"available_evses" json:"availableEvses"`
	TotalEvses               int32             `db:"total_evses" json:"totalEvses"`
	IsIntermediateCdrCapable bool              `db:"is_intermediate_cdr_capable" json:"isIntermediateCdrCapable"`
	IsRemoteCapable          bool              `db:"is_remote_capable" json:"isRemoteCapable"`
	IsRfidCapable            bool              `db:"is_rfid_capable" json:"isRfidCapable"`
	IsRemoved                bool              `db:"is_removed" json:"isRemoved"`
	OperatorID               sql.NullInt64     `db:"operator_id" json:"operatorId"`
	SuboperatorID            sql.NullInt64     `db:"suboperator_id" json:"suboperatorId"`
	OwnerID                  sql.NullInt64     `db:"owner_id" json:"ownerId"`
	TimeZone                 sql.NullString    `db:"time_zone" json:"timeZone"`
	OpeningTimeID            sql.NullInt64     `db:"opening_time_id" json:"openingTimeId"`
	ChargingWhenClosed       bool              `db:"charging_when_closed" json:"chargingWhenClosed"`
	EnergyMixID              sql.NullInt64     `db:"energy_mix_id" json:"energyMixId"`
	LastUpdated              time.Time         `db:"last_updated" json:"lastUpdated"`
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
		arg.AvailableEvses,
		arg.TotalEvses,
		arg.IsIntermediateCdrCapable,
		arg.IsRemoteCapable,
		arg.IsRfidCapable,
		arg.IsRemoved,
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
		&i.CredentialID,
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
		&i.AvailableEvses,
		&i.TotalEvses,
		&i.IsRemoteCapable,
		&i.IsRfidCapable,
		&i.OperatorID,
		&i.SuboperatorID,
		&i.OwnerID,
		&i.TimeZone,
		&i.OpeningTimeID,
		&i.ChargingWhenClosed,
		&i.EnergyMixID,
		&i.LastUpdated,
		&i.IsPublished,
		&i.AddedDate,
		&i.IsIntermediateCdrCapable,
		&i.IsRemoved,
	)
	return i, err
}

const updateLocationLastUpdated = `-- name: UpdateLocationLastUpdated :exec
UPDATE locations SET (
    available_evses, 
    total_evses, 
    is_intermediate_cdr_capable,
    is_remote_capable, 
    is_rfid_capable,
    is_removed,
    last_updated
  ) = ($2, $3, $4, $5, $6, $7, $8)
  WHERE id = $1
`

type UpdateLocationLastUpdatedParams struct {
	ID                       int64     `db:"id" json:"id"`
	AvailableEvses           int32     `db:"available_evses" json:"availableEvses"`
	TotalEvses               int32     `db:"total_evses" json:"totalEvses"`
	IsIntermediateCdrCapable bool      `db:"is_intermediate_cdr_capable" json:"isIntermediateCdrCapable"`
	IsRemoteCapable          bool      `db:"is_remote_capable" json:"isRemoteCapable"`
	IsRfidCapable            bool      `db:"is_rfid_capable" json:"isRfidCapable"`
	IsRemoved                bool      `db:"is_removed" json:"isRemoved"`
	LastUpdated              time.Time `db:"last_updated" json:"lastUpdated"`
}

func (q *Queries) UpdateLocationLastUpdated(ctx context.Context, arg UpdateLocationLastUpdatedParams) error {
	_, err := q.db.ExecContext(ctx, updateLocationLastUpdated,
		arg.ID,
		arg.AvailableEvses,
		arg.TotalEvses,
		arg.IsIntermediateCdrCapable,
		arg.IsRemoteCapable,
		arg.IsRfidCapable,
		arg.IsRemoved,
		arg.LastUpdated,
	)
	return err
}

const updateLocationPublished = `-- name: UpdateLocationPublished :exec
UPDATE locations SET is_published = $2
  WHERE id = $1
`

type UpdateLocationPublishedParams struct {
	ID          int64 `db:"id" json:"id"`
	IsPublished bool  `db:"is_published" json:"isPublished"`
}

func (q *Queries) UpdateLocationPublished(ctx context.Context, arg UpdateLocationPublishedParams) error {
	_, err := q.db.ExecContext(ctx, updateLocationPublished, arg.ID, arg.IsPublished)
	return err
}

const updateLocationsPublishedByCredential = `-- name: UpdateLocationsPublishedByCredential :exec
UPDATE locations SET is_published = $2
  WHERE credential_id = $1
`

type UpdateLocationsPublishedByCredentialParams struct {
	CredentialID int64 `db:"credential_id" json:"credentialId"`
	IsPublished  bool  `db:"is_published" json:"isPublished"`
}

func (q *Queries) UpdateLocationsPublishedByCredential(ctx context.Context, arg UpdateLocationsPublishedByCredentialParams) error {
	_, err := q.db.ExecContext(ctx, updateLocationsPublishedByCredential, arg.CredentialID, arg.IsPublished)
	return err
}

const updateLocationsPublishedByPartyAndCountryCode = `-- name: UpdateLocationsPublishedByPartyAndCountryCode :exec
UPDATE locations SET is_published = $3
  WHERE party_id = $1 AND country_code = $2
`

type UpdateLocationsPublishedByPartyAndCountryCodeParams struct {
	PartyID     sql.NullString `db:"party_id" json:"partyId"`
	CountryCode sql.NullString `db:"country_code" json:"countryCode"`
	IsPublished bool           `db:"is_published" json:"isPublished"`
}

func (q *Queries) UpdateLocationsPublishedByPartyAndCountryCode(ctx context.Context, arg UpdateLocationsPublishedByPartyAndCountryCodeParams) error {
	_, err := q.db.ExecContext(ctx, updateLocationsPublishedByPartyAndCountryCode, arg.PartyID, arg.CountryCode, arg.IsPublished)
	return err
}

const updateLocationsRemovedByCredential = `-- name: UpdateLocationsRemovedByCredential :exec
UPDATE locations SET is_removed = $2
  WHERE credential_id = $1
`

type UpdateLocationsRemovedByCredentialParams struct {
	CredentialID int64 `db:"credential_id" json:"credentialId"`
	IsRemoved    bool  `db:"is_removed" json:"isRemoved"`
}

func (q *Queries) UpdateLocationsRemovedByCredential(ctx context.Context, arg UpdateLocationsRemovedByCredentialParams) error {
	_, err := q.db.ExecContext(ctx, updateLocationsRemovedByCredential, arg.CredentialID, arg.IsRemoved)
	return err
}

const updateLocationsRemovedByPartyAndCountryCode = `-- name: UpdateLocationsRemovedByPartyAndCountryCode :exec
UPDATE locations SET is_removed = $3
  WHERE party_id = $1 AND country_code = $2
`

type UpdateLocationsRemovedByPartyAndCountryCodeParams struct {
	PartyID     sql.NullString `db:"party_id" json:"partyId"`
	CountryCode sql.NullString `db:"country_code" json:"countryCode"`
	IsRemoved   bool           `db:"is_removed" json:"isRemoved"`
}

func (q *Queries) UpdateLocationsRemovedByPartyAndCountryCode(ctx context.Context, arg UpdateLocationsRemovedByPartyAndCountryCodeParams) error {
	_, err := q.db.ExecContext(ctx, updateLocationsRemovedByPartyAndCountryCode, arg.PartyID, arg.CountryCode, arg.IsRemoved)
	return err
}
