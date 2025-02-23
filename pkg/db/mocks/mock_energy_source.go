package mocks

import (
	"context"

	"github.com/gaia-charge/go-datastore/pkg/db"
)

type EnergySourceMockData struct {
	EnergySource db.EnergySource
	Error        error
}

type EnergySourcesMockData struct {
	EnergySources []db.EnergySource
	Error         error
}

func (r *MockRepositoryService) CreateEnergySource(ctx context.Context, arg db.CreateEnergySourceParams) (db.EnergySource, error) {
	r.createEnergySourceMockData = append(r.createEnergySourceMockData, arg)
	return db.EnergySource{}, nil
}

func (r *MockRepositoryService) DeleteEnergySources(ctx context.Context, evseID int64) error {
	r.deleteEnergySourcesMockData = append(r.deleteEnergySourcesMockData, evseID)
	return nil
}

func (r *MockRepositoryService) ListEnergySources(ctx context.Context, id int64) ([]db.EnergySource, error) {
	if len(r.listEnergySourcesMockData) == 0 {
		return []db.EnergySource{}, nil
	}

	response := r.listEnergySourcesMockData[0]
	r.listEnergySourcesMockData = r.listEnergySourcesMockData[1:]
	return response.EnergySources, response.Error
}

func (r *MockRepositoryService) GetCreateEnergySourceMockData() (db.CreateEnergySourceParams, error) {
	if len(r.createEnergySourceMockData) == 0 {
		return db.CreateEnergySourceParams{}, ErrorNotFound()
	}

	response := r.createEnergySourceMockData[0]
	r.createEnergySourceMockData = r.createEnergySourceMockData[1:]
	return response, nil
}

func (r *MockRepositoryService) GetDeleteEnergySourcesMockData() (int64, error) {
	if len(r.deleteEnergySourcesMockData) == 0 {
		return 0, ErrorNotFound()
	}

	response := r.deleteEnergySourcesMockData[0]
	r.deleteEnergySourcesMockData = r.deleteEnergySourcesMockData[1:]
	return response, nil
}

func (r *MockRepositoryService) SetListEnergySourcesMockData(response EnergySourcesMockData) {
	r.listEnergySourcesMockData = append(r.listEnergySourcesMockData, response)
}
