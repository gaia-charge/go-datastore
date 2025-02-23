package mocks

import (
	mocks "github.com/gaia-charge/go-datastore/pkg/db/mocks"
	"github.com/gaia-charge/go-datastore/pkg/evse"
)

func NewRepository(repositoryService *mocks.MockRepositoryService) evse.EvseRepository {
	return evse.EvseRepository(repositoryService)
}
