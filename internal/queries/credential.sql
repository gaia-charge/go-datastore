-- name: CreateCredential :one
INSERT INTO credentials (
    client_token, 
    server_token, 
    url, 
    business_detail_id, 
    country_code, 
    party_id, 
    is_hub,
    is_base64,
    last_updated
  ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
  RETURNING *;

-- name: DeleteCredential :exec
DELETE FROM credentials
  WHERE id = $1;

-- name: GetCredential :one
SELECT * FROM credentials
  WHERE id = $1;

-- name: GetCredentialByPartyAndCountryCode :one
SELECT * FROM credentials
  WHERE party_id = $1 AND country_code = $2;

-- name: GetCredentialByServerToken :one
SELECT * FROM credentials
  WHERE server_token = $1;

-- name: ListCredentials :many
SELECT * FROM credentials
  ORDER BY id;

-- name: UpdateCredential :one
UPDATE credentials SET (
    client_token, 
    server_token, 
    url, 
    country_code,
    party_id, 
    is_available,
    is_hub,
    is_base64,
    version_id,
    last_updated
  ) = ($2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
  WHERE id = $1
  RETURNING *;
