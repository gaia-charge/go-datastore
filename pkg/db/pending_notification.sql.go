// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: pending_notification.sql

package db

import (
	"context"
	"database/sql"
	"time"

	"github.com/lib/pq"
)

const createPendingNotification = `-- name: CreatePendingNotification :one
INSERT INTO pending_notifications (
    user_id,
    node_id,
    invoice_request_id,
    device_token,
    type,
    send_date
  ) VALUES ($1, $2, $3, $4, $5, $6)
  RETURNING id, user_id, node_id, invoice_request_id, device_token, type, send_date
`

type CreatePendingNotificationParams struct {
	UserID           int64          `db:"user_id" json:"userId"`
	NodeID           int64          `db:"node_id" json:"nodeId"`
	InvoiceRequestID sql.NullInt64  `db:"invoice_request_id" json:"invoiceRequestId"`
	DeviceToken      sql.NullString `db:"device_token" json:"deviceToken"`
	Type             string         `db:"type" json:"type"`
	SendDate         time.Time      `db:"send_date" json:"sendDate"`
}

func (q *Queries) CreatePendingNotification(ctx context.Context, arg CreatePendingNotificationParams) (PendingNotification, error) {
	row := q.db.QueryRowContext(ctx, createPendingNotification,
		arg.UserID,
		arg.NodeID,
		arg.InvoiceRequestID,
		arg.DeviceToken,
		arg.Type,
		arg.SendDate,
	)
	var i PendingNotification
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.NodeID,
		&i.InvoiceRequestID,
		&i.DeviceToken,
		&i.Type,
		&i.SendDate,
	)
	return i, err
}

const deletePendingNotification = `-- name: DeletePendingNotification :exec
DELETE FROM pending_notifications
  WHERE id = $1
`

func (q *Queries) DeletePendingNotification(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deletePendingNotification, id)
	return err
}

const deletePendingNotificationByInvoiceRequest = `-- name: DeletePendingNotificationByInvoiceRequest :exec
DELETE FROM pending_notifications
  WHERE invoice_request_id = $1
  RETURNING id, user_id, node_id, invoice_request_id, device_token, type, send_date
`

func (q *Queries) DeletePendingNotificationByInvoiceRequest(ctx context.Context, invoiceRequestID sql.NullInt64) error {
	_, err := q.db.ExecContext(ctx, deletePendingNotificationByInvoiceRequest, invoiceRequestID)
	return err
}

const deletePendingNotifications = `-- name: DeletePendingNotifications :exec
DELETE FROM pending_notifications
  WHERE id = ANY($1::BIGINT[])
`

func (q *Queries) DeletePendingNotifications(ctx context.Context, ids []int64) error {
	_, err := q.db.ExecContext(ctx, deletePendingNotifications, pq.Array(ids))
	return err
}

const listPendingNotifications = `-- name: ListPendingNotifications :many
SELECT id, user_id, node_id, invoice_request_id, device_token, type, send_date FROM pending_notifications
  WHERE node_id = $1 AND send_date < NOW()
  ORDER BY id 
  LIMIT 1000
`

func (q *Queries) ListPendingNotifications(ctx context.Context, nodeID int64) ([]PendingNotification, error) {
	rows, err := q.db.QueryContext(ctx, listPendingNotifications, nodeID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []PendingNotification
	for rows.Next() {
		var i PendingNotification
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.NodeID,
			&i.InvoiceRequestID,
			&i.DeviceToken,
			&i.Type,
			&i.SendDate,
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

const updatePendingNotifications = `-- name: UpdatePendingNotifications :exec
UPDATE pending_notifications SET send_date = $1::TIMESTAMPTZ
  WHERE id = ANY($2::BIGINT[])
`

type UpdatePendingNotificationsParams struct {
	SendDate time.Time `db:"send_date" json:"sendDate"`
	Ids      []int64   `db:"ids" json:"ids"`
}

func (q *Queries) UpdatePendingNotifications(ctx context.Context, arg UpdatePendingNotificationsParams) error {
	_, err := q.db.ExecContext(ctx, updatePendingNotifications, arg.SendDate, pq.Array(arg.Ids))
	return err
}

const updatePendingNotificationsByUser = `-- name: UpdatePendingNotificationsByUser :exec
UPDATE pending_notifications SET device_token = $2
  WHERE user_id = $1
`

type UpdatePendingNotificationsByUserParams struct {
	UserID      int64          `db:"user_id" json:"userId"`
	DeviceToken sql.NullString `db:"device_token" json:"deviceToken"`
}

func (q *Queries) UpdatePendingNotificationsByUser(ctx context.Context, arg UpdatePendingNotificationsByUserParams) error {
	_, err := q.db.ExecContext(ctx, updatePendingNotificationsByUser, arg.UserID, arg.DeviceToken)
	return err
}
