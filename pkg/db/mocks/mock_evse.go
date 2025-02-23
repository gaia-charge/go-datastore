package mocks

import (
	"context"
	"database/sql"

	"github.com/gaia-charge/go-datastore/pkg/db"
)

type EvseMockData struct {
	Evse  db.Evse
	Error error
}

type EvsesMockData struct {
	Evses []db.Evse
	Error error
}

func (r *MockRepositoryService) CreateEvse(ctx context.Context, arg db.CreateEvseParams) (db.Evse, error) {
	r.createEvseMockData = append(r.createEvseMockData, arg)
	return db.Evse{}, nil
}

func (r *MockRepositoryService) GetEvse(ctx context.Context, id int64) (db.Evse, error) {
	if len(r.getEvseMockData) == 0 {
		return db.Evse{}, ErrorNotFound()
	}

	response := r.getEvseMockData[0]
	r.getEvseMockData = r.getEvseMockData[1:]
	return response.Evse, response.Error
}

func (r *MockRepositoryService) GetEvseByEvseID(ctx context.Context, evseID sql.NullString) (db.Evse, error) {
	if len(r.getEvseByEvseIDMockData) == 0 {
		return db.Evse{}, ErrorNotFound()
	}

	response := r.getEvseByEvseIDMockData[0]
	r.getEvseByEvseIDMockData = r.getEvseByEvseIDMockData[1:]
	return response.Evse, response.Error
}

func (r *MockRepositoryService) GetEvseByIdentifier(ctx context.Context, identifier sql.NullString) (db.Evse, error) {
	if len(r.getEvseByIdentifierMockData) == 0 {
		return db.Evse{}, ErrorNotFound()
	}

	response := r.getEvseByIdentifierMockData[0]
	r.getEvseByIdentifierMockData = r.getEvseByIdentifierMockData[1:]
	return response.Evse, response.Error
}

func (r *MockRepositoryService) GetEvseByUid(ctx context.Context, uid string) (db.Evse, error) {
	if len(r.getEvseByUidMockData) == 0 {
		return db.Evse{}, ErrorNotFound()
	}

	response := r.getEvseByUidMockData[0]
	r.getEvseByUidMockData = r.getEvseByUidMockData[1:]
	return response.Evse, response.Error
}

func (r *MockRepositoryService) ListEvses(ctx context.Context, locationID int64) ([]db.Evse, error) {
	if len(r.listEvsesMockData) == 0 {
		return []db.Evse{}, nil
	}

	response := r.listEvsesMockData[0]
	r.listEvsesMockData = r.listEvsesMockData[1:]
	return response.Evses, response.Error
}

func (r *MockRepositoryService) ListEvsesLikeEvseID(ctx context.Context, evseID sql.NullString) ([]db.Evse, error) {
	if len(r.listEvsesLikeEvseIDMockData) == 0 {
		return []db.Evse{}, nil
	}

	response := r.listEvsesLikeEvseIDMockData[0]
	r.listEvsesLikeEvseIDMockData = r.listEvsesLikeEvseIDMockData[1:]
	return response.Evses, response.Error
}

func (r *MockRepositoryService) ListActiveEvses(ctx context.Context, locationID int64) ([]db.Evse, error) {
	if len(r.listActiveEvsesMockData) == 0 {
		return []db.Evse{}, nil
	}

	response := r.listActiveEvsesMockData[0]
	r.listActiveEvsesMockData = r.listActiveEvsesMockData[1:]
	return response.Evses, response.Error
}

func (r *MockRepositoryService) UpdateEvseByUid(ctx context.Context, arg db.UpdateEvseByUidParams) (db.Evse, error) {
	r.updateEvseByUidMockData = append(r.updateEvseByUidMockData, arg)
	return db.Evse{
		Uid:               arg.Uid,
		EvseID:            arg.EvseID,
		Identifier:        arg.Identifier,
		Status:            arg.Status,
		FloorLevel:        arg.FloorLevel,
		Geom:              arg.Geom,
		GeoLocationID:     arg.GeoLocationID,
		IsRemoteCapable:   arg.IsRemoteCapable,
		IsRfidCapable:     arg.IsRfidCapable,
		PhysicalReference: arg.PhysicalReference,
		LastUpdated:       arg.LastUpdated,
	}, nil
}

func (r *MockRepositoryService) UpdateEvseLastUpdated(ctx context.Context, arg db.UpdateEvseLastUpdatedParams) error {
	r.updateEvseLastUpdatedMockData = append(r.updateEvseLastUpdatedMockData, arg)
	return nil
}

func (r *MockRepositoryService) GetCreateEvseMockData() (db.CreateEvseParams, error) {
	if len(r.createEvseMockData) == 0 {
		return db.CreateEvseParams{}, ErrorNotFound()
	}

	response := r.createEvseMockData[0]
	r.createEvseMockData = r.createEvseMockData[1:]
	return response, nil
}

func (r *MockRepositoryService) GetUpdateEvseByUidMockData() (db.UpdateEvseByUidParams, error) {
	if len(r.updateEvseByUidMockData) == 0 {
		return db.UpdateEvseByUidParams{}, ErrorNotFound()
	}

	response := r.updateEvseByUidMockData[0]
	r.updateEvseByUidMockData = r.updateEvseByUidMockData[1:]
	return response, nil
}

func (r *MockRepositoryService) SetUpdateEvseLastUpdatedMockData() (db.UpdateEvseLastUpdatedParams, error) {
	if len(r.updateEvseLastUpdatedMockData) == 0 {
		return db.UpdateEvseLastUpdatedParams{}, ErrorNotFound()
	}

	response := r.updateEvseLastUpdatedMockData[0]
	r.updateEvseLastUpdatedMockData = r.updateEvseLastUpdatedMockData[1:]
	return response, nil
}

func (r *MockRepositoryService) SetGetEvseMockData(response EvseMockData) {
	r.getEvseMockData = append(r.getEvseMockData, response)
}

func (r *MockRepositoryService) SetGetEvseByEvseIDMockData(response EvseMockData) {
	r.getEvseByEvseIDMockData = append(r.getEvseByEvseIDMockData, response)
}

func (r *MockRepositoryService) SetGetEvseByIdentifierMockData(response EvseMockData) {
	r.getEvseByIdentifierMockData = append(r.getEvseByIdentifierMockData, response)
}

func (r *MockRepositoryService) SetGetEvseByUidMockData(response EvseMockData) {
	r.getEvseByUidMockData = append(r.getEvseByUidMockData, response)
}

func (r *MockRepositoryService) SetListEvsesMockData(response EvsesMockData) {
	r.listEvsesMockData = append(r.listEvsesMockData, response)
}

func (r *MockRepositoryService) SetListEvsesLikeEvseIDMockData(response EvsesMockData) {
	r.listEvsesLikeEvseIDMockData = append(r.listEvsesLikeEvseIDMockData, response)
}

func (r *MockRepositoryService) SetListActiveEvsesMockData(response EvsesMockData) {
	r.listActiveEvsesMockData = append(r.listActiveEvsesMockData, response)
}
