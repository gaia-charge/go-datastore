package mocks

import (
	"context"

	"github.com/gaia-charge/go-datastore/pkg/db"
)

type LocationMockData struct {
	Location db.Location
	Error    error
}

type LocationsMockData struct {
	Locations []db.Location
	Error     error
}

func (r *MockRepositoryService) CreateLocation(ctx context.Context, arg db.CreateLocationParams) (db.Location, error) {
	r.createLocationMockData = append(r.createLocationMockData, arg)
	return db.Location{}, nil
}

func (r *MockRepositoryService) GetLocation(ctx context.Context, id int64) (db.Location, error) {
	if len(r.getLocationMockData) == 0 {
		return db.Location{}, ErrorNotFound()
	}

	response := r.getLocationMockData[0]
	r.getLocationMockData = r.getLocationMockData[1:]
	return response.Location, response.Error
}

func (r *MockRepositoryService) GetLocationByLastUpdated(ctx context.Context, arg db.GetLocationByLastUpdatedParams) (db.Location, error) {
	if len(r.getLocationByLastUpdatedMockData) == 0 {
		return db.Location{}, ErrorNotFound()
	}

	response := r.getLocationByLastUpdatedMockData[0]
	r.getLocationByLastUpdatedMockData = r.getLocationByLastUpdatedMockData[1:]
	return response.Location, response.Error
}

func (r *MockRepositoryService) GetLocationByUid(ctx context.Context, uid string) (db.Location, error) {
	if len(r.getLocationByUidMockData) == 0 {
		return db.Location{}, ErrorNotFound()
	}

	response := r.getLocationByUidMockData[0]
	r.getLocationByUidMockData = r.getLocationByUidMockData[1:]
	return response.Location, response.Error
}

func (r *MockRepositoryService) ListLocations(ctx context.Context) ([]db.Location, error) {
	if len(r.listLocationsMockData) == 0 {
		return []db.Location{}, nil
	}

	response := r.listLocationsMockData[0]
	r.listLocationsMockData = r.listLocationsMockData[1:]
	return response.Locations, response.Error
}

func (r *MockRepositoryService) ListLocationsByCountry(ctx context.Context, country string) ([]db.Location, error) {
	if len(r.listLocationsByCountryMockData) == 0 {
		return []db.Location{}, nil
	}

	response := r.listLocationsByCountryMockData[0]
	r.listLocationsByCountryMockData = r.listLocationsByCountryMockData[1:]
	return response.Locations, response.Error
}

func (r *MockRepositoryService) ListLocationsByGeom(ctx context.Context, arg db.ListLocationsByGeomParams) ([]db.Location, error) {
	if len(r.listLocationsByGeomMockData) == 0 {
		return []db.Location{}, nil
	}

	response := r.listLocationsByGeomMockData[0]
	r.listLocationsByGeomMockData = r.listLocationsByGeomMockData[1:]
	return response.Locations, response.Error
}

func (r *MockRepositoryService) UpdateLocationByUid(ctx context.Context, arg db.UpdateLocationByUidParams) (db.Location, error) {
	r.updateLocationByUidMockData = append(r.updateLocationByUidMockData, arg)
	return db.Location{}, nil
}

func (r *MockRepositoryService) UpdateLocationLastUpdated(ctx context.Context, arg db.UpdateLocationLastUpdatedParams) error {
	r.updateLocationLastUpdatedMockData = append(r.updateLocationLastUpdatedMockData, arg)
	return nil
}

func (r *MockRepositoryService) UpdateLocationPublished(ctx context.Context, arg db.UpdateLocationPublishedParams) error {
	r.updateLocationPublishedMockData = append(r.updateLocationPublishedMockData, arg)
	return nil
}

func (r *MockRepositoryService) UpdateLocationsPublishedByCredential(ctx context.Context, arg db.UpdateLocationsPublishedByCredentialParams) error {
	r.updateLocationsPublishedByCredentialMockData = append(r.updateLocationsPublishedByCredentialMockData, arg)
	return nil
}

func (r *MockRepositoryService) UpdateLocationsPublishedByPartyAndCountryCode(ctx context.Context, arg db.UpdateLocationsPublishedByPartyAndCountryCodeParams) error {
	r.updateLocationsPublishedByPartyAndCountryCodeMockData = append(r.updateLocationsPublishedByPartyAndCountryCodeMockData, arg)
	return nil
}

func (r *MockRepositoryService) UpdateLocationsRemovedByCredential(ctx context.Context, arg db.UpdateLocationsRemovedByCredentialParams) error {
	r.updateLocationsRemovedByCredentialMockData = append(r.updateLocationsRemovedByCredentialMockData, arg)
	return nil
}

func (r *MockRepositoryService) UpdateLocationsRemovedByPartyAndCountryCode(ctx context.Context, arg db.UpdateLocationsRemovedByPartyAndCountryCodeParams) error {
	r.updateLocationsRemovedByPartyAndCountryCodeMockData = append(r.updateLocationsRemovedByPartyAndCountryCodeMockData, arg)
	return nil
}

func (r *MockRepositoryService) GetCreateLocationMockData() (db.CreateLocationParams, error) {
	if len(r.createLocationMockData) == 0 {
		return db.CreateLocationParams{}, ErrorNotFound()
	}

	response := r.createLocationMockData[0]
	r.createLocationMockData = r.createLocationMockData[1:]
	return response, nil
}

func (r *MockRepositoryService) GetUpdateLocationByUidMockData() (db.UpdateLocationByUidParams, error) {
	if len(r.updateLocationByUidMockData) == 0 {
		return db.UpdateLocationByUidParams{}, ErrorNotFound()
	}

	response := r.updateLocationByUidMockData[0]
	r.updateLocationByUidMockData = r.updateLocationByUidMockData[1:]
	return response, nil
}

func (r *MockRepositoryService) GetUpdateLocationPublishedMockData() (db.UpdateLocationPublishedParams, error) {
	if len(r.updateLocationPublishedMockData) == 0 {
		return db.UpdateLocationPublishedParams{}, ErrorNotFound()
	}

	response := r.updateLocationPublishedMockData[0]
	r.updateLocationPublishedMockData = r.updateLocationPublishedMockData[1:]
	return response, nil
}

func (r *MockRepositoryService) GetUpdateLocationsPublishedByCredentialMockData() (db.UpdateLocationsPublishedByCredentialParams, error) {
	if len(r.updateLocationsPublishedByCredentialMockData) == 0 {
		return db.UpdateLocationsPublishedByCredentialParams{}, ErrorNotFound()
	}

	response := r.updateLocationsPublishedByCredentialMockData[0]
	r.updateLocationsPublishedByCredentialMockData = r.updateLocationsPublishedByCredentialMockData[1:]
	return response, nil
}

func (r *MockRepositoryService) GetUpdateLocationsPublishedByPartyAndCountryCodeMockData() (db.UpdateLocationsPublishedByPartyAndCountryCodeParams, error) {
	if len(r.updateLocationsPublishedByPartyAndCountryCodeMockData) == 0 {
		return db.UpdateLocationsPublishedByPartyAndCountryCodeParams{}, ErrorNotFound()
	}

	response := r.updateLocationsPublishedByPartyAndCountryCodeMockData[0]
	r.updateLocationsPublishedByPartyAndCountryCodeMockData = r.updateLocationsPublishedByPartyAndCountryCodeMockData[1:]
	return response, nil
}

func (r *MockRepositoryService) GetUpdateLocationsRemovedByCredentialMockData() (db.UpdateLocationsRemovedByCredentialParams, error) {
	if len(r.updateLocationsRemovedByCredentialMockData) == 0 {
		return db.UpdateLocationsRemovedByCredentialParams{}, ErrorNotFound()
	}

	response := r.updateLocationsRemovedByCredentialMockData[0]
	r.updateLocationsRemovedByCredentialMockData = r.updateLocationsRemovedByCredentialMockData[1:]
	return response, nil
}

func (r *MockRepositoryService) GetUpdateLocationsRemovedByPartyAndCountryCodeMockData() (db.UpdateLocationsRemovedByPartyAndCountryCodeParams, error) {
	if len(r.updateLocationsRemovedByPartyAndCountryCodeMockData) == 0 {
		return db.UpdateLocationsRemovedByPartyAndCountryCodeParams{}, ErrorNotFound()
	}

	response := r.updateLocationsRemovedByPartyAndCountryCodeMockData[0]
	r.updateLocationsRemovedByPartyAndCountryCodeMockData = r.updateLocationsRemovedByPartyAndCountryCodeMockData[1:]
	return response, nil
}

func (r *MockRepositoryService) SetGetLocationMockData(response LocationMockData) {
	r.getLocationMockData = append(r.getLocationMockData, response)
}

func (r *MockRepositoryService) SetGetLocationByLastUpdatedMockData(response LocationMockData) {
	r.getLocationByLastUpdatedMockData = append(r.getLocationByLastUpdatedMockData, response)
}

func (r *MockRepositoryService) SetGetLocationByUidMockData(response LocationMockData) {
	r.getLocationByUidMockData = append(r.getLocationByUidMockData, response)
}

func (r *MockRepositoryService) SetListLocationsMockData(response LocationsMockData) {
	r.listLocationsMockData = append(r.listLocationsMockData, response)
}

func (r *MockRepositoryService) SetListLocationsByCountryMockData(response LocationsMockData) {
	r.listLocationsByCountryMockData = append(r.listLocationsByCountryMockData, response)
}

func (r *MockRepositoryService) SetListLocationsByGeomMockData(response LocationsMockData) {
	r.listLocationsByGeomMockData = append(r.listLocationsByGeomMockData, response)
}
