package mocks

import (
	mocks "github.com/gaia-charge/go-datastore/pkg/db/mocks"
	"github.com/gaia-charge/go-datastore/pkg/routingevent"
)

func NewRepository(repositoryService *mocks.MockRepositoryService) routingevent.RoutingEventRepository {
	return routingevent.RoutingEventRepository(repositoryService)
}
