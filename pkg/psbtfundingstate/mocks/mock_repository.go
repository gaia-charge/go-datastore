package mocks

import (
	"github.com/gaia-charge/go-datastore/pkg/channelrequest"
	mocks "github.com/gaia-charge/go-datastore/pkg/db/mocks"
)

func NewRepository(repositoryService *mocks.MockRepositoryService) channelrequest.ChannelRequestRepository {
	return channelrequest.ChannelRequestRepository(repositoryService)
}
