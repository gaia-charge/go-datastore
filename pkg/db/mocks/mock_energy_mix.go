package mocks

import (
	"context"

	"github.com/gaia-charge/go-datastore/pkg/db"
)

type EnergyMixMockData struct {
	EnergyMix db.EnergyMix
	Error     error
}

func (r *MockRepositoryService) CreateEnergyMix(ctx context.Context, arg db.CreateEnergyMixParams) (db.EnergyMix, error) {
	r.createEnergyMixMockData = append(r.createEnergyMixMockData, arg)
	return db.EnergyMix{}, nil
}

func (r *MockRepositoryService) GetEnergyMix(ctx context.Context, id int64) (db.EnergyMix, error) {
	if len(r.getEnergyMixMockData) == 0 {
		return db.EnergyMix{}, ErrorNotFound()
	}

	response := r.getEnergyMixMockData[0]
	r.getEnergyMixMockData = r.getEnergyMixMockData[1:]
	return response.EnergyMix, response.Error
}

func (r *MockRepositoryService) UpdateEnergyMix(ctx context.Context, arg db.UpdateEnergyMixParams) (db.EnergyMix, error) {
	r.updateEnergyMixMockData = append(r.updateEnergyMixMockData, arg)
	return db.EnergyMix{}, nil
}

func (r *MockRepositoryService) GetCreateEnergyMixMockData() (db.CreateEnergyMixParams, error) {
	if len(r.createEnergyMixMockData) == 0 {
		return db.CreateEnergyMixParams{}, ErrorNotFound()
	}

	response := r.createEnergyMixMockData[0]
	r.createEnergyMixMockData = r.createEnergyMixMockData[1:]
	return response, nil
}

func (r *MockRepositoryService) GetUpdateEnergyMixMockData() (db.UpdateEnergyMixParams, error) {
	if len(r.updateEnergyMixMockData) == 0 {
		return db.UpdateEnergyMixParams{}, ErrorNotFound()
	}

	response := r.updateEnergyMixMockData[0]
	r.updateEnergyMixMockData = r.updateEnergyMixMockData[1:]
	return response, nil
}

func (r *MockRepositoryService) SetGetEnergyMixMockData(response EnergyMixMockData) {
	r.getEnergyMixMockData = append(r.getEnergyMixMockData, response)
}
