package chargingperiod

import (
	"context"
	"time"

	"github.com/gaia-charge/go-datastore/pkg/db"
)

type ChargingPeriodRepository interface {
	CreateChargingPeriod(ctx context.Context, startDateTime time.Time) (db.ChargingPeriod, error)
	CreateChargingPeriodDimension(ctx context.Context, arg db.CreateChargingPeriodDimensionParams) (db.ChargingPeriodDimension, error)
	DeleteChargingPeriodDimensions(ctx context.Context, chargingPeriodID int64) error
	ListChargingPeriodDimensions(ctx context.Context, chargingPeriodID int64) ([]db.ChargingPeriodDimension, error)
}

func NewRepository(repositoryService *db.RepositoryService) ChargingPeriodRepository {
	return ChargingPeriodRepository(repositoryService)
}
