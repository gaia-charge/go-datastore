// Code generated by sqlc. DO NOT EDIT.
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
    country_code,
    party_id,
    location_id
  ) VALUES ($1, $2, $3, $4, $5)
  RETURNING id, token_id, authorization_id, country_code, party_id, location_id
`

type CreateTokenAuthorizationParams struct {
	TokenID         int64          `db:"token_id" json:"tokenID"`
	AuthorizationID string         `db:"authorization_id" json:"authorizationID"`
	CountryCode     sql.NullString `db:"country_code" json:"countryCode"`
	PartyID         sql.NullString `db:"party_id" json:"partyID"`
	LocationID      sql.NullString `db:"location_id" json:"locationID"`
}

func (q *Queries) CreateTokenAuthorization(ctx context.Context, arg CreateTokenAuthorizationParams) (TokenAuthorization, error) {
	row := q.db.QueryRowContext(ctx, createTokenAuthorization,
		arg.TokenID,
		arg.AuthorizationID,
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
	)
	return i, err
}

const getTokenAuthorizationByAuthorizationID = `-- name: GetTokenAuthorizationByAuthorizationID :one
SELECT id, token_id, authorization_id, country_code, party_id, location_id FROM token_authorizations
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
	)
	return i, err
}

const updateTariffByAuthorizationID = `-- name: UpdateTariffByAuthorizationID :one
UPDATE token_authorizations SET (
    country_code,
    party_id
  ) = ($2, $3)
  WHERE authorization_id = $1
  RETURNING id, token_id, authorization_id, country_code, party_id, location_id
`

type UpdateTariffByAuthorizationIDParams struct {
	AuthorizationID string         `db:"authorization_id" json:"authorizationID"`
	CountryCode     sql.NullString `db:"country_code" json:"countryCode"`
	PartyID         sql.NullString `db:"party_id" json:"partyID"`
}

func (q *Queries) UpdateTariffByAuthorizationID(ctx context.Context, arg UpdateTariffByAuthorizationIDParams) (TokenAuthorization, error) {
	row := q.db.QueryRowContext(ctx, updateTariffByAuthorizationID, arg.AuthorizationID, arg.CountryCode, arg.PartyID)
	var i TokenAuthorization
	err := row.Scan(
		&i.ID,
		&i.TokenID,
		&i.AuthorizationID,
		&i.CountryCode,
		&i.PartyID,
		&i.LocationID,
	)
	return i, err
}
