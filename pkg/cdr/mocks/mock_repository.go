package mocks

import (
	"github.com/gaia-charge/go-datastore/pkg/cdr"
	mocks "github.com/gaia-charge/go-datastore/pkg/db/mocks"
)

func NewRepository(repositoryService *mocks.MockRepositoryService) cdr.CdrRepository {
	return cdr.CdrRepository(repositoryService)
}
