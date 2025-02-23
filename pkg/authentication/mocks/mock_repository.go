package mocks

import (
	"github.com/gaia-charge/go-datastore/pkg/authentication"
	mocks "github.com/gaia-charge/go-datastore/pkg/db/mocks"
)

func NewRepository(repositoryService *mocks.MockRepositoryService) authentication.AuthenticationRepository {
	return authentication.AuthenticationRepository(repositoryService)
}
