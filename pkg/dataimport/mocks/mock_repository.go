package mocks

import (
	"github.com/gaia-charge/go-datastore/pkg/dataimport"
	mocks "github.com/gaia-charge/go-datastore/pkg/db/mocks"
)

func NewRepository(repositoryService *mocks.MockRepositoryService) dataimport.DataImportRepository {
	return dataimport.DataImportRepository(repositoryService)
}
