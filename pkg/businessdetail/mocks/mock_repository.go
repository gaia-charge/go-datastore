package mocks

import (
	"github.com/gaia-charge/go-datastore/pkg/businessdetail"
	mocks "github.com/gaia-charge/go-datastore/pkg/db/mocks"
)

func NewRepository(repositoryService *mocks.MockRepositoryService) businessdetail.BusinessDetailRepository {
	return businessdetail.BusinessDetailRepository(repositoryService)
}
