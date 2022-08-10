// Code generated by sqlc. DO NOT EDIT.
// source: routing_event.sql

package db

import (
	"context"
	"database/sql"
	"time"
)

const createRoutingEvent = `-- name: CreateRoutingEvent :one
INSERT INTO routing_events (
    node_id,
    event_type,
    event_status,
    currency,
    currency_rate,
    currency_rate_msat,
    incoming_chan_id,
    incoming_htlc_id,
    incoming_fiat,
    incoming_msat,
    outgoing_chan_id,
    outgoing_htlc_id,
    outgoing_fiat,
    outgoing_msat,
    fee_fiat,
    fee_msat,
    wire_failure,
    failure_detail,
    failure_string,
    last_updated
  ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20)
  RETURNING id, node_id, event_type, currency, currency_rate, currency_rate_msat, incoming_chan_id, incoming_htlc_id, incoming_fiat, incoming_msat, outgoing_chan_id, outgoing_htlc_id, outgoing_fiat, outgoing_msat, fee_fiat, fee_msat, last_updated, event_status, wire_failure, failure_detail, failure_string
`

type CreateRoutingEventParams struct {
	NodeID           int64              `db:"node_id" json:"nodeID"`
	EventType        RoutingEventType   `db:"event_type" json:"eventType"`
	EventStatus      RoutingEventStatus `db:"event_status" json:"eventStatus"`
	Currency         string             `db:"currency" json:"currency"`
	CurrencyRate     int64              `db:"currency_rate" json:"currencyRate"`
	CurrencyRateMsat int64              `db:"currency_rate_msat" json:"currencyRateMsat"`
	IncomingChanID   int64              `db:"incoming_chan_id" json:"incomingChanID"`
	IncomingHtlcID   int64              `db:"incoming_htlc_id" json:"incomingHtlcID"`
	IncomingFiat     float64            `db:"incoming_fiat" json:"incomingFiat"`
	IncomingMsat     int64              `db:"incoming_msat" json:"incomingMsat"`
	OutgoingChanID   int64              `db:"outgoing_chan_id" json:"outgoingChanID"`
	OutgoingHtlcID   int64              `db:"outgoing_htlc_id" json:"outgoingHtlcID"`
	OutgoingFiat     float64            `db:"outgoing_fiat" json:"outgoingFiat"`
	OutgoingMsat     int64              `db:"outgoing_msat" json:"outgoingMsat"`
	FeeFiat          float64            `db:"fee_fiat" json:"feeFiat"`
	FeeMsat          int64              `db:"fee_msat" json:"feeMsat"`
	WireFailure      sql.NullInt32      `db:"wire_failure" json:"wireFailure"`
	FailureDetail    sql.NullInt32      `db:"failure_detail" json:"failureDetail"`
	FailureString    sql.NullString     `db:"failure_string" json:"failureString"`
	LastUpdated      time.Time          `db:"last_updated" json:"lastUpdated"`
}

func (q *Queries) CreateRoutingEvent(ctx context.Context, arg CreateRoutingEventParams) (RoutingEvent, error) {
	row := q.db.QueryRowContext(ctx, createRoutingEvent,
		arg.NodeID,
		arg.EventType,
		arg.EventStatus,
		arg.Currency,
		arg.CurrencyRate,
		arg.CurrencyRateMsat,
		arg.IncomingChanID,
		arg.IncomingHtlcID,
		arg.IncomingFiat,
		arg.IncomingMsat,
		arg.OutgoingChanID,
		arg.OutgoingHtlcID,
		arg.OutgoingFiat,
		arg.OutgoingMsat,
		arg.FeeFiat,
		arg.FeeMsat,
		arg.WireFailure,
		arg.FailureDetail,
		arg.FailureString,
		arg.LastUpdated,
	)
	var i RoutingEvent
	err := row.Scan(
		&i.ID,
		&i.NodeID,
		&i.EventType,
		&i.Currency,
		&i.CurrencyRate,
		&i.CurrencyRateMsat,
		&i.IncomingChanID,
		&i.IncomingHtlcID,
		&i.IncomingFiat,
		&i.IncomingMsat,
		&i.OutgoingChanID,
		&i.OutgoingHtlcID,
		&i.OutgoingFiat,
		&i.OutgoingMsat,
		&i.FeeFiat,
		&i.FeeMsat,
		&i.LastUpdated,
		&i.EventStatus,
		&i.WireFailure,
		&i.FailureDetail,
		&i.FailureString,
	)
	return i, err
}

const updateRoutingEvent = `-- name: UpdateRoutingEvent :one
UPDATE routing_events SET (
    event_status,
    wire_failure,
    failure_detail,
    failure_string,
    last_updated
  ) = ($5, $6, $7, $8, $9)
  WHERE incoming_chan_id = $1 AND incoming_htlc_id = $2 AND outgoing_chan_id = $3 AND outgoing_htlc_id = $4
  RETURNING id, node_id, event_type, currency, currency_rate, currency_rate_msat, incoming_chan_id, incoming_htlc_id, incoming_fiat, incoming_msat, outgoing_chan_id, outgoing_htlc_id, outgoing_fiat, outgoing_msat, fee_fiat, fee_msat, last_updated, event_status, wire_failure, failure_detail, failure_string
`

type UpdateRoutingEventParams struct {
	IncomingChanID int64              `db:"incoming_chan_id" json:"incomingChanID"`
	IncomingHtlcID int64              `db:"incoming_htlc_id" json:"incomingHtlcID"`
	OutgoingChanID int64              `db:"outgoing_chan_id" json:"outgoingChanID"`
	OutgoingHtlcID int64              `db:"outgoing_htlc_id" json:"outgoingHtlcID"`
	EventStatus    RoutingEventStatus `db:"event_status" json:"eventStatus"`
	WireFailure    sql.NullInt32      `db:"wire_failure" json:"wireFailure"`
	FailureDetail  sql.NullInt32      `db:"failure_detail" json:"failureDetail"`
	FailureString  sql.NullString     `db:"failure_string" json:"failureString"`
	LastUpdated    time.Time          `db:"last_updated" json:"lastUpdated"`
}

func (q *Queries) UpdateRoutingEvent(ctx context.Context, arg UpdateRoutingEventParams) (RoutingEvent, error) {
	row := q.db.QueryRowContext(ctx, updateRoutingEvent,
		arg.IncomingChanID,
		arg.IncomingHtlcID,
		arg.OutgoingChanID,
		arg.OutgoingHtlcID,
		arg.EventStatus,
		arg.WireFailure,
		arg.FailureDetail,
		arg.FailureString,
		arg.LastUpdated,
	)
	var i RoutingEvent
	err := row.Scan(
		&i.ID,
		&i.NodeID,
		&i.EventType,
		&i.Currency,
		&i.CurrencyRate,
		&i.CurrencyRateMsat,
		&i.IncomingChanID,
		&i.IncomingHtlcID,
		&i.IncomingFiat,
		&i.IncomingMsat,
		&i.OutgoingChanID,
		&i.OutgoingHtlcID,
		&i.OutgoingFiat,
		&i.OutgoingMsat,
		&i.FeeFiat,
		&i.FeeMsat,
		&i.LastUpdated,
		&i.EventStatus,
		&i.WireFailure,
		&i.FailureDetail,
		&i.FailureString,
	)
	return i, err
}
