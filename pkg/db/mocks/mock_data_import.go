package mocks

import (
	"context"

	"github.com/gaia-charge/go-datastore/pkg/db"
)

type HtbTariffMockData struct {
	HtbTariff db.HtbTariff
	Error     error
}

type HtbTariffsMockData struct {
	HtbTariffs []db.HtbTariff
	Error      error
}

func (r *MockRepositoryService) GetHtbTariffByName(ctx context.Context, name string) (db.HtbTariff, error) {
	if len(r.getHtbTariffByNameMockData) == 0 {
		return db.HtbTariff{}, ErrorNotFound()
	}

	response := r.getHtbTariffByNameMockData[0]
	r.getHtbTariffByNameMockData = r.getHtbTariffByNameMockData[1:]
	return response.HtbTariff, response.Error
}

func (r *MockRepositoryService) ListHtbTariffs(ctx context.Context) ([]db.HtbTariff, error) {
	if len(r.listHtbTariffsMockData) == 0 {
		return []db.HtbTariff{}, nil
	}

	response := r.listHtbTariffsMockData[0]
	r.listHtbTariffsMockData = r.listHtbTariffsMockData[1:]
	return response.HtbTariffs, response.Error
}

func (r *MockRepositoryService) SetGetHtbTariffByNameMockData(response HtbTariffMockData) {
	r.getHtbTariffByNameMockData = append(r.getHtbTariffByNameMockData, response)
}

func (r *MockRepositoryService) SetListHtbTariffsMockData(response HtbTariffsMockData) {
	r.listHtbTariffsMockData = append(r.listHtbTariffsMockData, response)
}
