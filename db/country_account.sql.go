// Code generated by sqlc. DO NOT EDIT.
// source: country_account.sql

package db

import (
	"context"
)

const createCountryAccount = `-- name: CreateCountryAccount :one
INSERT INTO country_accounts (
    country,
    tax_percent
  ) VALUES ($1, $2)
  RETURNING id, country, tax_percent
`

type CreateCountryAccountParams struct {
	Country    string  `db:"country" json:"country"`
	TaxPercent float64 `db:"tax_percent" json:"taxPercent"`
}

func (q *Queries) CreateCountryAccount(ctx context.Context, arg CreateCountryAccountParams) (CountryAccount, error) {
	row := q.db.QueryRowContext(ctx, createCountryAccount, arg.Country, arg.TaxPercent)
	var i CountryAccount
	err := row.Scan(&i.ID, &i.Country, &i.TaxPercent)
	return i, err
}

const getCountryAccountByCountry = `-- name: GetCountryAccountByCountry :one
SELECT id, country, tax_percent FROM country_accounts
  WHERE country = $1
`

func (q *Queries) GetCountryAccountByCountry(ctx context.Context, country string) (CountryAccount, error) {
	row := q.db.QueryRowContext(ctx, getCountryAccountByCountry, country)
	var i CountryAccount
	err := row.Scan(&i.ID, &i.Country, &i.TaxPercent)
	return i, err
}
