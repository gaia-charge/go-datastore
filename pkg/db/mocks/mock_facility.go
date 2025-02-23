package mocks

import (
	"context"

	"github.com/gaia-charge/go-datastore/pkg/db"
)

type FacilitiesMockData struct {
	Facilities []db.Facility
	Error      error
}

func (r *MockRepositoryService) ListFacilities(ctx context.Context) ([]db.Facility, error) {
	if len(r.listFacilitiesMockData) == 0 {
		return []db.Facility{}, nil
	}

	response := r.listFacilitiesMockData[0]
	r.listFacilitiesMockData = r.listFacilitiesMockData[1:]
	return response.Facilities, response.Error
}

func (r *MockRepositoryService) SetListFacilitiesMockData(response FacilitiesMockData) {
	r.listFacilitiesMockData = append(r.listFacilitiesMockData, response)
}
