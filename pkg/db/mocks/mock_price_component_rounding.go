package mocks

import (
	"context"

	"github.com/gaia-charge/go-datastore/pkg/db"
)

type PriceComponentRoundingMockData struct {
	PriceComponentRounding db.PriceComponentRounding
	Error                  error
}

type PriceComponentRoundingsMockData struct {
	PriceComponentRoundings []db.PriceComponentRounding
	Error                   error
}

func (r *MockRepositoryService) CreatePriceComponentRounding(ctx context.Context, arg db.CreatePriceComponentRoundingParams) (db.PriceComponentRounding, error) {
	r.createPriceComponentRoundingMockData = append(r.createPriceComponentRoundingMockData, arg)
	return db.PriceComponentRounding{}, nil
}

func (r *MockRepositoryService) DeletePriceComponentRoundings(ctx context.Context, tariffID int64) error {
	r.deletePriceComponentRoundingsMockData = append(r.deletePriceComponentRoundingsMockData, tariffID)
	return nil
}

func (r *MockRepositoryService) GetPriceComponentRounding(ctx context.Context, id int64) (db.PriceComponentRounding, error) {
	if len(r.getPriceComponentRoundingMockData) == 0 {
		return db.PriceComponentRounding{}, nil
	}

	response := r.getPriceComponentRoundingMockData[0]
	r.getPriceComponentRoundingMockData = r.getPriceComponentRoundingMockData[1:]
	return response.PriceComponentRounding, response.Error
}

func (r *MockRepositoryService) GetCreatePriceComponentRoundingMockData() (db.CreatePriceComponentRoundingParams, error) {
	if len(r.createPriceComponentRoundingMockData) == 0 {
		return db.CreatePriceComponentRoundingParams{}, ErrorNotFound()
	}

	response := r.createPriceComponentRoundingMockData[0]
	r.createPriceComponentRoundingMockData = r.createPriceComponentRoundingMockData[1:]
	return response, nil
}

func (r *MockRepositoryService) GetDeletePriceComponentRoundingsMockData() (int64, error) {
	if len(r.deletePriceComponentRoundingsMockData) == 0 {
		return 0, ErrorNotFound()
	}

	response := r.deletePriceComponentRoundingsMockData[0]
	r.deletePriceComponentRoundingsMockData = r.deletePriceComponentRoundingsMockData[1:]
	return response, nil
}

func (r *MockRepositoryService) SetGetPriceComponentRoundingMockData(response PriceComponentRoundingMockData) {
	r.getPriceComponentRoundingMockData = append(r.getPriceComponentRoundingMockData, response)
}
