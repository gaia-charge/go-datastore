// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: connector.sql

package db

import (
	"context"
	"database/sql"
	"time"
)

const createConnector = `-- name: CreateConnector :one
INSERT INTO connectors (
    evse_id, 
    uid, 
    identifier, 
    standard, 
    format, 
    power_type, 
    voltage, 
    amperage, 
    wattage, 
    tariff_id, 
    terms_and_conditions, 
    is_published,
    last_updated)
  VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)
  RETURNING id, evse_id, uid, identifier, standard, format, power_type, voltage, amperage, wattage, tariff_id, terms_and_conditions, last_updated, is_published, is_removed
`

type CreateConnectorParams struct {
	EvseID             int64           `db:"evse_id" json:"evseId"`
	Uid                string          `db:"uid" json:"uid"`
	Identifier         sql.NullString  `db:"identifier" json:"identifier"`
	Standard           ConnectorType   `db:"standard" json:"standard"`
	Format             ConnectorFormat `db:"format" json:"format"`
	PowerType          PowerType       `db:"power_type" json:"powerType"`
	Voltage            int32           `db:"voltage" json:"voltage"`
	Amperage           int32           `db:"amperage" json:"amperage"`
	Wattage            int32           `db:"wattage" json:"wattage"`
	TariffID           sql.NullString  `db:"tariff_id" json:"tariffId"`
	TermsAndConditions sql.NullString  `db:"terms_and_conditions" json:"termsAndConditions"`
	IsPublished        bool            `db:"is_published" json:"isPublished"`
	LastUpdated        time.Time       `db:"last_updated" json:"lastUpdated"`
}

func (q *Queries) CreateConnector(ctx context.Context, arg CreateConnectorParams) (Connector, error) {
	row := q.db.QueryRowContext(ctx, createConnector,
		arg.EvseID,
		arg.Uid,
		arg.Identifier,
		arg.Standard,
		arg.Format,
		arg.PowerType,
		arg.Voltage,
		arg.Amperage,
		arg.Wattage,
		arg.TariffID,
		arg.TermsAndConditions,
		arg.IsPublished,
		arg.LastUpdated,
	)
	var i Connector
	err := row.Scan(
		&i.ID,
		&i.EvseID,
		&i.Uid,
		&i.Identifier,
		&i.Standard,
		&i.Format,
		&i.PowerType,
		&i.Voltage,
		&i.Amperage,
		&i.Wattage,
		&i.TariffID,
		&i.TermsAndConditions,
		&i.LastUpdated,
		&i.IsPublished,
		&i.IsRemoved,
	)
	return i, err
}

const deleteConnector = `-- name: DeleteConnector :exec
DELETE FROM connectors
  WHERE id = $1
`

func (q *Queries) DeleteConnector(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteConnector, id)
	return err
}

const deleteConnectorByEvse = `-- name: DeleteConnectorByEvse :exec
DELETE FROM connectors
  WHERE evse_id = $1 AND uid = $2
`

type DeleteConnectorByEvseParams struct {
	EvseID int64  `db:"evse_id" json:"evseId"`
	Uid    string `db:"uid" json:"uid"`
}

func (q *Queries) DeleteConnectorByEvse(ctx context.Context, arg DeleteConnectorByEvseParams) error {
	_, err := q.db.ExecContext(ctx, deleteConnectorByEvse, arg.EvseID, arg.Uid)
	return err
}

const deleteConnectors = `-- name: DeleteConnectors :exec
DELETE FROM connectors
  WHERE evse_id = $1
`

func (q *Queries) DeleteConnectors(ctx context.Context, evseID int64) error {
	_, err := q.db.ExecContext(ctx, deleteConnectors, evseID)
	return err
}

const getConnector = `-- name: GetConnector :one
SELECT id, evse_id, uid, identifier, standard, format, power_type, voltage, amperage, wattage, tariff_id, terms_and_conditions, last_updated, is_published, is_removed FROM connectors
  WHERE id = $1
`

func (q *Queries) GetConnector(ctx context.Context, id int64) (Connector, error) {
	row := q.db.QueryRowContext(ctx, getConnector, id)
	var i Connector
	err := row.Scan(
		&i.ID,
		&i.EvseID,
		&i.Uid,
		&i.Identifier,
		&i.Standard,
		&i.Format,
		&i.PowerType,
		&i.Voltage,
		&i.Amperage,
		&i.Wattage,
		&i.TariffID,
		&i.TermsAndConditions,
		&i.LastUpdated,
		&i.IsPublished,
		&i.IsRemoved,
	)
	return i, err
}

const getConnectorByEvse = `-- name: GetConnectorByEvse :one
SELECT id, evse_id, uid, identifier, standard, format, power_type, voltage, amperage, wattage, tariff_id, terms_and_conditions, last_updated, is_published, is_removed FROM connectors
  WHERE evse_id = $1 AND uid = $2
`

type GetConnectorByEvseParams struct {
	EvseID int64  `db:"evse_id" json:"evseId"`
	Uid    string `db:"uid" json:"uid"`
}

func (q *Queries) GetConnectorByEvse(ctx context.Context, arg GetConnectorByEvseParams) (Connector, error) {
	row := q.db.QueryRowContext(ctx, getConnectorByEvse, arg.EvseID, arg.Uid)
	var i Connector
	err := row.Scan(
		&i.ID,
		&i.EvseID,
		&i.Uid,
		&i.Identifier,
		&i.Standard,
		&i.Format,
		&i.PowerType,
		&i.Voltage,
		&i.Amperage,
		&i.Wattage,
		&i.TariffID,
		&i.TermsAndConditions,
		&i.LastUpdated,
		&i.IsPublished,
		&i.IsRemoved,
	)
	return i, err
}

const getConnectorByIdentifier = `-- name: GetConnectorByIdentifier :one
SELECT id, evse_id, uid, identifier, standard, format, power_type, voltage, amperage, wattage, tariff_id, terms_and_conditions, last_updated, is_published, is_removed FROM connectors
  WHERE identifier = $1
`

func (q *Queries) GetConnectorByIdentifier(ctx context.Context, identifier sql.NullString) (Connector, error) {
	row := q.db.QueryRowContext(ctx, getConnectorByIdentifier, identifier)
	var i Connector
	err := row.Scan(
		&i.ID,
		&i.EvseID,
		&i.Uid,
		&i.Identifier,
		&i.Standard,
		&i.Format,
		&i.PowerType,
		&i.Voltage,
		&i.Amperage,
		&i.Wattage,
		&i.TariffID,
		&i.TermsAndConditions,
		&i.LastUpdated,
		&i.IsPublished,
		&i.IsRemoved,
	)
	return i, err
}

const listConnectors = `-- name: ListConnectors :many
SELECT id, evse_id, uid, identifier, standard, format, power_type, voltage, amperage, wattage, tariff_id, terms_and_conditions, last_updated, is_published, is_removed FROM connectors
  WHERE evse_id = $1 AND is_published = true AND is_removed = false
  ORDER BY id
`

func (q *Queries) ListConnectors(ctx context.Context, evseID int64) ([]Connector, error) {
	rows, err := q.db.QueryContext(ctx, listConnectors, evseID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Connector
	for rows.Next() {
		var i Connector
		if err := rows.Scan(
			&i.ID,
			&i.EvseID,
			&i.Uid,
			&i.Identifier,
			&i.Standard,
			&i.Format,
			&i.PowerType,
			&i.Voltage,
			&i.Amperage,
			&i.Wattage,
			&i.TariffID,
			&i.TermsAndConditions,
			&i.LastUpdated,
			&i.IsPublished,
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

const listConnectorsByEvseID = `-- name: ListConnectorsByEvseID :many
SELECT c.id, c.evse_id, c.uid, c.identifier, c.standard, c.format, c.power_type, c.voltage, c.amperage, c.wattage, c.tariff_id, c.terms_and_conditions, c.last_updated, c.is_published, c.is_removed FROM connectors c
  INNER JOIN evses e ON c.evse_id = e.id
  WHERE e.evse_id = $1
`

func (q *Queries) ListConnectorsByEvseID(ctx context.Context, evseID sql.NullString) ([]Connector, error) {
	rows, err := q.db.QueryContext(ctx, listConnectorsByEvseID, evseID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Connector
	for rows.Next() {
		var i Connector
		if err := rows.Scan(
			&i.ID,
			&i.EvseID,
			&i.Uid,
			&i.Identifier,
			&i.Standard,
			&i.Format,
			&i.PowerType,
			&i.Voltage,
			&i.Amperage,
			&i.Wattage,
			&i.TariffID,
			&i.TermsAndConditions,
			&i.LastUpdated,
			&i.IsPublished,
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

const listConnectorsByPartyAndCountryCode = `-- name: ListConnectorsByPartyAndCountryCode :many
SELECT c.id, c.evse_id, c.uid, c.identifier, c.standard, c.format, c.power_type, c.voltage, c.amperage, c.wattage, c.tariff_id, c.terms_and_conditions, c.last_updated, c.is_published, c.is_removed FROM connectors c
  INNER JOIN evses e ON c.evse_id = e.id
  INNER JOIN locations l ON e.location_id = l.id
  WHERE l.country_code = $1 AND l.party_id = $2
`

type ListConnectorsByPartyAndCountryCodeParams struct {
	CountryCode sql.NullString `db:"country_code" json:"countryCode"`
	PartyID     sql.NullString `db:"party_id" json:"partyId"`
}

func (q *Queries) ListConnectorsByPartyAndCountryCode(ctx context.Context, arg ListConnectorsByPartyAndCountryCodeParams) ([]Connector, error) {
	rows, err := q.db.QueryContext(ctx, listConnectorsByPartyAndCountryCode, arg.CountryCode, arg.PartyID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Connector
	for rows.Next() {
		var i Connector
		if err := rows.Scan(
			&i.ID,
			&i.EvseID,
			&i.Uid,
			&i.Identifier,
			&i.Standard,
			&i.Format,
			&i.PowerType,
			&i.Voltage,
			&i.Amperage,
			&i.Wattage,
			&i.TariffID,
			&i.TermsAndConditions,
			&i.LastUpdated,
			&i.IsPublished,
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

const updateConnector = `-- name: UpdateConnector :one
UPDATE connectors SET (
    identifier,
    standard, 
    format, 
    power_type, 
    voltage, 
    amperage, 
    wattage, 
    tariff_id, 
    terms_and_conditions, 
    is_published,
    is_removed,
    last_updated
  ) = ($2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)
  WHERE id = $1
  RETURNING id, evse_id, uid, identifier, standard, format, power_type, voltage, amperage, wattage, tariff_id, terms_and_conditions, last_updated, is_published, is_removed
`

type UpdateConnectorParams struct {
	ID                 int64           `db:"id" json:"id"`
	Identifier         sql.NullString  `db:"identifier" json:"identifier"`
	Standard           ConnectorType   `db:"standard" json:"standard"`
	Format             ConnectorFormat `db:"format" json:"format"`
	PowerType          PowerType       `db:"power_type" json:"powerType"`
	Voltage            int32           `db:"voltage" json:"voltage"`
	Amperage           int32           `db:"amperage" json:"amperage"`
	Wattage            int32           `db:"wattage" json:"wattage"`
	TariffID           sql.NullString  `db:"tariff_id" json:"tariffId"`
	TermsAndConditions sql.NullString  `db:"terms_and_conditions" json:"termsAndConditions"`
	IsPublished        bool            `db:"is_published" json:"isPublished"`
	IsRemoved          bool            `db:"is_removed" json:"isRemoved"`
	LastUpdated        time.Time       `db:"last_updated" json:"lastUpdated"`
}

func (q *Queries) UpdateConnector(ctx context.Context, arg UpdateConnectorParams) (Connector, error) {
	row := q.db.QueryRowContext(ctx, updateConnector,
		arg.ID,
		arg.Identifier,
		arg.Standard,
		arg.Format,
		arg.PowerType,
		arg.Voltage,
		arg.Amperage,
		arg.Wattage,
		arg.TariffID,
		arg.TermsAndConditions,
		arg.IsPublished,
		arg.IsRemoved,
		arg.LastUpdated,
	)
	var i Connector
	err := row.Scan(
		&i.ID,
		&i.EvseID,
		&i.Uid,
		&i.Identifier,
		&i.Standard,
		&i.Format,
		&i.PowerType,
		&i.Voltage,
		&i.Amperage,
		&i.Wattage,
		&i.TariffID,
		&i.TermsAndConditions,
		&i.LastUpdated,
		&i.IsPublished,
		&i.IsRemoved,
	)
	return i, err
}

const updateConnectorByEvse = `-- name: UpdateConnectorByEvse :one
UPDATE connectors SET (
    identifier,
    standard, 
    format, 
    power_type, 
    voltage, 
    amperage, 
    wattage, 
    tariff_id, 
    terms_and_conditions, 
    is_published,
    is_removed,
    last_updated
  ) = ($3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)
  WHERE evse_id = $1 AND uid = $2
  RETURNING id, evse_id, uid, identifier, standard, format, power_type, voltage, amperage, wattage, tariff_id, terms_and_conditions, last_updated, is_published, is_removed
`

type UpdateConnectorByEvseParams struct {
	EvseID             int64           `db:"evse_id" json:"evseId"`
	Uid                string          `db:"uid" json:"uid"`
	Identifier         sql.NullString  `db:"identifier" json:"identifier"`
	Standard           ConnectorType   `db:"standard" json:"standard"`
	Format             ConnectorFormat `db:"format" json:"format"`
	PowerType          PowerType       `db:"power_type" json:"powerType"`
	Voltage            int32           `db:"voltage" json:"voltage"`
	Amperage           int32           `db:"amperage" json:"amperage"`
	Wattage            int32           `db:"wattage" json:"wattage"`
	TariffID           sql.NullString  `db:"tariff_id" json:"tariffId"`
	TermsAndConditions sql.NullString  `db:"terms_and_conditions" json:"termsAndConditions"`
	IsPublished        bool            `db:"is_published" json:"isPublished"`
	IsRemoved          bool            `db:"is_removed" json:"isRemoved"`
	LastUpdated        time.Time       `db:"last_updated" json:"lastUpdated"`
}

func (q *Queries) UpdateConnectorByEvse(ctx context.Context, arg UpdateConnectorByEvseParams) (Connector, error) {
	row := q.db.QueryRowContext(ctx, updateConnectorByEvse,
		arg.EvseID,
		arg.Uid,
		arg.Identifier,
		arg.Standard,
		arg.Format,
		arg.PowerType,
		arg.Voltage,
		arg.Amperage,
		arg.Wattage,
		arg.TariffID,
		arg.TermsAndConditions,
		arg.IsPublished,
		arg.IsRemoved,
		arg.LastUpdated,
	)
	var i Connector
	err := row.Scan(
		&i.ID,
		&i.EvseID,
		&i.Uid,
		&i.Identifier,
		&i.Standard,
		&i.Format,
		&i.PowerType,
		&i.Voltage,
		&i.Amperage,
		&i.Wattage,
		&i.TariffID,
		&i.TermsAndConditions,
		&i.LastUpdated,
		&i.IsPublished,
		&i.IsRemoved,
	)
	return i, err
}
