package mocks

import (
	"context"

	"github.com/gaia-charge/go-datastore/pkg/db"
)

func (r *MockRepositoryService) ListTokenAuthorizationEvses(ctx context.Context, tokenAuthorizationID int64) ([]db.Evse, error) {
	if len(r.listTokenAuthorizationEvsesMockData) == 0 {
		return []db.Evse{}, nil
	}

	response := r.listTokenAuthorizationEvsesMockData[0]
	r.listTokenAuthorizationEvsesMockData = r.listTokenAuthorizationEvsesMockData[1:]
	return response.Evses, response.Error
}

func (r *MockRepositoryService) SetTokenAuthorizationEvse(ctx context.Context, arg db.SetTokenAuthorizationEvseParams) error {
	r.setTokenAuthorizationEvseMockData = append(r.setTokenAuthorizationEvseMockData, arg)
	return nil
}

func (r *MockRepositoryService) GetSetTokenAuthorizationEvseMockData() (db.SetTokenAuthorizationEvseParams, error) {
	if len(r.setTokenAuthorizationEvseMockData) == 0 {
		return db.SetTokenAuthorizationEvseParams{}, ErrorNotFound()
	}

	response := r.setTokenAuthorizationEvseMockData[0]
	r.setTokenAuthorizationEvseMockData = r.setTokenAuthorizationEvseMockData[1:]
	return response, nil
}

func (r *MockRepositoryService) SetListTokenAuthorizationEvsesMockData(response EvsesMockData) {
	r.listTokenAuthorizationEvsesMockData = append(r.listTokenAuthorizationEvsesMockData, response)
}
