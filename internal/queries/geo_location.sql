-- name: CreateGeoLocation :one
INSERT INTO geo_locations (
     latitude, 
     latitude_float, 
     longitude, 
     longitude_float 
  ) VALUES ($1, $2, $3, $4)
  RETURNING *;

-- name: DeleteGeoLocation :exec
DELETE FROM geo_locations
  WHERE id = $1;

-- name: GetGeoLocation :one
SELECT * FROM geo_locations
  WHERE id = $1;

-- name: ListGeoLocations :many
SELECT * FROM geo_locations
  ORDER BY id;

-- name: UpdateGeoLocation :one
UPDATE geo_locations SET (
    latitude, 
    latitude_float, 
    longitude, 
    longitude_float 
  ) = ($2, $3, $4, $5)
  WHERE id = $1
  RETURNING *;
