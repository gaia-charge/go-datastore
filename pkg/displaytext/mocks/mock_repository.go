package mocks

import (
	mocks "github.com/gaia-charge/go-datastore/pkg/db/mocks"
	"github.com/gaia-charge/go-datastore/pkg/displaytext"
)

func NewRepository(repositoryService *mocks.MockRepositoryService) displaytext.DisplayTextRepository {
	return displaytext.DisplayTextRepository(repositoryService)
}
