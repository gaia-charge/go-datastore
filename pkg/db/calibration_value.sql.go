// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: calibration_value.sql

package db

import (
	"context"
)

const createCalibrationValue = `-- name: CreateCalibrationValue :one
INSERT INTO calibration_values (
    calibration_id,
    nature,
    plain_data,
    signed_data
  ) VALUES ($1, $2, $3, $4)
  RETURNING id, calibration_id, nature, plain_data, signed_data
`

type CreateCalibrationValueParams struct {
	CalibrationID int64  `db:"calibration_id" json:"calibrationID"`
	Nature        string `db:"nature" json:"nature"`
	PlainData     string `db:"plain_data" json:"plainData"`
	SignedData    string `db:"signed_data" json:"signedData"`
}

func (q *Queries) CreateCalibrationValue(ctx context.Context, arg CreateCalibrationValueParams) (CalibrationValue, error) {
	row := q.db.QueryRowContext(ctx, createCalibrationValue,
		arg.CalibrationID,
		arg.Nature,
		arg.PlainData,
		arg.SignedData,
	)
	var i CalibrationValue
	err := row.Scan(
		&i.ID,
		&i.CalibrationID,
		&i.Nature,
		&i.PlainData,
		&i.SignedData,
	)
	return i, err
}

const listCalibrationValues = `-- name: ListCalibrationValues :many
SELECT id, calibration_id, nature, plain_data, signed_data FROM calibration_values
  WHERE calibration_id = $1
  ORDER BY id
`

func (q *Queries) ListCalibrationValues(ctx context.Context, calibrationID int64) ([]CalibrationValue, error) {
	rows, err := q.db.QueryContext(ctx, listCalibrationValues, calibrationID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []CalibrationValue
	for rows.Next() {
		var i CalibrationValue
		if err := rows.Scan(
			&i.ID,
			&i.CalibrationID,
			&i.Nature,
			&i.PlainData,
			&i.SignedData,
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
