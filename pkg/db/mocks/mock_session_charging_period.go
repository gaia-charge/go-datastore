package mocks

import (
	"context"

	"github.com/gaia-charge/go-datastore/pkg/db"
)

func (r *MockRepositoryService) SetSessionChargingPeriod(ctx context.Context, arg db.SetSessionChargingPeriodParams) error {
	r.setSessionChargingPeriodMockData = append(r.setSessionChargingPeriodMockData, arg)
	return nil
}

func (r *MockRepositoryService) ListSessionChargingPeriods(ctx context.Context, sessionID int64) ([]db.ChargingPeriod, error) {
	if len(r.listSessionChargingPeriodsMockData) == 0 {
		return []db.ChargingPeriod{}, nil
	}

	response := r.listSessionChargingPeriodsMockData[0]
	r.listSessionChargingPeriodsMockData = r.listSessionChargingPeriodsMockData[1:]
	return response.ChargingPeriods, response.Error
}

func (r *MockRepositoryService) DeleteSessionChargingPeriods(ctx context.Context, sessionID int64) error {
	r.deleteSessionChargingPeriodsMockData = append(r.deleteSessionChargingPeriodsMockData, sessionID)
	return nil
}

func (r *MockRepositoryService) GetSetSessionChargingPeriodMockData() (db.SetSessionChargingPeriodParams, error) {
	if len(r.setSessionChargingPeriodMockData) == 0 {
		return db.SetSessionChargingPeriodParams{}, ErrorNotFound()
	}

	response := r.setSessionChargingPeriodMockData[0]
	r.setSessionChargingPeriodMockData = r.setSessionChargingPeriodMockData[1:]
	return response, nil
}

func (r *MockRepositoryService) SetListSessionChargingPeriodsMockData(response ChargingPeriodsMockData) {
	r.listSessionChargingPeriodsMockData = append(r.listSessionChargingPeriodsMockData, response)
}

func (r *MockRepositoryService) GetDeleteSessionChargingPeriodsMockData() (int64, error) {
	if len(r.deleteSessionChargingPeriodsMockData) == 0 {
		return 0, ErrorNotFound()
	}

	response := r.deleteSessionChargingPeriodsMockData[0]
	r.deleteSessionChargingPeriodsMockData = r.deleteSessionChargingPeriodsMockData[1:]
	return response, nil
}
