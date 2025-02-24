// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: energy_mix.sql

package db

import (
	"context"
	"database/sql"
)

const createEnergyMix = `-- name: CreateEnergyMix :one
INSERT INTO energy_mixes (
    is_green_energy, 
    supplier_name, 
    energy_product_name
  ) VALUES ($1, $2, $3)
  RETURNING id, is_green_energy, supplier_name, energy_product_name
`

type CreateEnergyMixParams struct {
	IsGreenEnergy     bool           `db:"is_green_energy" json:"isGreenEnergy"`
	SupplierName      sql.NullString `db:"supplier_name" json:"supplierName"`
	EnergyProductName sql.NullString `db:"energy_product_name" json:"energyProductName"`
}

func (q *Queries) CreateEnergyMix(ctx context.Context, arg CreateEnergyMixParams) (EnergyMix, error) {
	row := q.db.QueryRowContext(ctx, createEnergyMix, arg.IsGreenEnergy, arg.SupplierName, arg.EnergyProductName)
	var i EnergyMix
	err := row.Scan(
		&i.ID,
		&i.IsGreenEnergy,
		&i.SupplierName,
		&i.EnergyProductName,
	)
	return i, err
}

const deleteEnergyMix = `-- name: DeleteEnergyMix :exec
DELETE FROM energy_mixes
  WHERE id = $1
`

func (q *Queries) DeleteEnergyMix(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteEnergyMix, id)
	return err
}

const getEnergyMix = `-- name: GetEnergyMix :one
SELECT id, is_green_energy, supplier_name, energy_product_name FROM energy_mixes
  WHERE id = $1
`

func (q *Queries) GetEnergyMix(ctx context.Context, id int64) (EnergyMix, error) {
	row := q.db.QueryRowContext(ctx, getEnergyMix, id)
	var i EnergyMix
	err := row.Scan(
		&i.ID,
		&i.IsGreenEnergy,
		&i.SupplierName,
		&i.EnergyProductName,
	)
	return i, err
}

const updateEnergyMix = `-- name: UpdateEnergyMix :one
UPDATE energy_mixes SET (
    is_green_energy, 
    supplier_name, 
    energy_product_name
  ) = ($2, $3, $4)
  WHERE id = $1
  RETURNING id, is_green_energy, supplier_name, energy_product_name
`

type UpdateEnergyMixParams struct {
	ID                int64          `db:"id" json:"id"`
	IsGreenEnergy     bool           `db:"is_green_energy" json:"isGreenEnergy"`
	SupplierName      sql.NullString `db:"supplier_name" json:"supplierName"`
	EnergyProductName sql.NullString `db:"energy_product_name" json:"energyProductName"`
}

func (q *Queries) UpdateEnergyMix(ctx context.Context, arg UpdateEnergyMixParams) (EnergyMix, error) {
	row := q.db.QueryRowContext(ctx, updateEnergyMix,
		arg.ID,
		arg.IsGreenEnergy,
		arg.SupplierName,
		arg.EnergyProductName,
	)
	var i EnergyMix
	err := row.Scan(
		&i.ID,
		&i.IsGreenEnergy,
		&i.SupplierName,
		&i.EnergyProductName,
	)
	return i, err
}
