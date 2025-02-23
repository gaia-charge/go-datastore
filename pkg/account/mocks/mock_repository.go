package mocks

import (
	"github.com/gaia-charge/go-datastore/pkg/account"
	mocks "github.com/gaia-charge/go-datastore/pkg/db/mocks"
)

func NewRepository(repositoryService *mocks.MockRepositoryService) account.AccountRepository {
	return account.AccountRepository(repositoryService)
}
