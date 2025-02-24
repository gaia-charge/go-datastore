// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: credential.sql

package db

import (
	"context"
	"database/sql"
	"time"
)

const createCredential = `-- name: CreateCredential :one
INSERT INTO credentials (
    client_token, 
    server_token, 
    url, 
    business_detail_id, 
    country_code, 
    party_id, 
    is_hub,
    is_base64,
    last_updated
  ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
  RETURNING id, client_token, server_token, url, country_code, party_id, is_hub, last_updated, business_detail_id, version_id, is_available, is_base64
`

type CreateCredentialParams struct {
	ClientToken      sql.NullString `db:"client_token" json:"clientToken"`
	ServerToken      sql.NullString `db:"server_token" json:"serverToken"`
	Url              string         `db:"url" json:"url"`
	BusinessDetailID int64          `db:"business_detail_id" json:"businessDetailId"`
	CountryCode      string         `db:"country_code" json:"countryCode"`
	PartyID          string         `db:"party_id" json:"partyId"`
	IsHub            bool           `db:"is_hub" json:"isHub"`
	IsBase64         bool           `db:"is_base64" json:"isBase64"`
	LastUpdated      time.Time      `db:"last_updated" json:"lastUpdated"`
}

func (q *Queries) CreateCredential(ctx context.Context, arg CreateCredentialParams) (Credential, error) {
	row := q.db.QueryRowContext(ctx, createCredential,
		arg.ClientToken,
		arg.ServerToken,
		arg.Url,
		arg.BusinessDetailID,
		arg.CountryCode,
		arg.PartyID,
		arg.IsHub,
		arg.IsBase64,
		arg.LastUpdated,
	)
	var i Credential
	err := row.Scan(
		&i.ID,
		&i.ClientToken,
		&i.ServerToken,
		&i.Url,
		&i.CountryCode,
		&i.PartyID,
		&i.IsHub,
		&i.LastUpdated,
		&i.BusinessDetailID,
		&i.VersionID,
		&i.IsAvailable,
		&i.IsBase64,
	)
	return i, err
}

const deleteCredential = `-- name: DeleteCredential :exec
DELETE FROM credentials
  WHERE id = $1
`

func (q *Queries) DeleteCredential(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteCredential, id)
	return err
}

const getCredential = `-- name: GetCredential :one
SELECT id, client_token, server_token, url, country_code, party_id, is_hub, last_updated, business_detail_id, version_id, is_available, is_base64 FROM credentials
  WHERE id = $1
`

func (q *Queries) GetCredential(ctx context.Context, id int64) (Credential, error) {
	row := q.db.QueryRowContext(ctx, getCredential, id)
	var i Credential
	err := row.Scan(
		&i.ID,
		&i.ClientToken,
		&i.ServerToken,
		&i.Url,
		&i.CountryCode,
		&i.PartyID,
		&i.IsHub,
		&i.LastUpdated,
		&i.BusinessDetailID,
		&i.VersionID,
		&i.IsAvailable,
		&i.IsBase64,
	)
	return i, err
}

const getCredentialByPartyAndCountryCode = `-- name: GetCredentialByPartyAndCountryCode :one
SELECT id, client_token, server_token, url, country_code, party_id, is_hub, last_updated, business_detail_id, version_id, is_available, is_base64 FROM credentials
  WHERE party_id = $1 AND country_code = $2
`

type GetCredentialByPartyAndCountryCodeParams struct {
	PartyID     string `db:"party_id" json:"partyId"`
	CountryCode string `db:"country_code" json:"countryCode"`
}

func (q *Queries) GetCredentialByPartyAndCountryCode(ctx context.Context, arg GetCredentialByPartyAndCountryCodeParams) (Credential, error) {
	row := q.db.QueryRowContext(ctx, getCredentialByPartyAndCountryCode, arg.PartyID, arg.CountryCode)
	var i Credential
	err := row.Scan(
		&i.ID,
		&i.ClientToken,
		&i.ServerToken,
		&i.Url,
		&i.CountryCode,
		&i.PartyID,
		&i.IsHub,
		&i.LastUpdated,
		&i.BusinessDetailID,
		&i.VersionID,
		&i.IsAvailable,
		&i.IsBase64,
	)
	return i, err
}

const getCredentialByServerToken = `-- name: GetCredentialByServerToken :one
SELECT id, client_token, server_token, url, country_code, party_id, is_hub, last_updated, business_detail_id, version_id, is_available, is_base64 FROM credentials
  WHERE server_token = $1
`

func (q *Queries) GetCredentialByServerToken(ctx context.Context, serverToken sql.NullString) (Credential, error) {
	row := q.db.QueryRowContext(ctx, getCredentialByServerToken, serverToken)
	var i Credential
	err := row.Scan(
		&i.ID,
		&i.ClientToken,
		&i.ServerToken,
		&i.Url,
		&i.CountryCode,
		&i.PartyID,
		&i.IsHub,
		&i.LastUpdated,
		&i.BusinessDetailID,
		&i.VersionID,
		&i.IsAvailable,
		&i.IsBase64,
	)
	return i, err
}

const listCredentials = `-- name: ListCredentials :many
SELECT id, client_token, server_token, url, country_code, party_id, is_hub, last_updated, business_detail_id, version_id, is_available, is_base64 FROM credentials
  ORDER BY id
`

func (q *Queries) ListCredentials(ctx context.Context) ([]Credential, error) {
	rows, err := q.db.QueryContext(ctx, listCredentials)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Credential
	for rows.Next() {
		var i Credential
		if err := rows.Scan(
			&i.ID,
			&i.ClientToken,
			&i.ServerToken,
			&i.Url,
			&i.CountryCode,
			&i.PartyID,
			&i.IsHub,
			&i.LastUpdated,
			&i.BusinessDetailID,
			&i.VersionID,
			&i.IsAvailable,
			&i.IsBase64,
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

const updateCredential = `-- name: UpdateCredential :one
UPDATE credentials SET (
    client_token, 
    server_token, 
    url, 
    country_code,
    party_id, 
    is_available,
    is_hub,
    is_base64,
    version_id,
    last_updated
  ) = ($2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
  WHERE id = $1
  RETURNING id, client_token, server_token, url, country_code, party_id, is_hub, last_updated, business_detail_id, version_id, is_available, is_base64
`

type UpdateCredentialParams struct {
	ID          int64          `db:"id" json:"id"`
	ClientToken sql.NullString `db:"client_token" json:"clientToken"`
	ServerToken sql.NullString `db:"server_token" json:"serverToken"`
	Url         string         `db:"url" json:"url"`
	CountryCode string         `db:"country_code" json:"countryCode"`
	PartyID     string         `db:"party_id" json:"partyId"`
	IsAvailable bool           `db:"is_available" json:"isAvailable"`
	IsHub       bool           `db:"is_hub" json:"isHub"`
	IsBase64    bool           `db:"is_base64" json:"isBase64"`
	VersionID   sql.NullInt64  `db:"version_id" json:"versionId"`
	LastUpdated time.Time      `db:"last_updated" json:"lastUpdated"`
}

func (q *Queries) UpdateCredential(ctx context.Context, arg UpdateCredentialParams) (Credential, error) {
	row := q.db.QueryRowContext(ctx, updateCredential,
		arg.ID,
		arg.ClientToken,
		arg.ServerToken,
		arg.Url,
		arg.CountryCode,
		arg.PartyID,
		arg.IsAvailable,
		arg.IsHub,
		arg.IsBase64,
		arg.VersionID,
		arg.LastUpdated,
	)
	var i Credential
	err := row.Scan(
		&i.ID,
		&i.ClientToken,
		&i.ServerToken,
		&i.Url,
		&i.CountryCode,
		&i.PartyID,
		&i.IsHub,
		&i.LastUpdated,
		&i.BusinessDetailID,
		&i.VersionID,
		&i.IsAvailable,
		&i.IsBase64,
	)
	return i, err
}
