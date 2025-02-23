package mocks

import (
	mocks "github.com/gaia-charge/go-datastore/pkg/db/mocks"
	"github.com/gaia-charge/go-datastore/pkg/user"
)

func NewRepository(repositoryService *mocks.MockRepositoryService) user.UserRepository {
	return user.UserRepository(repositoryService)
}
