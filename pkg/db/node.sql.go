// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: node.sql

package db

import (
	"context"
)

const createNode = `-- name: CreateNode :one
INSERT INTO nodes (
    pubkey,
    node_addr,
    rpc_addr,
    alias,
    color,
    commit_hash,
    version,
    channels,
    peers,
    is_active,
    is_lsp
  ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
  RETURNING id, pubkey, node_addr, rpc_addr, alias, color, commit_hash, version, channels, peers, is_active, is_lsp
`

type CreateNodeParams struct {
	Pubkey     string `db:"pubkey" json:"pubkey"`
	NodeAddr   string `db:"node_addr" json:"nodeAddr"`
	RpcAddr    string `db:"rpc_addr" json:"rpcAddr"`
	Alias      string `db:"alias" json:"alias"`
	Color      string `db:"color" json:"color"`
	CommitHash string `db:"commit_hash" json:"commitHash"`
	Version    string `db:"version" json:"version"`
	Channels   int64  `db:"channels" json:"channels"`
	Peers      int64  `db:"peers" json:"peers"`
	IsActive   bool   `db:"is_active" json:"isActive"`
	IsLsp      bool   `db:"is_lsp" json:"isLsp"`
}

func (q *Queries) CreateNode(ctx context.Context, arg CreateNodeParams) (Node, error) {
	row := q.db.QueryRowContext(ctx, createNode,
		arg.Pubkey,
		arg.NodeAddr,
		arg.RpcAddr,
		arg.Alias,
		arg.Color,
		arg.CommitHash,
		arg.Version,
		arg.Channels,
		arg.Peers,
		arg.IsActive,
		arg.IsLsp,
	)
	var i Node
	err := row.Scan(
		&i.ID,
		&i.Pubkey,
		&i.NodeAddr,
		&i.RpcAddr,
		&i.Alias,
		&i.Color,
		&i.CommitHash,
		&i.Version,
		&i.Channels,
		&i.Peers,
		&i.IsActive,
		&i.IsLsp,
	)
	return i, err
}

const getNode = `-- name: GetNode :one
SELECT id, pubkey, node_addr, rpc_addr, alias, color, commit_hash, version, channels, peers, is_active, is_lsp FROM nodes
  WHERE id = $1
`

func (q *Queries) GetNode(ctx context.Context, id int64) (Node, error) {
	row := q.db.QueryRowContext(ctx, getNode, id)
	var i Node
	err := row.Scan(
		&i.ID,
		&i.Pubkey,
		&i.NodeAddr,
		&i.RpcAddr,
		&i.Alias,
		&i.Color,
		&i.CommitHash,
		&i.Version,
		&i.Channels,
		&i.Peers,
		&i.IsActive,
		&i.IsLsp,
	)
	return i, err
}

const getNodeByPubkey = `-- name: GetNodeByPubkey :one
SELECT id, pubkey, node_addr, rpc_addr, alias, color, commit_hash, version, channels, peers, is_active, is_lsp FROM nodes
  WHERE pubkey = $1
`

func (q *Queries) GetNodeByPubkey(ctx context.Context, pubkey string) (Node, error) {
	row := q.db.QueryRowContext(ctx, getNodeByPubkey, pubkey)
	var i Node
	err := row.Scan(
		&i.ID,
		&i.Pubkey,
		&i.NodeAddr,
		&i.RpcAddr,
		&i.Alias,
		&i.Color,
		&i.CommitHash,
		&i.Version,
		&i.Channels,
		&i.Peers,
		&i.IsActive,
		&i.IsLsp,
	)
	return i, err
}

const getNodeByUserID = `-- name: GetNodeByUserID :one
SELECT n.id, n.pubkey, n.node_addr, n.rpc_addr, n.alias, n.color, n.commit_hash, n.version, n.channels, n.peers, n.is_active, n.is_lsp FROM nodes n
  INNER JOIN users u ON u.node_id = n.id
  WHERE u.id = $1
`

func (q *Queries) GetNodeByUserID(ctx context.Context, id int64) (Node, error) {
	row := q.db.QueryRowContext(ctx, getNodeByUserID, id)
	var i Node
	err := row.Scan(
		&i.ID,
		&i.Pubkey,
		&i.NodeAddr,
		&i.RpcAddr,
		&i.Alias,
		&i.Color,
		&i.CommitHash,
		&i.Version,
		&i.Channels,
		&i.Peers,
		&i.IsActive,
		&i.IsLsp,
	)
	return i, err
}

const listActiveNodes = `-- name: ListActiveNodes :many
SELECT id, pubkey, node_addr, rpc_addr, alias, color, commit_hash, version, channels, peers, is_active, is_lsp FROM nodes
  WHERE is_active = true AND is_lsp = $1
  ORDER BY peers ASC
`

func (q *Queries) ListActiveNodes(ctx context.Context, isLsp bool) ([]Node, error) {
	rows, err := q.db.QueryContext(ctx, listActiveNodes, isLsp)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Node
	for rows.Next() {
		var i Node
		if err := rows.Scan(
			&i.ID,
			&i.Pubkey,
			&i.NodeAddr,
			&i.RpcAddr,
			&i.Alias,
			&i.Color,
			&i.CommitHash,
			&i.Version,
			&i.Channels,
			&i.Peers,
			&i.IsActive,
			&i.IsLsp,
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

const listNodes = `-- name: ListNodes :many
SELECT id, pubkey, node_addr, rpc_addr, alias, color, commit_hash, version, channels, peers, is_active, is_lsp FROM nodes
  ORDER BY peers ASC
`

func (q *Queries) ListNodes(ctx context.Context) ([]Node, error) {
	rows, err := q.db.QueryContext(ctx, listNodes)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Node
	for rows.Next() {
		var i Node
		if err := rows.Scan(
			&i.ID,
			&i.Pubkey,
			&i.NodeAddr,
			&i.RpcAddr,
			&i.Alias,
			&i.Color,
			&i.CommitHash,
			&i.Version,
			&i.Channels,
			&i.Peers,
			&i.IsActive,
			&i.IsLsp,
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

const updateNode = `-- name: UpdateNode :one
UPDATE nodes SET (
    node_addr,
    rpc_addr,
    alias,
    color,
    commit_hash,
    version,
    channels,
    peers,
    is_active,
    is_lsp
  ) = ($2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
  WHERE id = $1
  RETURNING id, pubkey, node_addr, rpc_addr, alias, color, commit_hash, version, channels, peers, is_active, is_lsp
`

type UpdateNodeParams struct {
	ID         int64  `db:"id" json:"id"`
	NodeAddr   string `db:"node_addr" json:"nodeAddr"`
	RpcAddr    string `db:"rpc_addr" json:"rpcAddr"`
	Alias      string `db:"alias" json:"alias"`
	Color      string `db:"color" json:"color"`
	CommitHash string `db:"commit_hash" json:"commitHash"`
	Version    string `db:"version" json:"version"`
	Channels   int64  `db:"channels" json:"channels"`
	Peers      int64  `db:"peers" json:"peers"`
	IsActive   bool   `db:"is_active" json:"isActive"`
	IsLsp      bool   `db:"is_lsp" json:"isLsp"`
}

func (q *Queries) UpdateNode(ctx context.Context, arg UpdateNodeParams) (Node, error) {
	row := q.db.QueryRowContext(ctx, updateNode,
		arg.ID,
		arg.NodeAddr,
		arg.RpcAddr,
		arg.Alias,
		arg.Color,
		arg.CommitHash,
		arg.Version,
		arg.Channels,
		arg.Peers,
		arg.IsActive,
		arg.IsLsp,
	)
	var i Node
	err := row.Scan(
		&i.ID,
		&i.Pubkey,
		&i.NodeAddr,
		&i.RpcAddr,
		&i.Alias,
		&i.Color,
		&i.CommitHash,
		&i.Version,
		&i.Channels,
		&i.Peers,
		&i.IsActive,
		&i.IsLsp,
	)
	return i, err
}
