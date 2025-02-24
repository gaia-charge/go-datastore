-- name: CreateVersionEndpoint :one
INSERT INTO version_endpoints (
    version_id, 
    identifier, 
    url
  ) VALUES ($1, $2, $3)
  RETURNING *;

-- name: DeleteVersionEndpoints :exec
DELETE FROM version_endpoints
  WHERE version_id = $1;

-- name: GetVersionEndpoint :one
SELECT * FROM version_endpoints
  WHERE id = $1;

-- name: GetVersionEndpointByIdentity :one
SELECT ve.* FROM version_endpoints ve
  INNER JOIN credentials c ON ve.version_id = c.version_id
  WHERE ve.identifier = $1 AND c.country_code = $2 AND c.party_id = $3;

-- name: ListVersionEndpoints :many
SELECT * FROM version_endpoints
  WHERE version_id = $1;
