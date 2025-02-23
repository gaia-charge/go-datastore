package command

import (
	"context"

	"github.com/gaia-charge/go-datastore/pkg/db"
)

type CommandRepository interface {
	CreateCommandReservation(ctx context.Context, arg db.CreateCommandReservationParams) (db.CommandReservation, error)
	CreateCommandStart(ctx context.Context, arg db.CreateCommandStartParams) (db.CommandStart, error)
	CreateCommandStop(ctx context.Context, arg db.CreateCommandStopParams) (db.CommandStop, error)
	CreateCommandUnlock(ctx context.Context, arg db.CreateCommandUnlockParams) (db.CommandUnlock, error)
	GetCommandReservation(ctx context.Context, id int64) (db.CommandReservation, error)
	GetCommandStart(ctx context.Context, id int64) (db.CommandStart, error)
	GetCommandStop(ctx context.Context, id int64) (db.CommandStop, error)
	GetCommandUnlock(ctx context.Context, id int64) (db.CommandUnlock, error)
	UpdateCommandReservation(ctx context.Context, arg db.UpdateCommandReservationParams) (db.CommandReservation, error)
	UpdateCommandStart(ctx context.Context, arg db.UpdateCommandStartParams) (db.CommandStart, error)
	UpdateCommandStartByAuthorizationID(ctx context.Context, arg db.UpdateCommandStartByAuthorizationIDParams) (db.CommandStart, error)
	UpdateCommandStop(ctx context.Context, arg db.UpdateCommandStopParams) (db.CommandStop, error)
	UpdateCommandStopBySessionID(ctx context.Context, arg db.UpdateCommandStopBySessionIDParams) (db.CommandStop, error)
	UpdateCommandUnlock(ctx context.Context, arg db.UpdateCommandUnlockParams) (db.CommandUnlock, error)
}

func NewRepository(repositoryService *db.RepositoryService) CommandRepository {
	return CommandRepository(repositoryService)
}
