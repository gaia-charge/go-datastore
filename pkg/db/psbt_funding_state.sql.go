// Code generated by sqlc. DO NOT EDIT.
// source: psbt_funding_state.sql

package db

import (
	"context"
	"time"
)

const createPsbtFundingState = `-- name: CreatePsbtFundingState :one
INSERT INTO psbt_funding_states (
    node_id,
    base_psbt,
    psbt, 
    expiry_date
  ) VALUES ($1, $2, $3, $4)
  RETURNING id, node_id, base_psbt, psbt, funded_psbt, signed_psbt, signed_tx, expiry_date, is_failed
`

type CreatePsbtFundingStateParams struct {
	NodeID     int64     `db:"node_id" json:"nodeID"`
	BasePsbt   []byte    `db:"base_psbt" json:"basePsbt"`
	Psbt       []byte    `db:"psbt" json:"psbt"`
	ExpiryDate time.Time `db:"expiry_date" json:"expiryDate"`
}

func (q *Queries) CreatePsbtFundingState(ctx context.Context, arg CreatePsbtFundingStateParams) (PsbtFundingState, error) {
	row := q.db.QueryRowContext(ctx, createPsbtFundingState,
		arg.NodeID,
		arg.BasePsbt,
		arg.Psbt,
		arg.ExpiryDate,
	)
	var i PsbtFundingState
	err := row.Scan(
		&i.ID,
		&i.NodeID,
		&i.BasePsbt,
		&i.Psbt,
		&i.FundedPsbt,
		&i.SignedPsbt,
		&i.SignedTx,
		&i.ExpiryDate,
		&i.IsFailed,
	)
	return i, err
}

const getPsbtFundingState = `-- name: GetPsbtFundingState :one
SELECT id, node_id, base_psbt, psbt, funded_psbt, signed_psbt, signed_tx, expiry_date, is_failed FROM psbt_funding_states
  WHERE id = $1
`

func (q *Queries) GetPsbtFundingState(ctx context.Context, id int64) (PsbtFundingState, error) {
	row := q.db.QueryRowContext(ctx, getPsbtFundingState, id)
	var i PsbtFundingState
	err := row.Scan(
		&i.ID,
		&i.NodeID,
		&i.BasePsbt,
		&i.Psbt,
		&i.FundedPsbt,
		&i.SignedPsbt,
		&i.SignedTx,
		&i.ExpiryDate,
		&i.IsFailed,
	)
	return i, err
}

const getUnfundedPsbtFundingState = `-- name: GetUnfundedPsbtFundingState :one
SELECT id, node_id, base_psbt, psbt, funded_psbt, signed_psbt, signed_tx, expiry_date, is_failed FROM psbt_funding_states
  WHERE node_id = $1 AND funded_psbt is null AND not is_failed
`

func (q *Queries) GetUnfundedPsbtFundingState(ctx context.Context, nodeID int64) (PsbtFundingState, error) {
	row := q.db.QueryRowContext(ctx, getUnfundedPsbtFundingState, nodeID)
	var i PsbtFundingState
	err := row.Scan(
		&i.ID,
		&i.NodeID,
		&i.BasePsbt,
		&i.Psbt,
		&i.FundedPsbt,
		&i.SignedPsbt,
		&i.SignedTx,
		&i.ExpiryDate,
		&i.IsFailed,
	)
	return i, err
}

const listUnfundedPsbtFundingStates = `-- name: ListUnfundedPsbtFundingStates :many
SELECT id, node_id, base_psbt, psbt, funded_psbt, signed_psbt, signed_tx, expiry_date, is_failed FROM psbt_funding_states
  WHERE node_id = $1 AND funded_psbt is null AND not is_failed
  ORDER BY id
`

func (q *Queries) ListUnfundedPsbtFundingStates(ctx context.Context, nodeID int64) ([]PsbtFundingState, error) {
	rows, err := q.db.QueryContext(ctx, listUnfundedPsbtFundingStates, nodeID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []PsbtFundingState
	for rows.Next() {
		var i PsbtFundingState
		if err := rows.Scan(
			&i.ID,
			&i.NodeID,
			&i.BasePsbt,
			&i.Psbt,
			&i.FundedPsbt,
			&i.SignedPsbt,
			&i.SignedTx,
			&i.ExpiryDate,
			&i.IsFailed,
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

const updatePsbtFundingState = `-- name: UpdatePsbtFundingState :one
UPDATE psbt_funding_states SET (
    psbt,
    funded_psbt,
    signed_psbt,
    signed_tx,
    is_failed
  ) = ($2, $3, $4, $5, $6)
  WHERE id = $1
  RETURNING id, node_id, base_psbt, psbt, funded_psbt, signed_psbt, signed_tx, expiry_date, is_failed
`

type UpdatePsbtFundingStateParams struct {
	ID         int64  `db:"id" json:"id"`
	Psbt       []byte `db:"psbt" json:"psbt"`
	FundedPsbt []byte `db:"funded_psbt" json:"fundedPsbt"`
	SignedPsbt []byte `db:"signed_psbt" json:"signedPsbt"`
	SignedTx   []byte `db:"signed_tx" json:"signedTx"`
	IsFailed   bool   `db:"is_failed" json:"isFailed"`
}

func (q *Queries) UpdatePsbtFundingState(ctx context.Context, arg UpdatePsbtFundingStateParams) (PsbtFundingState, error) {
	row := q.db.QueryRowContext(ctx, updatePsbtFundingState,
		arg.ID,
		arg.Psbt,
		arg.FundedPsbt,
		arg.SignedPsbt,
		arg.SignedTx,
		arg.IsFailed,
	)
	var i PsbtFundingState
	err := row.Scan(
		&i.ID,
		&i.NodeID,
		&i.BasePsbt,
		&i.Psbt,
		&i.FundedPsbt,
		&i.SignedPsbt,
		&i.SignedTx,
		&i.ExpiryDate,
		&i.IsFailed,
	)
	return i, err
}
