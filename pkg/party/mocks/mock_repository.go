package mocks

import (
	mocks "github.com/gaia-charge/go-datastore/pkg/db/mocks"
	"github.com/gaia-charge/go-datastore/pkg/party"
)

func NewRepository(repositoryService *mocks.MockRepositoryService) party.PartyRepository {
	return party.PartyRepository(repositoryService)
}
