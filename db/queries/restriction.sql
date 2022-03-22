-- name: CreateRestriction :one
INSERT INTO restrictions (
    start_time,
    end_time,
    start_date,
    end_date,
    min_kwh,
    max_kwh,
    min_power,
    max_power,
    min_duration,
    max_duration
  ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
  RETURNING *;

-- name: DeleteRestrictions :exec
DELETE FROM restrictions r
  USING elements e
  WHERE e.restriction_id == r.id AND e.tariff_id == $1;

-- name: GetRestriction :one
SELECT * FROM restrictions
  WHERE id = $1;
