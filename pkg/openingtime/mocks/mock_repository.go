package mocks

import (
	mocks "github.com/gaia-charge/go-datastore/pkg/db/mocks"
	"github.com/gaia-charge/go-datastore/pkg/openingtime"
)

func NewRepository(repositoryService *mocks.MockRepositoryService) openingtime.OpeningTimeRepository {
	return openingtime.OpeningTimeRepository(repositoryService)
}
