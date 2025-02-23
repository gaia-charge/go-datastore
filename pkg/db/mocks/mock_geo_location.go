package mocks

import (
	"context"

	"github.com/gaia-charge/go-datastore/pkg/db"
)

type GeoLocationMockData struct {
	GeoLocation db.GeoLocation
	Error       error
}

type GeoLocationsMockData struct {
	GeoLocations []db.GeoLocation
	Error        error
}

func (r *MockRepositoryService) CreateGeoLocation(ctx context.Context, arg db.CreateGeoLocationParams) (db.GeoLocation, error) {
	r.createGeoLocationMockData = append(r.createGeoLocationMockData, arg)
	return db.GeoLocation{
		ID:        1,
		Latitude:  arg.Latitude,
		Longitude: arg.Longitude,
	}, nil
}

func (r *MockRepositoryService) DeleteGeoLocation(ctx context.Context, id int64) error {
	r.deleteGeoLocationMockData = append(r.deleteGeoLocationMockData, id)
	return nil
}

func (r *MockRepositoryService) GetGeoLocation(ctx context.Context, id int64) (db.GeoLocation, error) {
	if len(r.getGeoLocationMockData) == 0 {
		return db.GeoLocation{}, ErrorNotFound()
	}

	response := r.getGeoLocationMockData[0]
	r.getGeoLocationMockData = r.getGeoLocationMockData[1:]
	return response.GeoLocation, response.Error
}

func (r *MockRepositoryService) UpdateGeoLocation(ctx context.Context, arg db.UpdateGeoLocationParams) (db.GeoLocation, error) {
	r.updateGeoLocationMockData = append(r.updateGeoLocationMockData, arg)
	return db.GeoLocation{
		ID:        arg.ID,
		Latitude:  arg.Latitude,
		Longitude: arg.Longitude,
	}, nil
}

func (r *MockRepositoryService) GetCreateGeoLocationMockData() (db.CreateGeoLocationParams, error) {
	if len(r.createGeoLocationMockData) == 0 {
		return db.CreateGeoLocationParams{}, ErrorNotFound()
	}

	response := r.createGeoLocationMockData[0]
	r.createGeoLocationMockData = r.createGeoLocationMockData[1:]
	return response, nil
}

func (r *MockRepositoryService) GetDeleteGeoLocationMockData() (int64, error) {
	if len(r.deleteGeoLocationMockData) == 0 {
		return 0, ErrorNotFound()
	}

	response := r.deleteGeoLocationMockData[0]
	r.deleteGeoLocationMockData = r.deleteGeoLocationMockData[1:]
	return response, nil
}

func (r *MockRepositoryService) GetUpdateGeoLocationMockData() (db.UpdateGeoLocationParams, error) {
	if len(r.updateGeoLocationMockData) == 0 {
		return db.UpdateGeoLocationParams{}, ErrorNotFound()
	}

	response := r.updateGeoLocationMockData[0]
	r.updateGeoLocationMockData = r.updateGeoLocationMockData[1:]
	return response, nil
}

func (r *MockRepositoryService) SetGetGeoLocationMockData(response GeoLocationMockData) {
	r.getGeoLocationMockData = append(r.getGeoLocationMockData, response)
}
