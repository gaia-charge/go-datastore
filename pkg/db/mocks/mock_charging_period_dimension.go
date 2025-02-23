package mocks

import (
	"context"

	"github.com/gaia-charge/go-datastore/pkg/db"
)

type ChargingPeriodDimensionMockData struct {
	ChargingPeriodDimension db.ChargingPeriodDimension
	Error                   error
}

type ChargingPeriodDimensionsMockData struct {
	ChargingPeriodDimensions []db.ChargingPeriodDimension
	Error                    error
}

func (r *MockRepositoryService) CreateChargingPeriodDimension(ctx context.Context, arg db.CreateChargingPeriodDimensionParams) (db.ChargingPeriodDimension, error) {
	r.createChargingPeriodDimensionMockData = append(r.createChargingPeriodDimensionMockData, arg)
	return db.ChargingPeriodDimension{}, nil
}

func (r *MockRepositoryService) ListChargingPeriodDimensions(ctx context.Context, chargingPeriodID int64) ([]db.ChargingPeriodDimension, error) {
	if len(r.listChargingPeriodDimensionsMockData) == 0 {
		return []db.ChargingPeriodDimension{}, nil
	}

	response := r.listChargingPeriodDimensionsMockData[0]
	r.listChargingPeriodDimensionsMockData = r.listChargingPeriodDimensionsMockData[1:]
	return response.ChargingPeriodDimensions, response.Error
}

func (r *MockRepositoryService) DeleteChargingPeriodDimensions(ctx context.Context, chargingPeriodID int64) error {
	r.deleteChargingPeriodDimensionsMockData = append(r.deleteChargingPeriodDimensionsMockData, chargingPeriodID)
	return nil
}

func (r *MockRepositoryService) GetCreateChargingPeriodDimensionMockData() (db.CreateChargingPeriodDimensionParams, error) {
	if len(r.createChargingPeriodDimensionMockData) == 0 {
		return db.CreateChargingPeriodDimensionParams{}, ErrorNotFound()
	}

	response := r.createChargingPeriodDimensionMockData[0]
	r.createChargingPeriodDimensionMockData = r.createChargingPeriodDimensionMockData[1:]
	return response, nil
}

func (r *MockRepositoryService) SetListChargingPeriodDimensionsMockData(response ChargingPeriodDimensionsMockData) {
	r.listChargingPeriodDimensionsMockData = append(r.listChargingPeriodDimensionsMockData, response)
}

func (r *MockRepositoryService) GetDeleteChargingPeriodDimensionsMockData() (int64, error) {
	if len(r.deleteChargingPeriodDimensionsMockData) == 0 {
		return 0, ErrorNotFound()
	}

	response := r.deleteChargingPeriodDimensionsMockData[0]
	r.deleteChargingPeriodDimensionsMockData = r.deleteChargingPeriodDimensionsMockData[1:]
	return response, nil
}
