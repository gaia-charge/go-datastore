// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: currency.sql

package db

import (
	"context"
	"database/sql"
)

const createCurrency = `-- name: CreateCurrency :one
INSERT INTO currencies (
    name,
    code,
    numeric_code,
    decimals,
    symbol
  ) VALUES ($1, $2, $3, $4, $5)
  RETURNING id, name, code, numeric_code, decimals, symbol
`

type CreateCurrencyParams struct {
	Name        string         `db:"name" json:"name"`
	Code        string         `db:"code" json:"code"`
	NumericCode int32          `db:"numeric_code" json:"numericCode"`
	Decimals    int32          `db:"decimals" json:"decimals"`
	Symbol      sql.NullString `db:"symbol" json:"symbol"`
}

func (q *Queries) CreateCurrency(ctx context.Context, arg CreateCurrencyParams) (Currency, error) {
	row := q.db.QueryRowContext(ctx, createCurrency,
		arg.Name,
		arg.Code,
		arg.NumericCode,
		arg.Decimals,
		arg.Symbol,
	)
	var i Currency
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Code,
		&i.NumericCode,
		&i.Decimals,
		&i.Symbol,
	)
	return i, err
}

const getCurrencyByCode = `-- name: GetCurrencyByCode :one
SELECT id, name, code, numeric_code, decimals, symbol FROM currencies
  WHERE code = $1
`

func (q *Queries) GetCurrencyByCode(ctx context.Context, code string) (Currency, error) {
	row := q.db.QueryRowContext(ctx, getCurrencyByCode, code)
	var i Currency
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Code,
		&i.NumericCode,
		&i.Decimals,
		&i.Symbol,
	)
	return i, err
}
