// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: invoice_request.sql

package db

import (
	"context"
	"database/sql"
)

const createInvoiceRequest = `-- name: CreateInvoiceRequest :one
INSERT INTO invoice_requests (
    user_id,
    promotion_id,
    session_id,
    release_date,
    currency,
    memo,
    price_fiat,
    price_msat,
    commission_fiat,
    commission_msat,
    tax_fiat,
    tax_msat,
    total_fiat,
    total_msat,
    is_settled, 
    payment_request
  ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16)
  RETURNING id, user_id, promotion_id, total_msat, is_settled, payment_request, currency, memo, total_fiat, price_fiat, price_msat, commission_fiat, commission_msat, tax_fiat, tax_msat, release_date, session_id
`

type CreateInvoiceRequestParams struct {
	UserID         int64           `db:"user_id" json:"userId"`
	PromotionID    int64           `db:"promotion_id" json:"promotionId"`
	SessionID      sql.NullInt64   `db:"session_id" json:"sessionId"`
	ReleaseDate    sql.NullTime    `db:"release_date" json:"releaseDate"`
	Currency       string          `db:"currency" json:"currency"`
	Memo           string          `db:"memo" json:"memo"`
	PriceFiat      sql.NullFloat64 `db:"price_fiat" json:"priceFiat"`
	PriceMsat      sql.NullInt64   `db:"price_msat" json:"priceMsat"`
	CommissionFiat sql.NullFloat64 `db:"commission_fiat" json:"commissionFiat"`
	CommissionMsat sql.NullInt64   `db:"commission_msat" json:"commissionMsat"`
	TaxFiat        sql.NullFloat64 `db:"tax_fiat" json:"taxFiat"`
	TaxMsat        sql.NullInt64   `db:"tax_msat" json:"taxMsat"`
	TotalFiat      float64         `db:"total_fiat" json:"totalFiat"`
	TotalMsat      int64           `db:"total_msat" json:"totalMsat"`
	IsSettled      bool            `db:"is_settled" json:"isSettled"`
	PaymentRequest sql.NullString  `db:"payment_request" json:"paymentRequest"`
}

func (q *Queries) CreateInvoiceRequest(ctx context.Context, arg CreateInvoiceRequestParams) (InvoiceRequest, error) {
	row := q.db.QueryRowContext(ctx, createInvoiceRequest,
		arg.UserID,
		arg.PromotionID,
		arg.SessionID,
		arg.ReleaseDate,
		arg.Currency,
		arg.Memo,
		arg.PriceFiat,
		arg.PriceMsat,
		arg.CommissionFiat,
		arg.CommissionMsat,
		arg.TaxFiat,
		arg.TaxMsat,
		arg.TotalFiat,
		arg.TotalMsat,
		arg.IsSettled,
		arg.PaymentRequest,
	)
	var i InvoiceRequest
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.PromotionID,
		&i.TotalMsat,
		&i.IsSettled,
		&i.PaymentRequest,
		&i.Currency,
		&i.Memo,
		&i.TotalFiat,
		&i.PriceFiat,
		&i.PriceMsat,
		&i.CommissionFiat,
		&i.CommissionMsat,
		&i.TaxFiat,
		&i.TaxMsat,
		&i.ReleaseDate,
		&i.SessionID,
	)
	return i, err
}

const deleteInvoiceRequest = `-- name: DeleteInvoiceRequest :exec
DELETE FROM invoice_requests
  WHERE id = $1
`

func (q *Queries) DeleteInvoiceRequest(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteInvoiceRequest, id)
	return err
}

const getInvoiceRequest = `-- name: GetInvoiceRequest :one
SELECT id, user_id, promotion_id, total_msat, is_settled, payment_request, currency, memo, total_fiat, price_fiat, price_msat, commission_fiat, commission_msat, tax_fiat, tax_msat, release_date, session_id FROM invoice_requests
  WHERE id = $1
`

func (q *Queries) GetInvoiceRequest(ctx context.Context, id int64) (InvoiceRequest, error) {
	row := q.db.QueryRowContext(ctx, getInvoiceRequest, id)
	var i InvoiceRequest
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.PromotionID,
		&i.TotalMsat,
		&i.IsSettled,
		&i.PaymentRequest,
		&i.Currency,
		&i.Memo,
		&i.TotalFiat,
		&i.PriceFiat,
		&i.PriceMsat,
		&i.CommissionFiat,
		&i.CommissionMsat,
		&i.TaxFiat,
		&i.TaxMsat,
		&i.ReleaseDate,
		&i.SessionID,
	)
	return i, err
}

const getUnsettledInvoiceRequest = `-- name: GetUnsettledInvoiceRequest :one
SELECT id, user_id, promotion_id, total_msat, is_settled, payment_request, currency, memo, total_fiat, price_fiat, price_msat, commission_fiat, commission_msat, tax_fiat, tax_msat, release_date, session_id FROM invoice_requests
  WHERE user_id = $1::BIGINT AND promotion_id = $2::BIGINT AND 
    ($3::TEXT = '' OR $3::TEXT = memo) AND
    NOT is_settled AND payment_request IS NULL
`

type GetUnsettledInvoiceRequestParams struct {
	UserID      int64  `db:"user_id" json:"userId"`
	PromotionID int64  `db:"promotion_id" json:"promotionId"`
	Memo        string `db:"memo" json:"memo"`
}

func (q *Queries) GetUnsettledInvoiceRequest(ctx context.Context, arg GetUnsettledInvoiceRequestParams) (InvoiceRequest, error) {
	row := q.db.QueryRowContext(ctx, getUnsettledInvoiceRequest, arg.UserID, arg.PromotionID, arg.Memo)
	var i InvoiceRequest
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.PromotionID,
		&i.TotalMsat,
		&i.IsSettled,
		&i.PaymentRequest,
		&i.Currency,
		&i.Memo,
		&i.TotalFiat,
		&i.PriceFiat,
		&i.PriceMsat,
		&i.CommissionFiat,
		&i.CommissionMsat,
		&i.TaxFiat,
		&i.TaxMsat,
		&i.ReleaseDate,
		&i.SessionID,
	)
	return i, err
}

const listInvoiceRequestsByPromotionCodeAndSessionID = `-- name: ListInvoiceRequestsByPromotionCodeAndSessionID :many
SELECT ir.id, ir.user_id, ir.promotion_id, ir.total_msat, ir.is_settled, ir.payment_request, ir.currency, ir.memo, ir.total_fiat, ir.price_fiat, ir.price_msat, ir.commission_fiat, ir.commission_msat, ir.tax_fiat, ir.tax_msat, ir.release_date, ir.session_id FROM invoice_requests ir
  LEFT JOIN promotions p ON p.id = ir.promotion_id
  WHERE ir.session_id = $1 AND p.code = $2 AND
    (ir.release_date IS NULL OR NOW() > ir.release_date)
  ORDER BY ir.id
`

type ListInvoiceRequestsByPromotionCodeAndSessionIDParams struct {
	SessionID sql.NullInt64 `db:"session_id" json:"sessionId"`
	Code      string        `db:"code" json:"code"`
}

func (q *Queries) ListInvoiceRequestsByPromotionCodeAndSessionID(ctx context.Context, arg ListInvoiceRequestsByPromotionCodeAndSessionIDParams) ([]InvoiceRequest, error) {
	rows, err := q.db.QueryContext(ctx, listInvoiceRequestsByPromotionCodeAndSessionID, arg.SessionID, arg.Code)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []InvoiceRequest
	for rows.Next() {
		var i InvoiceRequest
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.PromotionID,
			&i.TotalMsat,
			&i.IsSettled,
			&i.PaymentRequest,
			&i.Currency,
			&i.Memo,
			&i.TotalFiat,
			&i.PriceFiat,
			&i.PriceMsat,
			&i.CommissionFiat,
			&i.CommissionMsat,
			&i.TaxFiat,
			&i.TaxMsat,
			&i.ReleaseDate,
			&i.SessionID,
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

const listInvoiceRequestsBySessionID = `-- name: ListInvoiceRequestsBySessionID :many
SELECT id, user_id, promotion_id, total_msat, is_settled, payment_request, currency, memo, total_fiat, price_fiat, price_msat, commission_fiat, commission_msat, tax_fiat, tax_msat, release_date, session_id FROM invoice_requests
  WHERE session_id = $1 AND
    (release_date IS NULL OR NOW() > release_date)
  ORDER BY id
`

func (q *Queries) ListInvoiceRequestsBySessionID(ctx context.Context, sessionID sql.NullInt64) ([]InvoiceRequest, error) {
	rows, err := q.db.QueryContext(ctx, listInvoiceRequestsBySessionID, sessionID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []InvoiceRequest
	for rows.Next() {
		var i InvoiceRequest
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.PromotionID,
			&i.TotalMsat,
			&i.IsSettled,
			&i.PaymentRequest,
			&i.Currency,
			&i.Memo,
			&i.TotalFiat,
			&i.PriceFiat,
			&i.PriceMsat,
			&i.CommissionFiat,
			&i.CommissionMsat,
			&i.TaxFiat,
			&i.TaxMsat,
			&i.ReleaseDate,
			&i.SessionID,
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

const listUnsettledInvoiceRequests = `-- name: ListUnsettledInvoiceRequests :many
SELECT id, user_id, promotion_id, total_msat, is_settled, payment_request, currency, memo, total_fiat, price_fiat, price_msat, commission_fiat, commission_msat, tax_fiat, tax_msat, release_date, session_id FROM invoice_requests
  WHERE user_id = $1 AND is_settled = false AND payment_request IS NULL AND
    (release_date IS NULL OR NOW() > release_date)
  ORDER BY id
`

func (q *Queries) ListUnsettledInvoiceRequests(ctx context.Context, userID int64) ([]InvoiceRequest, error) {
	rows, err := q.db.QueryContext(ctx, listUnsettledInvoiceRequests, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []InvoiceRequest
	for rows.Next() {
		var i InvoiceRequest
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.PromotionID,
			&i.TotalMsat,
			&i.IsSettled,
			&i.PaymentRequest,
			&i.Currency,
			&i.Memo,
			&i.TotalFiat,
			&i.PriceFiat,
			&i.PriceMsat,
			&i.CommissionFiat,
			&i.CommissionMsat,
			&i.TaxFiat,
			&i.TaxMsat,
			&i.ReleaseDate,
			&i.SessionID,
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

const updateInvoiceRequest = `-- name: UpdateInvoiceRequest :one
UPDATE invoice_requests SET (
    price_fiat,
    price_msat,
    commission_fiat,
    commission_msat,
    tax_fiat,
    tax_msat,
    total_fiat,
    total_msat,
    is_settled,
    payment_request
  ) = ($2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
  WHERE id = $1
  RETURNING id, user_id, promotion_id, total_msat, is_settled, payment_request, currency, memo, total_fiat, price_fiat, price_msat, commission_fiat, commission_msat, tax_fiat, tax_msat, release_date, session_id
`

type UpdateInvoiceRequestParams struct {
	ID             int64           `db:"id" json:"id"`
	PriceFiat      sql.NullFloat64 `db:"price_fiat" json:"priceFiat"`
	PriceMsat      sql.NullInt64   `db:"price_msat" json:"priceMsat"`
	CommissionFiat sql.NullFloat64 `db:"commission_fiat" json:"commissionFiat"`
	CommissionMsat sql.NullInt64   `db:"commission_msat" json:"commissionMsat"`
	TaxFiat        sql.NullFloat64 `db:"tax_fiat" json:"taxFiat"`
	TaxMsat        sql.NullInt64   `db:"tax_msat" json:"taxMsat"`
	TotalFiat      float64         `db:"total_fiat" json:"totalFiat"`
	TotalMsat      int64           `db:"total_msat" json:"totalMsat"`
	IsSettled      bool            `db:"is_settled" json:"isSettled"`
	PaymentRequest sql.NullString  `db:"payment_request" json:"paymentRequest"`
}

func (q *Queries) UpdateInvoiceRequest(ctx context.Context, arg UpdateInvoiceRequestParams) (InvoiceRequest, error) {
	row := q.db.QueryRowContext(ctx, updateInvoiceRequest,
		arg.ID,
		arg.PriceFiat,
		arg.PriceMsat,
		arg.CommissionFiat,
		arg.CommissionMsat,
		arg.TaxFiat,
		arg.TaxMsat,
		arg.TotalFiat,
		arg.TotalMsat,
		arg.IsSettled,
		arg.PaymentRequest,
	)
	var i InvoiceRequest
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.PromotionID,
		&i.TotalMsat,
		&i.IsSettled,
		&i.PaymentRequest,
		&i.Currency,
		&i.Memo,
		&i.TotalFiat,
		&i.PriceFiat,
		&i.PriceMsat,
		&i.CommissionFiat,
		&i.CommissionMsat,
		&i.TaxFiat,
		&i.TaxMsat,
		&i.ReleaseDate,
		&i.SessionID,
	)
	return i, err
}
