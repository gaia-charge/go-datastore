package mocks

import (
	mocks "github.com/gaia-charge/go-datastore/pkg/db/mocks"
	"github.com/gaia-charge/go-datastore/pkg/node"
)

func NewRepository(repositoryService *mocks.MockRepositoryService) node.NodeRepository {
	return node.NodeRepository(repositoryService)
}
