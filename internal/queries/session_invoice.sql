-- name: CreateSessionInvoice :one
INSERT INTO session_invoices (
    session_id,
    user_id,
    currency,
    currency_rate,
    currency_rate_msat,
    price_fiat,
    price_msat,
    commission_fiat,
    commission_msat,
    tax_fiat,
    tax_msat,
    total_fiat,
    total_msat,
    payment_request,
    signature,
    is_settled,
    is_expired,
    estimated_energy,
    estimated_time,
    metered_energy,
    metered_time,
    last_updated
  ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22)
  RETURNING *;

-- name: GetSessionInvoice :one
SELECT * FROM session_invoices
  WHERE id = $1;

-- name: GetSessionInvoiceByPaymentRequest :one
SELECT * FROM session_invoices
  WHERE payment_request = $1;

-- name: GetUnsettledSessionInvoiceBySession :one
SELECT * FROM session_invoices
  WHERE session_id = $1 AND is_settled = false;

-- name: ListSessionInvoices :many
SELECT * FROM session_invoices
  WHERE is_expired = $1 AND is_settled = $2
  ORDER BY id;

-- name: ListSessionInvoicesByNodeID :many
SELECT si.* FROM session_invoices si
  INNER JOIN sessions s ON s.id = si.session_id
  INNER JOIN users u ON u.id = s.user_id
  WHERE u.node_id = $1 AND si.is_settled = $2 AND si.is_expired = $3
  ORDER BY si.id;

-- name: ListSessionInvoicesBySessionID :many
SELECT * FROM session_invoices
  WHERE session_id = $1
  ORDER BY id;

-- name: ListSessionInvoicesByUserID :many
SELECT si.* FROM session_invoices si
  INNER JOIN sessions s ON s.id = si.session_id
  INNER JOIN users u ON u.id = s.user_id
  WHERE u.id = $1 AND si.is_settled = $2 AND si.is_expired = $3
  ORDER BY si.id;

-- name: UpdateSessionInvoice :one
UPDATE session_invoices SET (
    currency,
    currency_rate,
    currency_rate_msat,
    price_fiat,
    price_msat,
    commission_fiat,
    commission_msat,
    tax_fiat,
    tax_msat,
    total_fiat,
    total_msat,
    payment_request,
    signature,
    is_settled,
    is_expired,
    estimated_energy,
    estimated_time,
    metered_energy,
    metered_time,
    last_updated
  ) = ($2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21)
  WHERE id = $1
  RETURNING *;