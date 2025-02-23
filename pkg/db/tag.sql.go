// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: tag.sql

package db

import (
	"context"
)

const createTag = `-- name: CreateTag :one
INSERT INTO tags (
    key, 
    value
  ) VALUES ($1, $2)
  RETURNING id, key, value
`

type CreateTagParams struct {
	Key   string `db:"key" json:"key"`
	Value string `db:"value" json:"value"`
}

func (q *Queries) CreateTag(ctx context.Context, arg CreateTagParams) (Tag, error) {
	row := q.db.QueryRowContext(ctx, createTag, arg.Key, arg.Value)
	var i Tag
	err := row.Scan(&i.ID, &i.Key, &i.Value)
	return i, err
}

const getTagByKeyValue = `-- name: GetTagByKeyValue :one
SELECT id, key, value FROM tags
  WHERE key = $1 AND value = $2
  LIMIT 1
`

type GetTagByKeyValueParams struct {
	Key   string `db:"key" json:"key"`
	Value string `db:"value" json:"value"`
}

func (q *Queries) GetTagByKeyValue(ctx context.Context, arg GetTagByKeyValueParams) (Tag, error) {
	row := q.db.QueryRowContext(ctx, getTagByKeyValue, arg.Key, arg.Value)
	var i Tag
	err := row.Scan(&i.ID, &i.Key, &i.Value)
	return i, err
}
