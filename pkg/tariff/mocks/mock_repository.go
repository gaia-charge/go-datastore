package mocks

import (
	mocks "github.com/gaia-charge/go-datastore/pkg/db/mocks"
	"github.com/gaia-charge/go-datastore/pkg/tariff"
)

func NewRepository(repositoryService *mocks.MockRepositoryService) tariff.TariffRepository {
	return tariff.TariffRepository(repositoryService)
}
