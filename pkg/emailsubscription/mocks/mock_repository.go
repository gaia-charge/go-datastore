package mocks

import (
	mocks "github.com/gaia-charge/go-datastore/pkg/db/mocks"
	"github.com/gaia-charge/go-datastore/pkg/emailsubscription"
)

func NewRepository(repositoryService *mocks.MockRepositoryService) emailsubscription.EmailSubscriptionRepository {
	return emailsubscription.EmailSubscriptionRepository(repositoryService)
}
