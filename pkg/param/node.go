package param

import "github.com/gaia-charge/go-datastore/pkg/db"

func NewUpdateNodeParams(node db.Node) db.UpdateNodeParams {
	return db.UpdateNodeParams{
		ID:         node.ID,
		NodeAddr:   node.NodeAddr,
		RpcAddr:    node.RpcAddr,
		Alias:      node.Alias,
		Color:      node.Color,
		CommitHash: node.CommitHash,
		Version:    node.Version,
		Channels:   node.Channels,
		Peers:      node.Peers,
		IsActive:   node.IsActive,
		IsLsp:      node.IsLsp,
	}
}
