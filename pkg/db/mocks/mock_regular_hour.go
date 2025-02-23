package mocks

import (
	"context"

	"github.com/gaia-charge/go-datastore/pkg/db"
)

type RegularHoursMockData struct {
	RegularHours []db.RegularHour
	Error        error
}

func (r *MockRepositoryService) CreateRegularHour(ctx context.Context, arg db.CreateRegularHourParams) (db.RegularHour, error) {
	r.createRegularHourMockData = append(r.createRegularHourMockData, arg)
	return db.RegularHour{}, nil
}

func (r *MockRepositoryService) DeleteRegularHours(ctx context.Context, id int64) error {
	r.deleteRegularHoursMockData = append(r.deleteRegularHoursMockData, id)
	return nil
}

func (r *MockRepositoryService) ListRegularHours(ctx context.Context, openingTimeID int64) ([]db.RegularHour, error) {
	if len(r.listRegularHoursMockData) == 0 {
		return []db.RegularHour{}, nil
	}

	response := r.listRegularHoursMockData[0]
	r.listRegularHoursMockData = r.listRegularHoursMockData[1:]
	return response.RegularHours, response.Error
}

func (r *MockRepositoryService) GetCreateRegularHourMockData() (db.CreateRegularHourParams, error) {
	if len(r.createRegularHourMockData) == 0 {
		return db.CreateRegularHourParams{}, ErrorNotFound()
	}

	response := r.createRegularHourMockData[0]
	r.createRegularHourMockData = r.createRegularHourMockData[1:]
	return response, nil
}

func (r *MockRepositoryService) GetDeleteRegularHoursMockData() (int64, error) {
	if len(r.deleteRegularHoursMockData) == 0 {
		return 0, ErrorNotFound()
	}

	response := r.deleteRegularHoursMockData[0]
	r.deleteRegularHoursMockData = r.deleteRegularHoursMockData[1:]
	return response, nil
}

func (r *MockRepositoryService) SetListRegularHoursMockData(response RegularHoursMockData) {
	r.listRegularHoursMockData = append(r.listRegularHoursMockData, response)
}
