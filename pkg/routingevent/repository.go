package routingevent

import (
	"context"

	"github.com/gaia-charge/go-datastore/pkg/db"
)

type RoutingEventRepository interface {
	CreateRoutingEvent(ctx context.Context, arg db.CreateRoutingEventParams) (db.RoutingEvent, error)
	UpdateRoutingEvent(ctx context.Context, arg db.UpdateRoutingEventParams) (db.RoutingEvent, error)
}

func NewRepository(repositoryService *db.RepositoryService) RoutingEventRepository {
	return RoutingEventRepository(repositoryService)
}
