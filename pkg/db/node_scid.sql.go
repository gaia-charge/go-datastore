// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: node_scid.sql

package db

import (
	"context"
)

const countNodeScids = `-- name: CountNodeScids :one
SELECT COUNT(*) FROM node_scids
  WHERE node_id = $1
`

func (q *Queries) CountNodeScids(ctx context.Context, nodeID int64) (int64, error) {
	row := q.db.QueryRowContext(ctx, countNodeScids, nodeID)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const createNodeScid = `-- name: CreateNodeScid :one
INSERT INTO node_scids (
    node_id, 
    scid
  ) VALUES ($1, $2)
  RETURNING id, node_id, scid
`

type CreateNodeScidParams struct {
	NodeID int64  `db:"node_id" json:"nodeId"`
	Scid   []byte `db:"scid" json:"scid"`
}

func (q *Queries) CreateNodeScid(ctx context.Context, arg CreateNodeScidParams) (NodeScid, error) {
	row := q.db.QueryRowContext(ctx, createNodeScid, arg.NodeID, arg.Scid)
	var i NodeScid
	err := row.Scan(&i.ID, &i.NodeID, &i.Scid)
	return i, err
}

const deleteNodeScid = `-- name: DeleteNodeScid :exec
DELETE FROM node_scids
  WHERE id = $1
`

func (q *Queries) DeleteNodeScid(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteNodeScid, id)
	return err
}

const getNodeScid = `-- name: GetNodeScid :one
SELECT id, node_id, scid FROM node_scids
  WHERE node_id = $1
  ORDER BY id ASC
  LIMIT 1
`

func (q *Queries) GetNodeScid(ctx context.Context, nodeID int64) (NodeScid, error) {
	row := q.db.QueryRowContext(ctx, getNodeScid, nodeID)
	var i NodeScid
	err := row.Scan(&i.ID, &i.NodeID, &i.Scid)
	return i, err
}
