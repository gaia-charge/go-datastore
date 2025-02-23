package mocks

import (
	"context"

	"github.com/gaia-charge/go-datastore/pkg/db"
)

func (r *MockRepositoryService) SetEvseParkingRestriction(ctx context.Context, arg db.SetEvseParkingRestrictionParams) error {
	r.setEvseParkingRestrictionMockData = append(r.setEvseParkingRestrictionMockData, arg)
	return nil
}

func (r *MockRepositoryService) ListEvseParkingRestrictions(ctx context.Context, evseID int64) ([]db.ParkingRestriction, error) {
	if len(r.listEvseParkingRestrictionsMockData) == 0 {
		return []db.ParkingRestriction{}, nil
	}

	response := r.listEvseParkingRestrictionsMockData[0]
	r.listEvseParkingRestrictionsMockData = r.listEvseParkingRestrictionsMockData[1:]
	return response.ParkingRestrictions, response.Error
}

func (r *MockRepositoryService) UnsetEvseParkingRestrictions(ctx context.Context, evseID int64) error {
	r.unsetEvseParkingRestrictionsMockData = append(r.unsetEvseParkingRestrictionsMockData, evseID)
	return nil
}

func (r *MockRepositoryService) GetSetEvseParkingRestrictionMockData() (db.SetEvseParkingRestrictionParams, error) {
	if len(r.setEvseParkingRestrictionMockData) == 0 {
		return db.SetEvseParkingRestrictionParams{}, ErrorNotFound()
	}

	response := r.setEvseParkingRestrictionMockData[0]
	r.setEvseParkingRestrictionMockData = r.setEvseParkingRestrictionMockData[1:]
	return response, nil
}

func (r *MockRepositoryService) SetListEvseParkingRestrictionsMockData(response ParkingRestrictionsMockData) {
	r.listEvseParkingRestrictionsMockData = append(r.listEvseParkingRestrictionsMockData, response)
}

func (r *MockRepositoryService) GetUnsetEvseParkingRestrictionsMockData() (int64, error) {
	if len(r.unsetEvseParkingRestrictionsMockData) == 0 {
		return 0, ErrorNotFound()
	}

	response := r.unsetEvseParkingRestrictionsMockData[0]
	r.unsetEvseParkingRestrictionsMockData = r.unsetEvseParkingRestrictionsMockData[1:]
	return response, nil
}
