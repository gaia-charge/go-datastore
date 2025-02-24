// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: channel_request.sql

package db

import (
	"context"
	"database/sql"
)

const createChannelRequest = `-- name: CreateChannelRequest :one
INSERT INTO channel_requests (
    user_id,
    node_id,
    status,
    pubkey, 
    payment_hash, 
    payment_addr,
    amount,
    amount_msat,
    settled_msat,
    pending_chan_id,
    scid,
    fee_base_msat,
    fee_proportional_millionths,
    cltv_expiry_delta
  ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)
  RETURNING id, user_id, status, pubkey, payment_hash, payment_addr, amount_msat, settled_msat, funding_tx_id_bytes, output_index, node_id, amount, funding_amount, pending_chan_id, scid, fee_base_msat, fee_proportional_millionths, cltv_expiry_delta, funding_tx_id
`

type CreateChannelRequestParams struct {
	UserID                    int64                `db:"user_id" json:"userId"`
	NodeID                    int64                `db:"node_id" json:"nodeId"`
	Status                    ChannelRequestStatus `db:"status" json:"status"`
	Pubkey                    string               `db:"pubkey" json:"pubkey"`
	PaymentHash               []byte               `db:"payment_hash" json:"paymentHash"`
	PaymentAddr               []byte               `db:"payment_addr" json:"paymentAddr"`
	Amount                    int64                `db:"amount" json:"amount"`
	AmountMsat                int64                `db:"amount_msat" json:"amountMsat"`
	SettledMsat               int64                `db:"settled_msat" json:"settledMsat"`
	PendingChanID             []byte               `db:"pending_chan_id" json:"pendingChanId"`
	Scid                      []byte               `db:"scid" json:"scid"`
	FeeBaseMsat               int64                `db:"fee_base_msat" json:"feeBaseMsat"`
	FeeProportionalMillionths int64                `db:"fee_proportional_millionths" json:"feeProportionalMillionths"`
	CltvExpiryDelta           int64                `db:"cltv_expiry_delta" json:"cltvExpiryDelta"`
}

func (q *Queries) CreateChannelRequest(ctx context.Context, arg CreateChannelRequestParams) (ChannelRequest, error) {
	row := q.db.QueryRowContext(ctx, createChannelRequest,
		arg.UserID,
		arg.NodeID,
		arg.Status,
		arg.Pubkey,
		arg.PaymentHash,
		arg.PaymentAddr,
		arg.Amount,
		arg.AmountMsat,
		arg.SettledMsat,
		arg.PendingChanID,
		arg.Scid,
		arg.FeeBaseMsat,
		arg.FeeProportionalMillionths,
		arg.CltvExpiryDelta,
	)
	var i ChannelRequest
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Status,
		&i.Pubkey,
		&i.PaymentHash,
		&i.PaymentAddr,
		&i.AmountMsat,
		&i.SettledMsat,
		&i.FundingTxIDBytes,
		&i.OutputIndex,
		&i.NodeID,
		&i.Amount,
		&i.FundingAmount,
		&i.PendingChanID,
		&i.Scid,
		&i.FeeBaseMsat,
		&i.FeeProportionalMillionths,
		&i.CltvExpiryDelta,
		&i.FundingTxID,
	)
	return i, err
}

const deleteChannelRequest = `-- name: DeleteChannelRequest :exec
DELETE FROM channel_requests
  WHERE id = $1
`

func (q *Queries) DeleteChannelRequest(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteChannelRequest, id)
	return err
}

const getChannelRequest = `-- name: GetChannelRequest :one
SELECT id, user_id, status, pubkey, payment_hash, payment_addr, amount_msat, settled_msat, funding_tx_id_bytes, output_index, node_id, amount, funding_amount, pending_chan_id, scid, fee_base_msat, fee_proportional_millionths, cltv_expiry_delta, funding_tx_id FROM channel_requests
  WHERE id = $1
`

func (q *Queries) GetChannelRequest(ctx context.Context, id int64) (ChannelRequest, error) {
	row := q.db.QueryRowContext(ctx, getChannelRequest, id)
	var i ChannelRequest
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Status,
		&i.Pubkey,
		&i.PaymentHash,
		&i.PaymentAddr,
		&i.AmountMsat,
		&i.SettledMsat,
		&i.FundingTxIDBytes,
		&i.OutputIndex,
		&i.NodeID,
		&i.Amount,
		&i.FundingAmount,
		&i.PendingChanID,
		&i.Scid,
		&i.FeeBaseMsat,
		&i.FeeProportionalMillionths,
		&i.CltvExpiryDelta,
		&i.FundingTxID,
	)
	return i, err
}

const getChannelRequestByChannelPoint = `-- name: GetChannelRequestByChannelPoint :one
SELECT id, user_id, status, pubkey, payment_hash, payment_addr, amount_msat, settled_msat, funding_tx_id_bytes, output_index, node_id, amount, funding_amount, pending_chan_id, scid, fee_base_msat, fee_proportional_millionths, cltv_expiry_delta, funding_tx_id FROM channel_requests
  WHERE funding_tx_id_bytes = $1 AND output_index = $2
`

type GetChannelRequestByChannelPointParams struct {
	FundingTxIDBytes []byte        `db:"funding_tx_id_bytes" json:"fundingTxIdBytes"`
	OutputIndex      sql.NullInt64 `db:"output_index" json:"outputIndex"`
}

func (q *Queries) GetChannelRequestByChannelPoint(ctx context.Context, arg GetChannelRequestByChannelPointParams) (ChannelRequest, error) {
	row := q.db.QueryRowContext(ctx, getChannelRequestByChannelPoint, arg.FundingTxIDBytes, arg.OutputIndex)
	var i ChannelRequest
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Status,
		&i.Pubkey,
		&i.PaymentHash,
		&i.PaymentAddr,
		&i.AmountMsat,
		&i.SettledMsat,
		&i.FundingTxIDBytes,
		&i.OutputIndex,
		&i.NodeID,
		&i.Amount,
		&i.FundingAmount,
		&i.PendingChanID,
		&i.Scid,
		&i.FeeBaseMsat,
		&i.FeeProportionalMillionths,
		&i.CltvExpiryDelta,
		&i.FundingTxID,
	)
	return i, err
}

const getChannelRequestByPaymentHash = `-- name: GetChannelRequestByPaymentHash :one
SELECT id, user_id, status, pubkey, payment_hash, payment_addr, amount_msat, settled_msat, funding_tx_id_bytes, output_index, node_id, amount, funding_amount, pending_chan_id, scid, fee_base_msat, fee_proportional_millionths, cltv_expiry_delta, funding_tx_id FROM channel_requests
  WHERE payment_hash = $1 OR digest('probing-01:' || payment_hash, 'sha256') = $1
`

func (q *Queries) GetChannelRequestByPaymentHash(ctx context.Context, paymentHash []byte) (ChannelRequest, error) {
	row := q.db.QueryRowContext(ctx, getChannelRequestByPaymentHash, paymentHash)
	var i ChannelRequest
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Status,
		&i.Pubkey,
		&i.PaymentHash,
		&i.PaymentAddr,
		&i.AmountMsat,
		&i.SettledMsat,
		&i.FundingTxIDBytes,
		&i.OutputIndex,
		&i.NodeID,
		&i.Amount,
		&i.FundingAmount,
		&i.PendingChanID,
		&i.Scid,
		&i.FeeBaseMsat,
		&i.FeeProportionalMillionths,
		&i.CltvExpiryDelta,
		&i.FundingTxID,
	)
	return i, err
}

const getChannelRequestByPendingChanId = `-- name: GetChannelRequestByPendingChanId :one
SELECT id, user_id, status, pubkey, payment_hash, payment_addr, amount_msat, settled_msat, funding_tx_id_bytes, output_index, node_id, amount, funding_amount, pending_chan_id, scid, fee_base_msat, fee_proportional_millionths, cltv_expiry_delta, funding_tx_id FROM channel_requests
  WHERE pending_chan_id = $1
`

func (q *Queries) GetChannelRequestByPendingChanId(ctx context.Context, pendingChanID []byte) (ChannelRequest, error) {
	row := q.db.QueryRowContext(ctx, getChannelRequestByPendingChanId, pendingChanID)
	var i ChannelRequest
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Status,
		&i.Pubkey,
		&i.PaymentHash,
		&i.PaymentAddr,
		&i.AmountMsat,
		&i.SettledMsat,
		&i.FundingTxIDBytes,
		&i.OutputIndex,
		&i.NodeID,
		&i.Amount,
		&i.FundingAmount,
		&i.PendingChanID,
		&i.Scid,
		&i.FeeBaseMsat,
		&i.FeeProportionalMillionths,
		&i.CltvExpiryDelta,
		&i.FundingTxID,
	)
	return i, err
}

const updateChannelRequest = `-- name: UpdateChannelRequest :one
UPDATE channel_requests SET (
    status,
    settled_msat,
    funding_amount
  ) = ($2, $3, $4)
  WHERE id = $1
  RETURNING id, user_id, status, pubkey, payment_hash, payment_addr, amount_msat, settled_msat, funding_tx_id_bytes, output_index, node_id, amount, funding_amount, pending_chan_id, scid, fee_base_msat, fee_proportional_millionths, cltv_expiry_delta, funding_tx_id
`

type UpdateChannelRequestParams struct {
	ID            int64                `db:"id" json:"id"`
	Status        ChannelRequestStatus `db:"status" json:"status"`
	SettledMsat   int64                `db:"settled_msat" json:"settledMsat"`
	FundingAmount sql.NullInt64        `db:"funding_amount" json:"fundingAmount"`
}

func (q *Queries) UpdateChannelRequest(ctx context.Context, arg UpdateChannelRequestParams) (ChannelRequest, error) {
	row := q.db.QueryRowContext(ctx, updateChannelRequest,
		arg.ID,
		arg.Status,
		arg.SettledMsat,
		arg.FundingAmount,
	)
	var i ChannelRequest
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Status,
		&i.Pubkey,
		&i.PaymentHash,
		&i.PaymentAddr,
		&i.AmountMsat,
		&i.SettledMsat,
		&i.FundingTxIDBytes,
		&i.OutputIndex,
		&i.NodeID,
		&i.Amount,
		&i.FundingAmount,
		&i.PendingChanID,
		&i.Scid,
		&i.FeeBaseMsat,
		&i.FeeProportionalMillionths,
		&i.CltvExpiryDelta,
		&i.FundingTxID,
	)
	return i, err
}

const updatePendingChannelRequestByPubkey = `-- name: UpdatePendingChannelRequestByPubkey :one
UPDATE channel_requests SET (
    funding_tx_id, 
    funding_tx_id_bytes, 
    output_index
  ) = ($3, $4, $5)
  WHERE status = 'OPENING_CHANNEL' AND pubkey = $1 AND funding_amount = $2
  RETURNING id, user_id, status, pubkey, payment_hash, payment_addr, amount_msat, settled_msat, funding_tx_id_bytes, output_index, node_id, amount, funding_amount, pending_chan_id, scid, fee_base_msat, fee_proportional_millionths, cltv_expiry_delta, funding_tx_id
`

type UpdatePendingChannelRequestByPubkeyParams struct {
	Pubkey           string         `db:"pubkey" json:"pubkey"`
	FundingAmount    sql.NullInt64  `db:"funding_amount" json:"fundingAmount"`
	FundingTxID      sql.NullString `db:"funding_tx_id" json:"fundingTxId"`
	FundingTxIDBytes []byte         `db:"funding_tx_id_bytes" json:"fundingTxIdBytes"`
	OutputIndex      sql.NullInt64  `db:"output_index" json:"outputIndex"`
}

func (q *Queries) UpdatePendingChannelRequestByPubkey(ctx context.Context, arg UpdatePendingChannelRequestByPubkeyParams) (ChannelRequest, error) {
	row := q.db.QueryRowContext(ctx, updatePendingChannelRequestByPubkey,
		arg.Pubkey,
		arg.FundingAmount,
		arg.FundingTxID,
		arg.FundingTxIDBytes,
		arg.OutputIndex,
	)
	var i ChannelRequest
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Status,
		&i.Pubkey,
		&i.PaymentHash,
		&i.PaymentAddr,
		&i.AmountMsat,
		&i.SettledMsat,
		&i.FundingTxIDBytes,
		&i.OutputIndex,
		&i.NodeID,
		&i.Amount,
		&i.FundingAmount,
		&i.PendingChanID,
		&i.Scid,
		&i.FeeBaseMsat,
		&i.FeeProportionalMillionths,
		&i.CltvExpiryDelta,
		&i.FundingTxID,
	)
	return i, err
}
