package mocks

import (
	mocks "github.com/gaia-charge/go-datastore/pkg/db/mocks"
	"github.com/gaia-charge/go-datastore/pkg/energymix"
)

func NewRepository(repositoryService *mocks.MockRepositoryService) energymix.EnergyMixRepository {
	return energymix.EnergyMixRepository(repositoryService)
}
