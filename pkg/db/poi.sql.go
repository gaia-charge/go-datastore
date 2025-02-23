// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: poi.sql

package db

import (
	"context"
	"database/sql"
	"time"

	"github.com/gaia-charge/go-datastore/pkg/geom"
)

const createPoi = `-- name: CreatePoi :one
INSERT INTO pois (
    uid,
    source,
    name,
    geom,
    description,
    address,
    city,
    postal_code,
    tag_key,
    tag_value,
    payment_on_chain,
    payment_ln,
    payment_ln_tap,
    payment_uri,
    opening_times,
    phone,
    website,
    last_updated
  ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18)
  RETURNING id, uid, source, name, geom, description, address, city, postal_code, tag_key, tag_value, payment_on_chain, payment_ln, payment_ln_tap, opening_times, phone, website, last_updated, payment_uri
`

type CreatePoiParams struct {
	Uid            string            `db:"uid" json:"uid"`
	Source         string            `db:"source" json:"source"`
	Name           string            `db:"name" json:"name"`
	Geom           geom.Geometry4326 `db:"geom" json:"geom"`
	Description    sql.NullString    `db:"description" json:"description"`
	Address        sql.NullString    `db:"address" json:"address"`
	City           sql.NullString    `db:"city" json:"city"`
	PostalCode     sql.NullString    `db:"postal_code" json:"postalCode"`
	TagKey         string            `db:"tag_key" json:"tagKey"`
	TagValue       string            `db:"tag_value" json:"tagValue"`
	PaymentOnChain bool              `db:"payment_on_chain" json:"paymentOnChain"`
	PaymentLn      bool              `db:"payment_ln" json:"paymentLn"`
	PaymentLnTap   bool              `db:"payment_ln_tap" json:"paymentLnTap"`
	PaymentUri     sql.NullString    `db:"payment_uri" json:"paymentUri"`
	OpeningTimes   sql.NullString    `db:"opening_times" json:"openingTimes"`
	Phone          sql.NullString    `db:"phone" json:"phone"`
	Website        sql.NullString    `db:"website" json:"website"`
	LastUpdated    time.Time         `db:"last_updated" json:"lastUpdated"`
}

func (q *Queries) CreatePoi(ctx context.Context, arg CreatePoiParams) (Poi, error) {
	row := q.db.QueryRowContext(ctx, createPoi,
		arg.Uid,
		arg.Source,
		arg.Name,
		arg.Geom,
		arg.Description,
		arg.Address,
		arg.City,
		arg.PostalCode,
		arg.TagKey,
		arg.TagValue,
		arg.PaymentOnChain,
		arg.PaymentLn,
		arg.PaymentLnTap,
		arg.PaymentUri,
		arg.OpeningTimes,
		arg.Phone,
		arg.Website,
		arg.LastUpdated,
	)
	var i Poi
	err := row.Scan(
		&i.ID,
		&i.Uid,
		&i.Source,
		&i.Name,
		&i.Geom,
		&i.Description,
		&i.Address,
		&i.City,
		&i.PostalCode,
		&i.TagKey,
		&i.TagValue,
		&i.PaymentOnChain,
		&i.PaymentLn,
		&i.PaymentLnTap,
		&i.OpeningTimes,
		&i.Phone,
		&i.Website,
		&i.LastUpdated,
		&i.PaymentUri,
	)
	return i, err
}

const deletePoiByUid = `-- name: DeletePoiByUid :exec
DELETE FROM pois
  WHERE uid = $1
`

func (q *Queries) DeletePoiByUid(ctx context.Context, uid string) error {
	_, err := q.db.ExecContext(ctx, deletePoiByUid, uid)
	return err
}

const getPoiByLastUpdated = `-- name: GetPoiByLastUpdated :one
SELECT id, uid, source, name, geom, description, address, city, postal_code, tag_key, tag_value, payment_on_chain, payment_ln, payment_ln_tap, opening_times, phone, website, last_updated, payment_uri FROM pois
  ORDER BY last_updated DESC
  LIMIT 1
`

func (q *Queries) GetPoiByLastUpdated(ctx context.Context) (Poi, error) {
	row := q.db.QueryRowContext(ctx, getPoiByLastUpdated)
	var i Poi
	err := row.Scan(
		&i.ID,
		&i.Uid,
		&i.Source,
		&i.Name,
		&i.Geom,
		&i.Description,
		&i.Address,
		&i.City,
		&i.PostalCode,
		&i.TagKey,
		&i.TagValue,
		&i.PaymentOnChain,
		&i.PaymentLn,
		&i.PaymentLnTap,
		&i.OpeningTimes,
		&i.Phone,
		&i.Website,
		&i.LastUpdated,
		&i.PaymentUri,
	)
	return i, err
}

const getPoiByUid = `-- name: GetPoiByUid :one
SELECT id, uid, source, name, geom, description, address, city, postal_code, tag_key, tag_value, payment_on_chain, payment_ln, payment_ln_tap, opening_times, phone, website, last_updated, payment_uri FROM pois
  WHERE uid = $1
`

func (q *Queries) GetPoiByUid(ctx context.Context, uid string) (Poi, error) {
	row := q.db.QueryRowContext(ctx, getPoiByUid, uid)
	var i Poi
	err := row.Scan(
		&i.ID,
		&i.Uid,
		&i.Source,
		&i.Name,
		&i.Geom,
		&i.Description,
		&i.Address,
		&i.City,
		&i.PostalCode,
		&i.TagKey,
		&i.TagValue,
		&i.PaymentOnChain,
		&i.PaymentLn,
		&i.PaymentLnTap,
		&i.OpeningTimes,
		&i.Phone,
		&i.Website,
		&i.LastUpdated,
		&i.PaymentUri,
	)
	return i, err
}

const listPoisByGeom = `-- name: ListPoisByGeom :many
SELECT id, uid, source, name, geom, description, address, city, postal_code, tag_key, tag_value, payment_on_chain, payment_ln, payment_ln_tap, opening_times, phone, website, last_updated, payment_uri FROM pois
  WHERE ST_Intersects(geom, ST_MakeEnvelope($1::FLOAT, $2::FLOAT, $3::FLOAT, $4::FLOAT, 4326))
  LIMIT 100
`

type ListPoisByGeomParams struct {
	XMin float64 `db:"x_min" json:"xMin"`
	YMin float64 `db:"y_min" json:"yMin"`
	XMax float64 `db:"x_max" json:"xMax"`
	YMax float64 `db:"y_max" json:"yMax"`
}

func (q *Queries) ListPoisByGeom(ctx context.Context, arg ListPoisByGeomParams) ([]Poi, error) {
	rows, err := q.db.QueryContext(ctx, listPoisByGeom,
		arg.XMin,
		arg.YMin,
		arg.XMax,
		arg.YMax,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Poi
	for rows.Next() {
		var i Poi
		if err := rows.Scan(
			&i.ID,
			&i.Uid,
			&i.Source,
			&i.Name,
			&i.Geom,
			&i.Description,
			&i.Address,
			&i.City,
			&i.PostalCode,
			&i.TagKey,
			&i.TagValue,
			&i.PaymentOnChain,
			&i.PaymentLn,
			&i.PaymentLnTap,
			&i.OpeningTimes,
			&i.Phone,
			&i.Website,
			&i.LastUpdated,
			&i.PaymentUri,
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

const updatePoiByUid = `-- name: UpdatePoiByUid :one
UPDATE pois SET (
    name,
    geom,
    description,
    address,
    city,
    postal_code,
    tag_key,
    tag_value,
    payment_on_chain,
    payment_ln,
    payment_ln_tap,
    payment_uri,
    opening_times,
    phone,
    website,
    last_updated
  ) = ($2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17)
  WHERE uid = $1
  RETURNING id, uid, source, name, geom, description, address, city, postal_code, tag_key, tag_value, payment_on_chain, payment_ln, payment_ln_tap, opening_times, phone, website, last_updated, payment_uri
`

type UpdatePoiByUidParams struct {
	Uid            string            `db:"uid" json:"uid"`
	Name           string            `db:"name" json:"name"`
	Geom           geom.Geometry4326 `db:"geom" json:"geom"`
	Description    sql.NullString    `db:"description" json:"description"`
	Address        sql.NullString    `db:"address" json:"address"`
	City           sql.NullString    `db:"city" json:"city"`
	PostalCode     sql.NullString    `db:"postal_code" json:"postalCode"`
	TagKey         string            `db:"tag_key" json:"tagKey"`
	TagValue       string            `db:"tag_value" json:"tagValue"`
	PaymentOnChain bool              `db:"payment_on_chain" json:"paymentOnChain"`
	PaymentLn      bool              `db:"payment_ln" json:"paymentLn"`
	PaymentLnTap   bool              `db:"payment_ln_tap" json:"paymentLnTap"`
	PaymentUri     sql.NullString    `db:"payment_uri" json:"paymentUri"`
	OpeningTimes   sql.NullString    `db:"opening_times" json:"openingTimes"`
	Phone          sql.NullString    `db:"phone" json:"phone"`
	Website        sql.NullString    `db:"website" json:"website"`
	LastUpdated    time.Time         `db:"last_updated" json:"lastUpdated"`
}

func (q *Queries) UpdatePoiByUid(ctx context.Context, arg UpdatePoiByUidParams) (Poi, error) {
	row := q.db.QueryRowContext(ctx, updatePoiByUid,
		arg.Uid,
		arg.Name,
		arg.Geom,
		arg.Description,
		arg.Address,
		arg.City,
		arg.PostalCode,
		arg.TagKey,
		arg.TagValue,
		arg.PaymentOnChain,
		arg.PaymentLn,
		arg.PaymentLnTap,
		arg.PaymentUri,
		arg.OpeningTimes,
		arg.Phone,
		arg.Website,
		arg.LastUpdated,
	)
	var i Poi
	err := row.Scan(
		&i.ID,
		&i.Uid,
		&i.Source,
		&i.Name,
		&i.Geom,
		&i.Description,
		&i.Address,
		&i.City,
		&i.PostalCode,
		&i.TagKey,
		&i.TagValue,
		&i.PaymentOnChain,
		&i.PaymentLn,
		&i.PaymentLnTap,
		&i.OpeningTimes,
		&i.Phone,
		&i.Website,
		&i.LastUpdated,
		&i.PaymentUri,
	)
	return i, err
}
