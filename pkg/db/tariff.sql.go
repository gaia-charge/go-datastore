// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: tariff.sql

package db

import (
	"context"
	"database/sql"
	"time"
)

const createTariff = `-- name: CreateTariff :one
INSERT INTO tariffs (
    uid, 
    credential_id,
    country_code,
    party_id,
    cdr_id,
    currency, 
    tariff_alt_url, 
    energy_mix_id, 
    tariff_restriction_id,
    last_updated
  ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
  RETURNING id, uid, credential_id, country_code, party_id, currency, tariff_alt_url, energy_mix_id, tariff_restriction_id, last_updated, cdr_id
`

type CreateTariffParams struct {
	Uid                 string         `db:"uid" json:"uid"`
	CredentialID        int64          `db:"credential_id" json:"credentialId"`
	CountryCode         sql.NullString `db:"country_code" json:"countryCode"`
	PartyID             sql.NullString `db:"party_id" json:"partyId"`
	CdrID               sql.NullInt64  `db:"cdr_id" json:"cdrId"`
	Currency            string         `db:"currency" json:"currency"`
	TariffAltUrl        sql.NullString `db:"tariff_alt_url" json:"tariffAltUrl"`
	EnergyMixID         sql.NullInt64  `db:"energy_mix_id" json:"energyMixId"`
	TariffRestrictionID sql.NullInt64  `db:"tariff_restriction_id" json:"tariffRestrictionId"`
	LastUpdated         time.Time      `db:"last_updated" json:"lastUpdated"`
}

func (q *Queries) CreateTariff(ctx context.Context, arg CreateTariffParams) (Tariff, error) {
	row := q.db.QueryRowContext(ctx, createTariff,
		arg.Uid,
		arg.CredentialID,
		arg.CountryCode,
		arg.PartyID,
		arg.CdrID,
		arg.Currency,
		arg.TariffAltUrl,
		arg.EnergyMixID,
		arg.TariffRestrictionID,
		arg.LastUpdated,
	)
	var i Tariff
	err := row.Scan(
		&i.ID,
		&i.Uid,
		&i.CredentialID,
		&i.CountryCode,
		&i.PartyID,
		&i.Currency,
		&i.TariffAltUrl,
		&i.EnergyMixID,
		&i.TariffRestrictionID,
		&i.LastUpdated,
		&i.CdrID,
	)
	return i, err
}

const deleteTariffByUid = `-- name: DeleteTariffByUid :exec
DELETE FROM tariffs
  WHERE uid = $1 AND cdr_id IS NULL
`

func (q *Queries) DeleteTariffByUid(ctx context.Context, uid string) error {
	_, err := q.db.ExecContext(ctx, deleteTariffByUid, uid)
	return err
}

const getTariff = `-- name: GetTariff :one
SELECT id, uid, credential_id, country_code, party_id, currency, tariff_alt_url, energy_mix_id, tariff_restriction_id, last_updated, cdr_id FROM tariffs
  WHERE id = $1
`

func (q *Queries) GetTariff(ctx context.Context, id int64) (Tariff, error) {
	row := q.db.QueryRowContext(ctx, getTariff, id)
	var i Tariff
	err := row.Scan(
		&i.ID,
		&i.Uid,
		&i.CredentialID,
		&i.CountryCode,
		&i.PartyID,
		&i.Currency,
		&i.TariffAltUrl,
		&i.EnergyMixID,
		&i.TariffRestrictionID,
		&i.LastUpdated,
		&i.CdrID,
	)
	return i, err
}

const getTariffByLastUpdated = `-- name: GetTariffByLastUpdated :one
SELECT id, uid, credential_id, country_code, party_id, currency, tariff_alt_url, energy_mix_id, tariff_restriction_id, last_updated, cdr_id FROM tariffs
  WHERE ($1::BIGINT = -1 OR $1::BIGINT = credential_id) AND
    ($2::TEXT = '' OR $2::TEXT = country_code) AND
    ($3::TEXT = '' OR $3::TEXT = party_id)
  ORDER BY last_updated DESC
  LIMIT 1
`

type GetTariffByLastUpdatedParams struct {
	CredentialID int64  `db:"credential_id" json:"credentialId"`
	CountryCode  string `db:"country_code" json:"countryCode"`
	PartyID      string `db:"party_id" json:"partyId"`
}

func (q *Queries) GetTariffByLastUpdated(ctx context.Context, arg GetTariffByLastUpdatedParams) (Tariff, error) {
	row := q.db.QueryRowContext(ctx, getTariffByLastUpdated, arg.CredentialID, arg.CountryCode, arg.PartyID)
	var i Tariff
	err := row.Scan(
		&i.ID,
		&i.Uid,
		&i.CredentialID,
		&i.CountryCode,
		&i.PartyID,
		&i.Currency,
		&i.TariffAltUrl,
		&i.EnergyMixID,
		&i.TariffRestrictionID,
		&i.LastUpdated,
		&i.CdrID,
	)
	return i, err
}

const getTariffByUid = `-- name: GetTariffByUid :one
SELECT id, uid, credential_id, country_code, party_id, currency, tariff_alt_url, energy_mix_id, tariff_restriction_id, last_updated, cdr_id FROM tariffs
  WHERE uid = $1 AND cdr_id IS NULL
`

func (q *Queries) GetTariffByUid(ctx context.Context, uid string) (Tariff, error) {
	row := q.db.QueryRowContext(ctx, getTariffByUid, uid)
	var i Tariff
	err := row.Scan(
		&i.ID,
		&i.Uid,
		&i.CredentialID,
		&i.CountryCode,
		&i.PartyID,
		&i.Currency,
		&i.TariffAltUrl,
		&i.EnergyMixID,
		&i.TariffRestrictionID,
		&i.LastUpdated,
		&i.CdrID,
	)
	return i, err
}

const getTariffLikeUid = `-- name: GetTariffLikeUid :one
SELECT id, uid, credential_id, country_code, party_id, currency, tariff_alt_url, energy_mix_id, tariff_restriction_id, last_updated, cdr_id FROM tariffs
  WHERE uid like $1
  LIMIT 1
`

func (q *Queries) GetTariffLikeUid(ctx context.Context, uid string) (Tariff, error) {
	row := q.db.QueryRowContext(ctx, getTariffLikeUid, uid)
	var i Tariff
	err := row.Scan(
		&i.ID,
		&i.Uid,
		&i.CredentialID,
		&i.CountryCode,
		&i.PartyID,
		&i.Currency,
		&i.TariffAltUrl,
		&i.EnergyMixID,
		&i.TariffRestrictionID,
		&i.LastUpdated,
		&i.CdrID,
	)
	return i, err
}

const listTariffsByCdr = `-- name: ListTariffsByCdr :many
SELECT id, uid, credential_id, country_code, party_id, currency, tariff_alt_url, energy_mix_id, tariff_restriction_id, last_updated, cdr_id FROM tariffs
  WHERE cdr_id = $1
  ORDER BY id
`

func (q *Queries) ListTariffsByCdr(ctx context.Context, cdrID sql.NullInt64) ([]Tariff, error) {
	rows, err := q.db.QueryContext(ctx, listTariffsByCdr, cdrID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Tariff
	for rows.Next() {
		var i Tariff
		if err := rows.Scan(
			&i.ID,
			&i.Uid,
			&i.CredentialID,
			&i.CountryCode,
			&i.PartyID,
			&i.Currency,
			&i.TariffAltUrl,
			&i.EnergyMixID,
			&i.TariffRestrictionID,
			&i.LastUpdated,
			&i.CdrID,
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

const updateTariffByUid = `-- name: UpdateTariffByUid :one
UPDATE tariffs SET (
    country_code,
    party_id,
    currency, 
    tariff_alt_url,
    energy_mix_id, 
    tariff_restriction_id,
    last_updated
  ) = ($2, $3, $4, $5, $6, $7, $8)
  WHERE uid = $1 AND cdr_id IS NULL
  RETURNING id, uid, credential_id, country_code, party_id, currency, tariff_alt_url, energy_mix_id, tariff_restriction_id, last_updated, cdr_id
`

type UpdateTariffByUidParams struct {
	Uid                 string         `db:"uid" json:"uid"`
	CountryCode         sql.NullString `db:"country_code" json:"countryCode"`
	PartyID             sql.NullString `db:"party_id" json:"partyId"`
	Currency            string         `db:"currency" json:"currency"`
	TariffAltUrl        sql.NullString `db:"tariff_alt_url" json:"tariffAltUrl"`
	EnergyMixID         sql.NullInt64  `db:"energy_mix_id" json:"energyMixId"`
	TariffRestrictionID sql.NullInt64  `db:"tariff_restriction_id" json:"tariffRestrictionId"`
	LastUpdated         time.Time      `db:"last_updated" json:"lastUpdated"`
}

func (q *Queries) UpdateTariffByUid(ctx context.Context, arg UpdateTariffByUidParams) (Tariff, error) {
	row := q.db.QueryRowContext(ctx, updateTariffByUid,
		arg.Uid,
		arg.CountryCode,
		arg.PartyID,
		arg.Currency,
		arg.TariffAltUrl,
		arg.EnergyMixID,
		arg.TariffRestrictionID,
		arg.LastUpdated,
	)
	var i Tariff
	err := row.Scan(
		&i.ID,
		&i.Uid,
		&i.CredentialID,
		&i.CountryCode,
		&i.PartyID,
		&i.Currency,
		&i.TariffAltUrl,
		&i.EnergyMixID,
		&i.TariffRestrictionID,
		&i.LastUpdated,
		&i.CdrID,
	)
	return i, err
}
