package mocks

import (
	"context"

	"github.com/gaia-charge/go-datastore/pkg/db"
)

func (r *MockRepositoryService) SetCdrChargingPeriod(ctx context.Context, arg db.SetCdrChargingPeriodParams) error {
	r.setCdrChargingPeriodMockData = append(r.setCdrChargingPeriodMockData, arg)
	return nil
}

func (r *MockRepositoryService) ListCdrChargingPeriods(ctx context.Context, sessionID int64) ([]db.ChargingPeriod, error) {
	if len(r.listCdrChargingPeriodsMockData) == 0 {
		return []db.ChargingPeriod{}, nil
	}

	response := r.listCdrChargingPeriodsMockData[0]
	r.listCdrChargingPeriodsMockData = r.listCdrChargingPeriodsMockData[1:]
	return response.ChargingPeriods, response.Error
}

func (r *MockRepositoryService) DeleteCdrChargingPeriods(ctx context.Context, sessionID int64) error {
	r.deleteCdrChargingPeriodsMockData = append(r.deleteCdrChargingPeriodsMockData, sessionID)
	return nil
}

func (r *MockRepositoryService) GetSetCdrChargingPeriodMockData() (db.SetCdrChargingPeriodParams, error) {
	if len(r.setCdrChargingPeriodMockData) == 0 {
		return db.SetCdrChargingPeriodParams{}, ErrorNotFound()
	}

	response := r.setCdrChargingPeriodMockData[0]
	r.setCdrChargingPeriodMockData = r.setCdrChargingPeriodMockData[1:]
	return response, nil
}

func (r *MockRepositoryService) SetListCdrChargingPeriodsMockData(response ChargingPeriodsMockData) {
	r.listCdrChargingPeriodsMockData = append(r.listCdrChargingPeriodsMockData, response)
}

func (r *MockRepositoryService) GetDeleteCdrChargingPeriodsMockData() (int64, error) {
	if len(r.deleteCdrChargingPeriodsMockData) == 0 {
		return 0, ErrorNotFound()
	}

	response := r.deleteCdrChargingPeriodsMockData[0]
	r.deleteCdrChargingPeriodsMockData = r.deleteCdrChargingPeriodsMockData[1:]
	return response, nil
}
