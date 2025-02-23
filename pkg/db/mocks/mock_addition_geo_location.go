package mocks

import (
	"context"

	"github.com/gaia-charge/go-datastore/pkg/db"
)

type AdditionalGeoLocationMockData struct {
	AdditionalGeoLocation db.AdditionalGeoLocation
	Error                 error
}

type AdditionalGeoLocationsMockData struct {
	AdditionalGeoLocations []db.AdditionalGeoLocation
	Error                  error
}

func (r *MockRepositoryService) CreateAdditionalGeoLocation(ctx context.Context, arg db.CreateAdditionalGeoLocationParams) (db.AdditionalGeoLocation, error) {
	r.createAdditionalGeoLocationMockData = append(r.createAdditionalGeoLocationMockData, arg)
	return db.AdditionalGeoLocation{}, nil
}

func (r *MockRepositoryService) DeleteAdditionalGeoLocations(ctx context.Context, locationID int64) error {
	r.deleteAdditionalGeoLocationsMockData = append(r.deleteAdditionalGeoLocationsMockData, locationID)
	return nil
}

func (r *MockRepositoryService) ListAdditionalGeoLocations(ctx context.Context, locationID int64) ([]db.AdditionalGeoLocation, error) {
	if len(r.listAdditionalGeoLocationsMockData) == 0 {
		return []db.AdditionalGeoLocation{}, nil
	}

	response := r.listAdditionalGeoLocationsMockData[0]
	r.listAdditionalGeoLocationsMockData = r.listAdditionalGeoLocationsMockData[1:]
	return response.AdditionalGeoLocations, response.Error
}

func (r *MockRepositoryService) GetCreateAdditionalGeoLocationMockData() (db.CreateAdditionalGeoLocationParams, error) {
	if len(r.createAdditionalGeoLocationMockData) == 0 {
		return db.CreateAdditionalGeoLocationParams{}, ErrorNotFound()
	}

	response := r.createAdditionalGeoLocationMockData[0]
	r.createAdditionalGeoLocationMockData = r.createAdditionalGeoLocationMockData[1:]
	return response, nil
}

func (r *MockRepositoryService) GetDeleteAdditionalGeoLocationsMockData() (int64, error) {
	if len(r.deleteAdditionalGeoLocationsMockData) == 0 {
		return 0, ErrorNotFound()
	}

	response := r.deleteAdditionalGeoLocationsMockData[0]
	r.deleteAdditionalGeoLocationsMockData = r.deleteAdditionalGeoLocationsMockData[1:]
	return response, nil
}

func (r *MockRepositoryService) SetListAdditionalGeoLocationsMockData(response AdditionalGeoLocationsMockData) {
	r.listAdditionalGeoLocationsMockData = append(r.listAdditionalGeoLocationsMockData, response)
}
