package cdr

import (
	"context"

	"github.com/gaia-charge/go-datastore/pkg/db"
)

type CdrRepository interface {
	CountCdrsByLocationID(ctx context.Context, locationID int64) (int64, error)
	CountCdrsByUserID(ctx context.Context, userID int64) (int64, error)
	CreateCdr(ctx context.Context, arg db.CreateCdrParams) (db.Cdr, error)
	DeleteCdrChargingPeriods(ctx context.Context, cdrID int64) error
	GetCdrByAuthorizationID(ctx context.Context, authorizationID string) (db.Cdr, error)
	GetCdrByLastUpdated(ctx context.Context, arg db.GetCdrByLastUpdatedParams) (db.Cdr, error)
	GetCdrByUid(ctx context.Context, uid string) (db.Cdr, error)
	ListCdrChargingPeriods(ctx context.Context, cdrID int64) ([]db.ChargingPeriod, error)
	ListCdrsByCompletedSessionStatus(ctx context.Context, nodeID int64) ([]db.Cdr, error)
	SetCdrChargingPeriod(ctx context.Context, arg db.SetCdrChargingPeriodParams) error
	UpdateCdrIsFlaggedByUid(ctx context.Context, arg db.UpdateCdrIsFlaggedByUidParams) error
}

func NewRepository(repositoryService *db.RepositoryService) CdrRepository {
	return CdrRepository(repositoryService)
}
