package mocks

import (
	mocks "github.com/gaia-charge/go-datastore/pkg/db/mocks"
	"github.com/gaia-charge/go-datastore/pkg/tokenauthorization"
)

func NewRepository(repositoryService *mocks.MockRepositoryService) tokenauthorization.TokenAuthorizationRepository {
	return tokenauthorization.TokenAuthorizationRepository(repositoryService)
}
