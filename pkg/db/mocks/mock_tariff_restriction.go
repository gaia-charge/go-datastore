package mocks

import (
	"context"

	"github.com/gaia-charge/go-datastore/pkg/db"
)

type TariffRestrictionMockData struct {
	TariffRestriction db.TariffRestriction
	Error             error
}

func (r *MockRepositoryService) CreateTariffRestriction(ctx context.Context, arg db.CreateTariffRestrictionParams) (db.TariffRestriction, error) {
	r.createTariffRestrictionMockData = append(r.createTariffRestrictionMockData, arg)
	return db.TariffRestriction{}, nil
}

func (r *MockRepositoryService) DeleteTariffRestriction(ctx context.Context, tariffID int64) error {
	r.deleteTariffRestrictionMockData = append(r.deleteTariffRestrictionMockData, tariffID)
	return nil
}

func (r *MockRepositoryService) GetTariffRestriction(ctx context.Context, id int64) (db.TariffRestriction, error) {
	if len(r.getTariffRestrictionMockData) == 0 {
		return db.TariffRestriction{}, ErrorNotFound()
	}

	response := r.getTariffRestrictionMockData[0]
	r.getTariffRestrictionMockData = r.getTariffRestrictionMockData[1:]
	return response.TariffRestriction, response.Error
}

func (r *MockRepositoryService) UpdateTariffRestriction(ctx context.Context, arg db.UpdateTariffRestrictionParams) (db.TariffRestriction, error) {
	r.updateTariffRestrictionMockData = append(r.updateTariffRestrictionMockData, arg)
	return db.TariffRestriction{}, nil
}

func (r *MockRepositoryService) GetCreateTariffRestrictionMockData() (db.CreateTariffRestrictionParams, error) {
	if len(r.createTariffRestrictionMockData) == 0 {
		return db.CreateTariffRestrictionParams{}, ErrorNotFound()
	}

	response := r.createTariffRestrictionMockData[0]
	r.createTariffRestrictionMockData = r.createTariffRestrictionMockData[1:]
	return response, nil
}

func (r *MockRepositoryService) GetDeleteTariffRestrictionMockData() (int64, error) {
	if len(r.deleteTariffRestrictionMockData) == 0 {
		return 0, ErrorNotFound()
	}

	response := r.deleteTariffRestrictionMockData[0]
	r.deleteTariffRestrictionMockData = r.deleteTariffRestrictionMockData[1:]
	return response, nil
}

func (r *MockRepositoryService) GetUpdateTariffRestrictionMockData() (db.UpdateTariffRestrictionParams, error) {
	if len(r.updateTariffRestrictionMockData) == 0 {
		return db.UpdateTariffRestrictionParams{}, ErrorNotFound()
	}

	response := r.updateTariffRestrictionMockData[0]
	r.updateTariffRestrictionMockData = r.updateTariffRestrictionMockData[1:]
	return response, nil
}

func (r *MockRepositoryService) SetGetTariffRestrictionMockData(response TariffRestrictionMockData) {
	r.getTariffRestrictionMockData = append(r.getTariffRestrictionMockData, response)
}
