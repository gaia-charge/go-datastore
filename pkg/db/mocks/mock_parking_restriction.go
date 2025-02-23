package mocks

import (
	"context"

	"github.com/gaia-charge/go-datastore/pkg/db"
)

type ParkingRestrictionsMockData struct {
	ParkingRestrictions []db.ParkingRestriction
	Error               error
}

func (r *MockRepositoryService) ListParkingRestrictions(ctx context.Context) ([]db.ParkingRestriction, error) {
	if len(r.listParkingRestrictionsMockData) == 0 {
		return []db.ParkingRestriction{}, nil
	}

	response := r.listParkingRestrictionsMockData[0]
	r.listParkingRestrictionsMockData = r.listParkingRestrictionsMockData[1:]
	return response.ParkingRestrictions, response.Error
}

func (r *MockRepositoryService) SetListParkingRestrictionsMockData(response ParkingRestrictionsMockData) {
	r.listParkingRestrictionsMockData = append(r.listParkingRestrictionsMockData, response)
}
