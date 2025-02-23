package mocks

import (
	mocks "github.com/gaia-charge/go-datastore/pkg/db/mocks"
	"github.com/gaia-charge/go-datastore/pkg/elementrestriction"
)

func NewRepository(repositoryService *mocks.MockRepositoryService) elementrestriction.ElementRestrictionRepository {
	return elementrestriction.ElementRestrictionRepository(repositoryService)
}
