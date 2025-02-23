package mocks

import (
	"context"

	"github.com/gaia-charge/go-datastore/pkg/db"
)

type PriceComponentMockData struct {
	PriceComponent db.PriceComponent
	Error          error
}

type PriceComponentsMockData struct {
	PriceComponents []db.PriceComponent
	Error           error
}

func (r *MockRepositoryService) CreatePriceComponent(ctx context.Context, arg db.CreatePriceComponentParams) (db.PriceComponent, error) {
	r.createPriceComponentMockData = append(r.createPriceComponentMockData, arg)
	return db.PriceComponent{}, nil
}

func (r *MockRepositoryService) DeletePriceComponents(ctx context.Context, tariffID int64) error {
	r.deletePriceComponentsMockData = append(r.deletePriceComponentsMockData, tariffID)
	return nil
}

func (r *MockRepositoryService) ListPriceComponents(ctx context.Context, elementID int64) ([]db.PriceComponent, error) {
	if len(r.listPriceComponentsMockData) == 0 {
		return []db.PriceComponent{}, nil
	}

	response := r.listPriceComponentsMockData[0]
	r.listPriceComponentsMockData = r.listPriceComponentsMockData[1:]
	return response.PriceComponents, response.Error
}

func (r *MockRepositoryService) GetCreatePriceComponentMockData() (db.CreatePriceComponentParams, error) {
	if len(r.createPriceComponentMockData) == 0 {
		return db.CreatePriceComponentParams{}, ErrorNotFound()
	}

	response := r.createPriceComponentMockData[0]
	r.createPriceComponentMockData = r.createPriceComponentMockData[1:]
	return response, nil
}

func (r *MockRepositoryService) GetDeletePriceComponentsMockData() (int64, error) {
	if len(r.deletePriceComponentsMockData) == 0 {
		return 0, ErrorNotFound()
	}

	response := r.deletePriceComponentsMockData[0]
	r.deletePriceComponentsMockData = r.deletePriceComponentsMockData[1:]
	return response, nil
}

func (r *MockRepositoryService) SetListPriceComponentsMockData(response PriceComponentsMockData) {
	r.listPriceComponentsMockData = append(r.listPriceComponentsMockData, response)
}
