package mocks

import (
	mocks "github.com/gaia-charge/go-datastore/pkg/db/mocks"
	"github.com/gaia-charge/go-datastore/pkg/element"
)

func NewRepository(repositoryService *mocks.MockRepositoryService) element.ElementRepository {
	return element.ElementRepository(repositoryService)
}
