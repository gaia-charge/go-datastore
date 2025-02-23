package node

import (
	"context"

	"github.com/gaia-charge/go-datastore/pkg/db"
)

type NodeRepository interface {
	CountNodeScids(ctx context.Context, nodeID int64) (int64, error)
	CreateNode(ctx context.Context, arg db.CreateNodeParams) (db.Node, error)
	CreateNodeScid(ctx context.Context, arg db.CreateNodeScidParams) (db.NodeScid, error)
	DeleteNodeScid(ctx context.Context, id int64) error
	GetNode(ctx context.Context, id int64) (db.Node, error)
	GetNodeByPubkey(ctx context.Context, pubkey string) (db.Node, error)
	GetNodeByUserID(ctx context.Context, id int64) (db.Node, error)
	GetNodeScid(ctx context.Context, nodeID int64) (db.NodeScid, error)
	ListNodes(ctx context.Context) ([]db.Node, error)
	ListActiveNodes(ctx context.Context, isLsp bool) ([]db.Node, error)
	UpdateNode(ctx context.Context, arg db.UpdateNodeParams) (db.Node, error)
}

func NewRepository(repositoryService *db.RepositoryService) NodeRepository {
	return NodeRepository(repositoryService)
}
