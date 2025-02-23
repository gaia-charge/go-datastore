package mocks

import (
	mocks "github.com/gaia-charge/go-datastore/pkg/db/mocks"
	"github.com/gaia-charge/go-datastore/pkg/pendingnotification"
)

func NewRepository(repositoryService *mocks.MockRepositoryService) pendingnotification.PendingNotificationRepository {
	return pendingnotification.PendingNotificationRepository(repositoryService)
}
