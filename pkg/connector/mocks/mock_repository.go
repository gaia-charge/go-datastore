package mocks

import (
	"github.com/gaia-charge/go-datastore/pkg/connector"
	mocks "github.com/gaia-charge/go-datastore/pkg/db/mocks"
)

func NewRepository(repositoryService *mocks.MockRepositoryService) connector.ConnectorRepository {
	return connector.ConnectorRepository(repositoryService)
}
