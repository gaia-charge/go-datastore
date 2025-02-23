// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: element_restriction_weekday.sql

package db

import (
	"context"
)

const listElementRestrictionWeekdays = `-- name: ListElementRestrictionWeekdays :many
SELECT w.id, w.text, w.description FROM weekdays w
  INNER JOIN element_restriction_weekdays rw ON rw.weekday_id = w.id
  WHERE rw.element_restriction_id = $1
  ORDER BY w.id
`

func (q *Queries) ListElementRestrictionWeekdays(ctx context.Context, elementRestrictionID int64) ([]Weekday, error) {
	rows, err := q.db.QueryContext(ctx, listElementRestrictionWeekdays, elementRestrictionID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Weekday
	for rows.Next() {
		var i Weekday
		if err := rows.Scan(&i.ID, &i.Text, &i.Description); err != nil {
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

const setElementRestrictionWeekday = `-- name: SetElementRestrictionWeekday :exec
INSERT INTO element_restriction_weekdays 
    (element_restriction_id, weekday_id)
  VALUES ($1, $2)
`

type SetElementRestrictionWeekdayParams struct {
	ElementRestrictionID int64 `db:"element_restriction_id" json:"elementRestrictionId"`
	WeekdayID            int64 `db:"weekday_id" json:"weekdayId"`
}

func (q *Queries) SetElementRestrictionWeekday(ctx context.Context, arg SetElementRestrictionWeekdayParams) error {
	_, err := q.db.ExecContext(ctx, setElementRestrictionWeekday, arg.ElementRestrictionID, arg.WeekdayID)
	return err
}

const unsetElementRestrictionWeekdays = `-- name: UnsetElementRestrictionWeekdays :exec
DELETE FROM element_restriction_weekdays rw
  WHERE rw.element_restriction_id = $1
`

func (q *Queries) UnsetElementRestrictionWeekdays(ctx context.Context, elementRestrictionID int64) error {
	_, err := q.db.ExecContext(ctx, unsetElementRestrictionWeekdays, elementRestrictionID)
	return err
}
