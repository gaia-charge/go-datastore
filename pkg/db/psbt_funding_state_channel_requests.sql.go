// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: psbt_funding_state_channel_requests.sql

package db

import (
	"context"
)

const listPsbtFundingStateChannelRequests = `-- name: ListPsbtFundingStateChannelRequests :many
SELECT cr.id, cr.user_id, cr.status, cr.pubkey, cr.payment_hash, cr.payment_addr, cr.amount_msat, cr.settled_msat, cr.funding_tx_id_bytes, cr.output_index, cr.node_id, cr.amount, cr.funding_amount, cr.pending_chan_id, cr.scid, cr.fee_base_msat, cr.fee_proportional_millionths, cr.cltv_expiry_delta, cr.funding_tx_id FROM channel_requests cr
  INNER JOIN psbt_funding_state_channel_requests pfscr ON pfscr.channel_request_id = cr.id
  WHERE pfscr.psbt_funding_state_id = $1
  ORDER BY cr.id
`

func (q *Queries) ListPsbtFundingStateChannelRequests(ctx context.Context, psbtFundingStateID int64) ([]ChannelRequest, error) {
	rows, err := q.db.QueryContext(ctx, listPsbtFundingStateChannelRequests, psbtFundingStateID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ChannelRequest
	for rows.Next() {
		var i ChannelRequest
		if err := rows.Scan(
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

const setPsbtFundingStateChannelRequest = `-- name: SetPsbtFundingStateChannelRequest :exec
INSERT INTO psbt_funding_state_channel_requests (psbt_funding_state_id, channel_request_id)
  VALUES ($1, $2)
`

type SetPsbtFundingStateChannelRequestParams struct {
	PsbtFundingStateID int64 `db:"psbt_funding_state_id" json:"psbtFundingStateId"`
	ChannelRequestID   int64 `db:"channel_request_id" json:"channelRequestId"`
}

func (q *Queries) SetPsbtFundingStateChannelRequest(ctx context.Context, arg SetPsbtFundingStateChannelRequestParams) error {
	_, err := q.db.ExecContext(ctx, setPsbtFundingStateChannelRequest, arg.PsbtFundingStateID, arg.ChannelRequestID)
	return err
}
