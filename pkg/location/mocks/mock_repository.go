package mocks

import (
	mocks "github.com/gaia-charge/go-datastore/pkg/db/mocks"
	"github.com/gaia-charge/go-datastore/pkg/location"
)

func NewRepository(repositoryService *mocks.MockRepositoryService) location.LocationRepository {
	return location.LocationRepository(repositoryService)
}
