// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: referral.sql

package db

import (
	"context"
	"time"
)

const createReferral = `-- name: CreateReferral :one
INSERT INTO referrals (
    promotion_id,
    user_id,
    ip_address,
    last_updated
  ) VALUES ($1, $2, $3, $4)
  RETURNING id, promotion_id, user_id, ip_address, last_updated
`

type CreateReferralParams struct {
	PromotionID int64     `db:"promotion_id" json:"promotionId"`
	UserID      int64     `db:"user_id" json:"userId"`
	IpAddress   string    `db:"ip_address" json:"ipAddress"`
	LastUpdated time.Time `db:"last_updated" json:"lastUpdated"`
}

func (q *Queries) CreateReferral(ctx context.Context, arg CreateReferralParams) (Referral, error) {
	row := q.db.QueryRowContext(ctx, createReferral,
		arg.PromotionID,
		arg.UserID,
		arg.IpAddress,
		arg.LastUpdated,
	)
	var i Referral
	err := row.Scan(
		&i.ID,
		&i.PromotionID,
		&i.UserID,
		&i.IpAddress,
		&i.LastUpdated,
	)
	return i, err
}

const getReferralByIpAddress = `-- name: GetReferralByIpAddress :one
SELECT id, promotion_id, user_id, ip_address, last_updated FROM referrals
  WHERE ip_address = $1
`

func (q *Queries) GetReferralByIpAddress(ctx context.Context, ipAddress string) (Referral, error) {
	row := q.db.QueryRowContext(ctx, getReferralByIpAddress, ipAddress)
	var i Referral
	err := row.Scan(
		&i.ID,
		&i.PromotionID,
		&i.UserID,
		&i.IpAddress,
		&i.LastUpdated,
	)
	return i, err
}
