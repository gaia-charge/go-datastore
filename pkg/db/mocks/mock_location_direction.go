package mocks

import (
	"context"

	"github.com/gaia-charge/go-datastore/pkg/db"
)

func (r *MockRepositoryService) DeleteLocationDirections(ctx context.Context, id int64) error {
	r.deleteLocationDirectionsMockData = append(r.deleteLocationDirectionsMockData, id)
	return nil
}

func (r *MockRepositoryService) ListLocationDirections(ctx context.Context, locationID int64) ([]db.DisplayText, error) {
	if len(r.listLocationDirectionsMockData) == 0 {
		return []db.DisplayText{}, nil
	}

	response := r.listLocationDirectionsMockData[0]
	r.listLocationDirectionsMockData = r.listLocationDirectionsMockData[1:]
	return response.DisplayTexts, response.Error
}

func (r *MockRepositoryService) SetLocationDirection(ctx context.Context, arg db.SetLocationDirectionParams) error {
	r.setLocationDirectionMockData = append(r.setLocationDirectionMockData, arg)
	return nil
}

func (r *MockRepositoryService) GetDeleteLocationDirectionsMockData() (int64, error) {
	if len(r.deleteLocationDirectionsMockData) == 0 {
		return 0, ErrorNotFound()
	}

	response := r.deleteLocationDirectionsMockData[0]
	r.deleteLocationDirectionsMockData = r.deleteLocationDirectionsMockData[1:]
	return response, nil
}

func (r *MockRepositoryService) GetSetLocationDirectionMockData() (db.SetLocationDirectionParams, error) {
	if len(r.setLocationDirectionMockData) == 0 {
		return db.SetLocationDirectionParams{}, ErrorNotFound()
	}

	response := r.setLocationDirectionMockData[0]
	r.setLocationDirectionMockData = r.setLocationDirectionMockData[1:]
	return response, nil
}

func (r *MockRepositoryService) SetListLocationDirectionsMockData(response DisplayTextsMockData) {
	r.listLocationDirectionsMockData = append(r.listLocationDirectionsMockData, response)
}
