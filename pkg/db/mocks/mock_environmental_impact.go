package mocks

import (
	"context"

	"github.com/gaia-charge/go-datastore/pkg/db"
)

type EnvironmentalImpactMockData struct {
	EnvironmentalImpact db.EnvironmentalImpact
	Error               error
}

type EnvironmentalImpactsMockData struct {
	EnvironmentalImpacts []db.EnvironmentalImpact
	Error                error
}

func (r *MockRepositoryService) CreateEnvironmentalImpact(ctx context.Context, arg db.CreateEnvironmentalImpactParams) (db.EnvironmentalImpact, error) {
	r.createEnvironmentalImpactMockData = append(r.createEnvironmentalImpactMockData, arg)
	return db.EnvironmentalImpact{}, nil
}

func (r *MockRepositoryService) DeleteEnvironmentalImpacts(ctx context.Context, evseID int64) error {
	r.deleteEnvironmentalImpactsMockData = append(r.deleteEnvironmentalImpactsMockData, evseID)
	return nil
}

func (r *MockRepositoryService) ListEnvironmentalImpacts(ctx context.Context, id int64) ([]db.EnvironmentalImpact, error) {
	if len(r.listEnvironmentalImpactsMockData) == 0 {
		return []db.EnvironmentalImpact{}, nil
	}

	response := r.listEnvironmentalImpactsMockData[0]
	r.listEnvironmentalImpactsMockData = r.listEnvironmentalImpactsMockData[1:]
	return response.EnvironmentalImpacts, response.Error
}

func (r *MockRepositoryService) GetCreateEnvironmentalImpactMockData() (db.CreateEnvironmentalImpactParams, error) {
	if len(r.createEnvironmentalImpactMockData) == 0 {
		return db.CreateEnvironmentalImpactParams{}, ErrorNotFound()
	}

	response := r.createEnvironmentalImpactMockData[0]
	r.createEnvironmentalImpactMockData = r.createEnvironmentalImpactMockData[1:]
	return response, nil
}

func (r *MockRepositoryService) GetDeleteEnvironmentalImpactsMockData() (int64, error) {
	if len(r.deleteEnvironmentalImpactsMockData) == 0 {
		return 0, ErrorNotFound()
	}

	response := r.deleteEnvironmentalImpactsMockData[0]
	r.deleteEnvironmentalImpactsMockData = r.deleteEnvironmentalImpactsMockData[1:]
	return response, nil
}

func (r *MockRepositoryService) SetListEnvironmentalImpactsMockData(response EnvironmentalImpactsMockData) {
	r.listEnvironmentalImpactsMockData = append(r.listEnvironmentalImpactsMockData, response)
}
