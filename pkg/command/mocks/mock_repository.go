package mocks

import (
	"github.com/gaia-charge/go-datastore/pkg/command"
	mocks "github.com/gaia-charge/go-datastore/pkg/db/mocks"
)

func NewRepository(repositoryService *mocks.MockRepositoryService) command.CommandRepository {
	return command.CommandRepository(repositoryService)
}
