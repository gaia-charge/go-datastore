package mocks

import (
	"context"
	"time"

	"github.com/gaia-charge/go-datastore/pkg/db"
)

type ChargingPeriodMockData struct {
	ChargingPeriod db.ChargingPeriod
	Error          error
}

type ChargingPeriodsMockData struct {
	ChargingPeriods []db.ChargingPeriod
	Error           error
}

func (r *MockRepositoryService) CreateChargingPeriod(ctx context.Context, startDateTime time.Time) (db.ChargingPeriod, error) {
	r.createChargingPeriodMockData = append(r.createChargingPeriodMockData, startDateTime)
	return db.ChargingPeriod{}, nil
}

func (r *MockRepositoryService) GetCreateChargingPeriodMockData() (time.Time, error) {
	if len(r.createChargingPeriodMockData) == 0 {
		return time.Time{}, ErrorNotFound()
	}

	response := r.createChargingPeriodMockData[0]
	r.createChargingPeriodMockData = r.createChargingPeriodMockData[1:]
	return response, nil
}
