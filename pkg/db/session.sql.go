// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: session.sql

package db

import (
	"context"
	"database/sql"
	"time"
)

const createSession = `-- name: CreateSession :one
INSERT INTO sessions (
    uid,
    credential_id,
    country_code,
    party_id,
    authorization_id,
    start_datetime,
    end_datetime,
    kwh,
    auth_id,
    auth_method,
    user_id,
    token_id,
    location_id,
    evse_id,
    connector_id,
    meter_id,
    currency,
    total_cost,
    status,
    last_updated
  ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20)
  RETURNING id, uid, credential_id, country_code, party_id, authorization_id, start_datetime, end_datetime, kwh, auth_id, auth_method, user_id, token_id, location_id, evse_id, connector_id, meter_id, currency, total_cost, status, last_updated, invoice_request_id, is_flagged, is_confirmed_started, is_confirmed_stopped
`

type CreateSessionParams struct {
	Uid             string            `db:"uid" json:"uid"`
	CredentialID    int64             `db:"credential_id" json:"credentialId"`
	CountryCode     sql.NullString    `db:"country_code" json:"countryCode"`
	PartyID         sql.NullString    `db:"party_id" json:"partyId"`
	AuthorizationID sql.NullString    `db:"authorization_id" json:"authorizationId"`
	StartDatetime   time.Time         `db:"start_datetime" json:"startDatetime"`
	EndDatetime     sql.NullTime      `db:"end_datetime" json:"endDatetime"`
	Kwh             float64           `db:"kwh" json:"kwh"`
	AuthID          string            `db:"auth_id" json:"authId"`
	AuthMethod      AuthMethodType    `db:"auth_method" json:"authMethod"`
	UserID          int64             `db:"user_id" json:"userId"`
	TokenID         int64             `db:"token_id" json:"tokenId"`
	LocationID      int64             `db:"location_id" json:"locationId"`
	EvseID          int64             `db:"evse_id" json:"evseId"`
	ConnectorID     int64             `db:"connector_id" json:"connectorId"`
	MeterID         sql.NullString    `db:"meter_id" json:"meterId"`
	Currency        string            `db:"currency" json:"currency"`
	TotalCost       sql.NullFloat64   `db:"total_cost" json:"totalCost"`
	Status          SessionStatusType `db:"status" json:"status"`
	LastUpdated     time.Time         `db:"last_updated" json:"lastUpdated"`
}

func (q *Queries) CreateSession(ctx context.Context, arg CreateSessionParams) (Session, error) {
	row := q.db.QueryRowContext(ctx, createSession,
		arg.Uid,
		arg.CredentialID,
		arg.CountryCode,
		arg.PartyID,
		arg.AuthorizationID,
		arg.StartDatetime,
		arg.EndDatetime,
		arg.Kwh,
		arg.AuthID,
		arg.AuthMethod,
		arg.UserID,
		arg.TokenID,
		arg.LocationID,
		arg.EvseID,
		arg.ConnectorID,
		arg.MeterID,
		arg.Currency,
		arg.TotalCost,
		arg.Status,
		arg.LastUpdated,
	)
	var i Session
	err := row.Scan(
		&i.ID,
		&i.Uid,
		&i.CredentialID,
		&i.CountryCode,
		&i.PartyID,
		&i.AuthorizationID,
		&i.StartDatetime,
		&i.EndDatetime,
		&i.Kwh,
		&i.AuthID,
		&i.AuthMethod,
		&i.UserID,
		&i.TokenID,
		&i.LocationID,
		&i.EvseID,
		&i.ConnectorID,
		&i.MeterID,
		&i.Currency,
		&i.TotalCost,
		&i.Status,
		&i.LastUpdated,
		&i.InvoiceRequestID,
		&i.IsFlagged,
		&i.IsConfirmedStarted,
		&i.IsConfirmedStopped,
	)
	return i, err
}

const getSession = `-- name: GetSession :one
SELECT id, uid, credential_id, country_code, party_id, authorization_id, start_datetime, end_datetime, kwh, auth_id, auth_method, user_id, token_id, location_id, evse_id, connector_id, meter_id, currency, total_cost, status, last_updated, invoice_request_id, is_flagged, is_confirmed_started, is_confirmed_stopped FROM sessions
  WHERE id = $1
`

func (q *Queries) GetSession(ctx context.Context, id int64) (Session, error) {
	row := q.db.QueryRowContext(ctx, getSession, id)
	var i Session
	err := row.Scan(
		&i.ID,
		&i.Uid,
		&i.CredentialID,
		&i.CountryCode,
		&i.PartyID,
		&i.AuthorizationID,
		&i.StartDatetime,
		&i.EndDatetime,
		&i.Kwh,
		&i.AuthID,
		&i.AuthMethod,
		&i.UserID,
		&i.TokenID,
		&i.LocationID,
		&i.EvseID,
		&i.ConnectorID,
		&i.MeterID,
		&i.Currency,
		&i.TotalCost,
		&i.Status,
		&i.LastUpdated,
		&i.InvoiceRequestID,
		&i.IsFlagged,
		&i.IsConfirmedStarted,
		&i.IsConfirmedStopped,
	)
	return i, err
}

const getSessionByAuthorizationID = `-- name: GetSessionByAuthorizationID :one
SELECT id, uid, credential_id, country_code, party_id, authorization_id, start_datetime, end_datetime, kwh, auth_id, auth_method, user_id, token_id, location_id, evse_id, connector_id, meter_id, currency, total_cost, status, last_updated, invoice_request_id, is_flagged, is_confirmed_started, is_confirmed_stopped FROM sessions
  WHERE authorization_id = $1::TEXT
  LIMIT 1
`

func (q *Queries) GetSessionByAuthorizationID(ctx context.Context, authorizationID string) (Session, error) {
	row := q.db.QueryRowContext(ctx, getSessionByAuthorizationID, authorizationID)
	var i Session
	err := row.Scan(
		&i.ID,
		&i.Uid,
		&i.CredentialID,
		&i.CountryCode,
		&i.PartyID,
		&i.AuthorizationID,
		&i.StartDatetime,
		&i.EndDatetime,
		&i.Kwh,
		&i.AuthID,
		&i.AuthMethod,
		&i.UserID,
		&i.TokenID,
		&i.LocationID,
		&i.EvseID,
		&i.ConnectorID,
		&i.MeterID,
		&i.Currency,
		&i.TotalCost,
		&i.Status,
		&i.LastUpdated,
		&i.InvoiceRequestID,
		&i.IsFlagged,
		&i.IsConfirmedStarted,
		&i.IsConfirmedStopped,
	)
	return i, err
}

const getSessionByLastUpdated = `-- name: GetSessionByLastUpdated :one
SELECT id, uid, credential_id, country_code, party_id, authorization_id, start_datetime, end_datetime, kwh, auth_id, auth_method, user_id, token_id, location_id, evse_id, connector_id, meter_id, currency, total_cost, status, last_updated, invoice_request_id, is_flagged, is_confirmed_started, is_confirmed_stopped FROM sessions
  WHERE ($1::BIGINT = -1 OR $1::BIGINT = credential_id) AND
    ($2::TEXT = '' OR $2::TEXT = country_code) AND
    ($3::TEXT = '' OR $3::TEXT = party_id)
  ORDER BY last_updated DESC
  LIMIT 1
`

type GetSessionByLastUpdatedParams struct {
	CredentialID int64  `db:"credential_id" json:"credentialId"`
	CountryCode  string `db:"country_code" json:"countryCode"`
	PartyID      string `db:"party_id" json:"partyId"`
}

func (q *Queries) GetSessionByLastUpdated(ctx context.Context, arg GetSessionByLastUpdatedParams) (Session, error) {
	row := q.db.QueryRowContext(ctx, getSessionByLastUpdated, arg.CredentialID, arg.CountryCode, arg.PartyID)
	var i Session
	err := row.Scan(
		&i.ID,
		&i.Uid,
		&i.CredentialID,
		&i.CountryCode,
		&i.PartyID,
		&i.AuthorizationID,
		&i.StartDatetime,
		&i.EndDatetime,
		&i.Kwh,
		&i.AuthID,
		&i.AuthMethod,
		&i.UserID,
		&i.TokenID,
		&i.LocationID,
		&i.EvseID,
		&i.ConnectorID,
		&i.MeterID,
		&i.Currency,
		&i.TotalCost,
		&i.Status,
		&i.LastUpdated,
		&i.InvoiceRequestID,
		&i.IsFlagged,
		&i.IsConfirmedStarted,
		&i.IsConfirmedStopped,
	)
	return i, err
}

const getSessionByUid = `-- name: GetSessionByUid :one
SELECT id, uid, credential_id, country_code, party_id, authorization_id, start_datetime, end_datetime, kwh, auth_id, auth_method, user_id, token_id, location_id, evse_id, connector_id, meter_id, currency, total_cost, status, last_updated, invoice_request_id, is_flagged, is_confirmed_started, is_confirmed_stopped FROM sessions
  WHERE uid = $1
`

func (q *Queries) GetSessionByUid(ctx context.Context, uid string) (Session, error) {
	row := q.db.QueryRowContext(ctx, getSessionByUid, uid)
	var i Session
	err := row.Scan(
		&i.ID,
		&i.Uid,
		&i.CredentialID,
		&i.CountryCode,
		&i.PartyID,
		&i.AuthorizationID,
		&i.StartDatetime,
		&i.EndDatetime,
		&i.Kwh,
		&i.AuthID,
		&i.AuthMethod,
		&i.UserID,
		&i.TokenID,
		&i.LocationID,
		&i.EvseID,
		&i.ConnectorID,
		&i.MeterID,
		&i.Currency,
		&i.TotalCost,
		&i.Status,
		&i.LastUpdated,
		&i.InvoiceRequestID,
		&i.IsFlagged,
		&i.IsConfirmedStarted,
		&i.IsConfirmedStopped,
	)
	return i, err
}

const listInProgressSessionsByNodeID = `-- name: ListInProgressSessionsByNodeID :many
SELECT s.id, s.uid, s.credential_id, s.country_code, s.party_id, s.authorization_id, s.start_datetime, s.end_datetime, s.kwh, s.auth_id, s.auth_method, s.user_id, s.token_id, s.location_id, s.evse_id, s.connector_id, s.meter_id, s.currency, s.total_cost, s.status, s.last_updated, s.invoice_request_id, s.is_flagged, s.is_confirmed_started, s.is_confirmed_stopped FROM sessions s
  INNER JOIN users u ON u.id = s.user_id
  WHERE u.node_id = $1 AND s.status IN ('PENDING', 'ACTIVE')
  ORDER BY s.id
`

func (q *Queries) ListInProgressSessionsByNodeID(ctx context.Context, nodeID sql.NullInt64) ([]Session, error) {
	rows, err := q.db.QueryContext(ctx, listInProgressSessionsByNodeID, nodeID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Session
	for rows.Next() {
		var i Session
		if err := rows.Scan(
			&i.ID,
			&i.Uid,
			&i.CredentialID,
			&i.CountryCode,
			&i.PartyID,
			&i.AuthorizationID,
			&i.StartDatetime,
			&i.EndDatetime,
			&i.Kwh,
			&i.AuthID,
			&i.AuthMethod,
			&i.UserID,
			&i.TokenID,
			&i.LocationID,
			&i.EvseID,
			&i.ConnectorID,
			&i.MeterID,
			&i.Currency,
			&i.TotalCost,
			&i.Status,
			&i.LastUpdated,
			&i.InvoiceRequestID,
			&i.IsFlagged,
			&i.IsConfirmedStarted,
			&i.IsConfirmedStopped,
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

const listInProgressSessionsByUserID = `-- name: ListInProgressSessionsByUserID :many
SELECT id, uid, credential_id, country_code, party_id, authorization_id, start_datetime, end_datetime, kwh, auth_id, auth_method, user_id, token_id, location_id, evse_id, connector_id, meter_id, currency, total_cost, status, last_updated, invoice_request_id, is_flagged, is_confirmed_started, is_confirmed_stopped FROM sessions
  WHERE user_id = $1 AND status IN ('PENDING', 'ACTIVE', 'ENDING')
  ORDER BY id DESC
`

func (q *Queries) ListInProgressSessionsByUserID(ctx context.Context, userID int64) ([]Session, error) {
	rows, err := q.db.QueryContext(ctx, listInProgressSessionsByUserID, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Session
	for rows.Next() {
		var i Session
		if err := rows.Scan(
			&i.ID,
			&i.Uid,
			&i.CredentialID,
			&i.CountryCode,
			&i.PartyID,
			&i.AuthorizationID,
			&i.StartDatetime,
			&i.EndDatetime,
			&i.Kwh,
			&i.AuthID,
			&i.AuthMethod,
			&i.UserID,
			&i.TokenID,
			&i.LocationID,
			&i.EvseID,
			&i.ConnectorID,
			&i.MeterID,
			&i.Currency,
			&i.TotalCost,
			&i.Status,
			&i.LastUpdated,
			&i.InvoiceRequestID,
			&i.IsFlagged,
			&i.IsConfirmedStarted,
			&i.IsConfirmedStopped,
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

const listInvoicedSessionsByUserID = `-- name: ListInvoicedSessionsByUserID :many
SELECT id, uid, credential_id, country_code, party_id, authorization_id, start_datetime, end_datetime, kwh, auth_id, auth_method, user_id, token_id, location_id, evse_id, connector_id, meter_id, currency, total_cost, status, last_updated, invoice_request_id, is_flagged, is_confirmed_started, is_confirmed_stopped FROM sessions
  WHERE user_id = $1 AND status = 'INVOICED'
  ORDER BY id DESC
`

func (q *Queries) ListInvoicedSessionsByUserID(ctx context.Context, userID int64) ([]Session, error) {
	rows, err := q.db.QueryContext(ctx, listInvoicedSessionsByUserID, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Session
	for rows.Next() {
		var i Session
		if err := rows.Scan(
			&i.ID,
			&i.Uid,
			&i.CredentialID,
			&i.CountryCode,
			&i.PartyID,
			&i.AuthorizationID,
			&i.StartDatetime,
			&i.EndDatetime,
			&i.Kwh,
			&i.AuthID,
			&i.AuthMethod,
			&i.UserID,
			&i.TokenID,
			&i.LocationID,
			&i.EvseID,
			&i.ConnectorID,
			&i.MeterID,
			&i.Currency,
			&i.TotalCost,
			&i.Status,
			&i.LastUpdated,
			&i.InvoiceRequestID,
			&i.IsFlagged,
			&i.IsConfirmedStarted,
			&i.IsConfirmedStopped,
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

const updateSessionByUid = `-- name: UpdateSessionByUid :one
UPDATE sessions SET (
    authorization_id,
    start_datetime,
    end_datetime,
    kwh,
    auth_method,
    meter_id,
    currency,
    total_cost,
    status,
    invoice_request_id,
    is_confirmed_started,
    is_confirmed_stopped,
    is_flagged,
    last_updated
  ) = ($2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15)
  WHERE uid = $1
  RETURNING id, uid, credential_id, country_code, party_id, authorization_id, start_datetime, end_datetime, kwh, auth_id, auth_method, user_id, token_id, location_id, evse_id, connector_id, meter_id, currency, total_cost, status, last_updated, invoice_request_id, is_flagged, is_confirmed_started, is_confirmed_stopped
`

type UpdateSessionByUidParams struct {
	Uid                string            `db:"uid" json:"uid"`
	AuthorizationID    sql.NullString    `db:"authorization_id" json:"authorizationId"`
	StartDatetime      time.Time         `db:"start_datetime" json:"startDatetime"`
	EndDatetime        sql.NullTime      `db:"end_datetime" json:"endDatetime"`
	Kwh                float64           `db:"kwh" json:"kwh"`
	AuthMethod         AuthMethodType    `db:"auth_method" json:"authMethod"`
	MeterID            sql.NullString    `db:"meter_id" json:"meterId"`
	Currency           string            `db:"currency" json:"currency"`
	TotalCost          sql.NullFloat64   `db:"total_cost" json:"totalCost"`
	Status             SessionStatusType `db:"status" json:"status"`
	InvoiceRequestID   sql.NullInt64     `db:"invoice_request_id" json:"invoiceRequestId"`
	IsConfirmedStarted bool              `db:"is_confirmed_started" json:"isConfirmedStarted"`
	IsConfirmedStopped bool              `db:"is_confirmed_stopped" json:"isConfirmedStopped"`
	IsFlagged          bool              `db:"is_flagged" json:"isFlagged"`
	LastUpdated        time.Time         `db:"last_updated" json:"lastUpdated"`
}

func (q *Queries) UpdateSessionByUid(ctx context.Context, arg UpdateSessionByUidParams) (Session, error) {
	row := q.db.QueryRowContext(ctx, updateSessionByUid,
		arg.Uid,
		arg.AuthorizationID,
		arg.StartDatetime,
		arg.EndDatetime,
		arg.Kwh,
		arg.AuthMethod,
		arg.MeterID,
		arg.Currency,
		arg.TotalCost,
		arg.Status,
		arg.InvoiceRequestID,
		arg.IsConfirmedStarted,
		arg.IsConfirmedStopped,
		arg.IsFlagged,
		arg.LastUpdated,
	)
	var i Session
	err := row.Scan(
		&i.ID,
		&i.Uid,
		&i.CredentialID,
		&i.CountryCode,
		&i.PartyID,
		&i.AuthorizationID,
		&i.StartDatetime,
		&i.EndDatetime,
		&i.Kwh,
		&i.AuthID,
		&i.AuthMethod,
		&i.UserID,
		&i.TokenID,
		&i.LocationID,
		&i.EvseID,
		&i.ConnectorID,
		&i.MeterID,
		&i.Currency,
		&i.TotalCost,
		&i.Status,
		&i.LastUpdated,
		&i.InvoiceRequestID,
		&i.IsFlagged,
		&i.IsConfirmedStarted,
		&i.IsConfirmedStopped,
	)
	return i, err
}

const updateSessionIsFlaggedByUid = `-- name: UpdateSessionIsFlaggedByUid :exec
UPDATE sessions SET is_flagged = $2
  WHERE uid = $1
`

type UpdateSessionIsFlaggedByUidParams struct {
	Uid       string `db:"uid" json:"uid"`
	IsFlagged bool   `db:"is_flagged" json:"isFlagged"`
}

func (q *Queries) UpdateSessionIsFlaggedByUid(ctx context.Context, arg UpdateSessionIsFlaggedByUidParams) error {
	_, err := q.db.ExecContext(ctx, updateSessionIsFlaggedByUid, arg.Uid, arg.IsFlagged)
	return err
}
