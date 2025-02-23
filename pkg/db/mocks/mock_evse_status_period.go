package mocks

import (
	"context"

	"github.com/gaia-charge/go-datastore/pkg/db"
)

type EvseStatusPeriodMockData struct {
	EvseStatusPeriod db.EvseStatusPeriod
	Error            error
}

type EvseStatusPeriodsMockData struct {
	EvseStatusPeriods []db.EvseStatusPeriod
	Error             error
}

func (r *MockRepositoryService) CreateEvseStatusPeriod(ctx context.Context, arg db.CreateEvseStatusPeriodParams) (db.EvseStatusPeriod, error) {
	r.createEvseStatusPeriodMockData = append(r.createEvseStatusPeriodMockData, arg)
	return db.EvseStatusPeriod{}, nil
}

func (r *MockRepositoryService) ListEvseStatusPeriods(ctx context.Context, evseID int64) ([]db.EvseStatusPeriod, error) {
	if len(r.listEvseStatusPeriodsMockData) == 0 {
		return []db.EvseStatusPeriod{}, nil
	}

	response := r.listEvseStatusPeriodsMockData[0]
	r.listEvseStatusPeriodsMockData = r.listEvseStatusPeriodsMockData[1:]
	return response.EvseStatusPeriods, response.Error
}

func (r *MockRepositoryService) GetCreateEvseStatusPeriodMockData() (db.CreateEvseStatusPeriodParams, error) {
	if len(r.createEvseStatusPeriodMockData) == 0 {
		return db.CreateEvseStatusPeriodParams{}, ErrorNotFound()
	}

	response := r.createEvseStatusPeriodMockData[0]
	r.createEvseStatusPeriodMockData = r.createEvseStatusPeriodMockData[1:]
	return response, nil
}

func (r *MockRepositoryService) SetListEvseStatusPeriodsMockData(response EvseStatusPeriodsMockData) {
	r.listEvseStatusPeriodsMockData = append(r.listEvseStatusPeriodsMockData, response)
}
