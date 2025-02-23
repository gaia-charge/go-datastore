package mocks

import (
	mocks "github.com/gaia-charge/go-datastore/pkg/db/mocks"
	"github.com/gaia-charge/go-datastore/pkg/pricecomponent"
)

func NewRepository(repositoryService *mocks.MockRepositoryService) pricecomponent.PriceComponentRepository {
	return pricecomponent.PriceComponentRepository(repositoryService)
}
