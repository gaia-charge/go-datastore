// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: token_authorization.sql

package db

import (
	"context"
	"database/sql"
)

const createTokenAuthorization = `-- name: CreateTokenAuthorization :one
INSERT INTO token_authorizations (
    token_id,
    authorization_id,
    authorized,
    country_code,
    party_id,
    location_id
  ) VALUES ($1, $2, $3, $4, $5, $6)
  RETURNING id, token_id, authorization_id, country_code, party_id, location_id, authorized
`

type CreateTokenAuthorizationParams struct {
	TokenID         int64          `db:"token_id" json:"tokenID"`
	AuthorizationID string         `db:"authorization_id" json:"authorizationID"`
	Authorized      bool           `db:"authorized" json:"authorized"`
	CountryCode     sql.NullString `db:"country_code" json:"countryCode"`
	PartyID         sql.NullString `db:"party_id" json:"partyID"`
	LocationID      sql.NullString `db:"location_id" json:"locationID"`
}

func (q *Queries) CreateTokenAuthorization(ctx context.Context, arg CreateTokenAuthorizationParams) (TokenAuthorization, error) {
	row := q.db.QueryRowContext(ctx, createTokenAuthorization,
		arg.TokenID,
		arg.AuthorizationID,
		arg.Authorized,
		arg.CountryCode,
		arg.PartyID,
		arg.LocationID,
	)
	var i TokenAuthorization
	err := row.Scan(
		&i.ID,
		&i.TokenID,
		&i.AuthorizationID,
		&i.CountryCode,
		&i.PartyID,
		&i.LocationID,
		&i.Authorized,
	)
	return i, err
}

const getLastTokenAuthorizationByTokenID = `-- name: GetLastTokenAuthorizationByTokenID :one
SELECT ta.id, ta.token_id, ta.authorization_id, ta.country_code, ta.party_id, ta.location_id, ta.authorized FROM token_authorizations ta
  LEFT JOIN sessions s ON s.authorization_id = ta.authorization_id
  WHERE ta.token_id = $1 AND s.authorization_id is NULL
  ORDER BY ta.id DESC
  LIMIT 1
`

func (q *Queries) GetLastTokenAuthorizationByTokenID(ctx context.Context, tokenID int64) (TokenAuthorization, error) {
	row := q.db.QueryRowContext(ctx, getLastTokenAuthorizationByTokenID, tokenID)
	var i TokenAuthorization
	err := row.Scan(
		&i.ID,
		&i.TokenID,
		&i.AuthorizationID,
		&i.CountryCode,
		&i.PartyID,
		&i.LocationID,
		&i.Authorized,
	)
	return i, err
}

const getTokenAuthorizationByAuthorizationID = `-- name: GetTokenAuthorizationByAuthorizationID :one
SELECT id, token_id, authorization_id, country_code, party_id, location_id, authorized FROM token_authorizations
  WHERE authorization_id = $1
`

func (q *Queries) GetTokenAuthorizationByAuthorizationID(ctx context.Context, authorizationID string) (TokenAuthorization, error) {
	row := q.db.QueryRowContext(ctx, getTokenAuthorizationByAuthorizationID, authorizationID)
	var i TokenAuthorization
	err := row.Scan(
		&i.ID,
		&i.TokenID,
		&i.AuthorizationID,
		&i.CountryCode,
		&i.PartyID,
		&i.LocationID,
		&i.Authorized,
	)
	return i, err
}

const updateTokenAuthorizationByAuthorizationID = `-- name: UpdateTokenAuthorizationByAuthorizationID :one
UPDATE token_authorizations SET (
    authorized,
    country_code,
    party_id
  ) = ($2, $3, $4)
  WHERE authorization_id = $1
  RETURNING id, token_id, authorization_id, country_code, party_id, location_id, authorized
`

type UpdateTokenAuthorizationByAuthorizationIDParams struct {
	AuthorizationID string         `db:"authorization_id" json:"authorizationID"`
	Authorized      bool           `db:"authorized" json:"authorized"`
	CountryCode     sql.NullString `db:"country_code" json:"countryCode"`
	PartyID         sql.NullString `db:"party_id" json:"partyID"`
}

func (q *Queries) UpdateTokenAuthorizationByAuthorizationID(ctx context.Context, arg UpdateTokenAuthorizationByAuthorizationIDParams) (TokenAuthorization, error) {
	row := q.db.QueryRowContext(ctx, updateTokenAuthorizationByAuthorizationID,
		arg.AuthorizationID,
		arg.Authorized,
		arg.CountryCode,
		arg.PartyID,
	)
	var i TokenAuthorization
	err := row.Scan(
		&i.ID,
		&i.TokenID,
		&i.AuthorizationID,
		&i.CountryCode,
		&i.PartyID,
		&i.LocationID,
		&i.Authorized,
	)
	return i, err
}
