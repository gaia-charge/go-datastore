package mocks

import (
	mocks "github.com/gaia-charge/go-datastore/pkg/db/mocks"
	"github.com/gaia-charge/go-datastore/pkg/tariffrestriction"
)

func NewRepository(repositoryService *mocks.MockRepositoryService) tariffrestriction.TariffRestrictionRepository {
	return tariffrestriction.TariffRestrictionRepository(repositoryService)
}
