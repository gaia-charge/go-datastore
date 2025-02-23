package mocks

import (
	"context"

	"github.com/gaia-charge/go-datastore/pkg/db"
)

func (r *MockRepositoryService) ListLocationFacilities(ctx context.Context, locationID int64) ([]db.Facility, error) {
	if len(r.listLocationFacilitiesMockData) == 0 {
		return []db.Facility{}, nil
	}

	response := r.listLocationFacilitiesMockData[0]
	r.listLocationFacilitiesMockData = r.listLocationFacilitiesMockData[1:]
	return response.Facilities, response.Error
}

func (r *MockRepositoryService) SetLocationFacility(ctx context.Context, arg db.SetLocationFacilityParams) error {
	r.setLocationFacilityMockData = append(r.setLocationFacilityMockData, arg)
	return nil
}

func (r *MockRepositoryService) UnsetLocationFacilities(ctx context.Context, evseID int64) error {
	r.unsetLocationFacilitiesMockData = append(r.unsetLocationFacilitiesMockData, evseID)
	return nil
}

func (r *MockRepositoryService) GetSetLocationFacilityMockData() (db.SetLocationFacilityParams, error) {
	if len(r.setLocationFacilityMockData) == 0 {
		return db.SetLocationFacilityParams{}, ErrorNotFound()
	}

	response := r.setLocationFacilityMockData[0]
	r.setLocationFacilityMockData = r.setLocationFacilityMockData[1:]
	return response, nil
}

func (r *MockRepositoryService) SetListLocationFacilitiesMockData(response FacilitiesMockData) {
	r.listLocationFacilitiesMockData = append(r.listLocationFacilitiesMockData, response)
}

func (r *MockRepositoryService) GetUnsetLocationFacilitiesMockData() (int64, error) {
	if len(r.unsetLocationFacilitiesMockData) == 0 {
		return 0, ErrorNotFound()
	}

	response := r.unsetLocationFacilitiesMockData[0]
	r.unsetLocationFacilitiesMockData = r.unsetLocationFacilitiesMockData[1:]
	return response, nil
}
