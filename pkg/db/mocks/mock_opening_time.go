package mocks

import (
	"context"

	"github.com/gaia-charge/go-datastore/pkg/db"
)

type OpeningTimeMockData struct {
	OpeningTime db.OpeningTime
	Error       error
}

func (r *MockRepositoryService) CreateOpeningTime(ctx context.Context, twentyfourseven bool) (db.OpeningTime, error) {
	r.createOpeningTimeMockData = append(r.createOpeningTimeMockData, twentyfourseven)
	return db.OpeningTime{}, nil
}

func (r *MockRepositoryService) DeleteOpeningTime(ctx context.Context, id int64) error {
	r.deleteOpeningTimeMockData = append(r.deleteOpeningTimeMockData, id)
	return nil
}

func (r *MockRepositoryService) GetOpeningTime(ctx context.Context, id int64) (db.OpeningTime, error) {
	if len(r.getOpeningTimeMockData) == 0 {
		return db.OpeningTime{}, ErrorNotFound()
	}

	response := r.getOpeningTimeMockData[0]
	r.getOpeningTimeMockData = r.getOpeningTimeMockData[1:]
	return response.OpeningTime, response.Error
}

func (r *MockRepositoryService) UpdateOpeningTime(ctx context.Context, arg db.UpdateOpeningTimeParams) (db.OpeningTime, error) {
	r.updateOpeningTimeMockData = append(r.updateOpeningTimeMockData, arg)
	return db.OpeningTime{}, nil
}

func (r *MockRepositoryService) GetCreateOpeningTimeMockData() (bool, error) {
	if len(r.createOpeningTimeMockData) == 0 {
		return false, ErrorNotFound()
	}

	response := r.createOpeningTimeMockData[0]
	r.createOpeningTimeMockData = r.createOpeningTimeMockData[1:]
	return response, nil
}

func (r *MockRepositoryService) GetDeleteOpeningTimeMockData() (int64, error) {
	if len(r.deleteOpeningTimeMockData) == 0 {
		return 0, ErrorNotFound()
	}

	response := r.deleteOpeningTimeMockData[0]
	r.deleteOpeningTimeMockData = r.deleteOpeningTimeMockData[1:]
	return response, nil
}

func (r *MockRepositoryService) GetUpdateOpeningTimeMockData() (db.UpdateOpeningTimeParams, error) {
	if len(r.updateOpeningTimeMockData) == 0 {
		return db.UpdateOpeningTimeParams{}, ErrorNotFound()
	}

	response := r.updateOpeningTimeMockData[0]
	r.updateOpeningTimeMockData = r.updateOpeningTimeMockData[1:]
	return response, nil
}

func (r *MockRepositoryService) SetGetOpeningTimeMockData(response OpeningTimeMockData) {
	r.getOpeningTimeMockData = append(r.getOpeningTimeMockData, response)
}
