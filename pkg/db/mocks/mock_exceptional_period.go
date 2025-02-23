package mocks

import (
	"context"

	"github.com/gaia-charge/go-datastore/pkg/db"
)

type ExceptionalPeriodsMockData struct {
	ExceptionalPeriods []db.ExceptionalPeriod
	Error              error
}

func (r *MockRepositoryService) CreateExceptionalPeriod(ctx context.Context, arg db.CreateExceptionalPeriodParams) (db.ExceptionalPeriod, error) {
	r.createExceptionalPeriodMockData = append(r.createExceptionalPeriodMockData, arg)
	return db.ExceptionalPeriod{}, nil
}

func (r *MockRepositoryService) DeleteExceptionalClosingPeriods(ctx context.Context, id int64) error {
	r.deleteExceptionalClosingPeriodsMockData = append(r.deleteExceptionalClosingPeriodsMockData, id)
	return nil
}

func (r *MockRepositoryService) DeleteExceptionalOpeningPeriods(ctx context.Context, id int64) error {
	r.deleteExceptionalOpeningPeriodsMockData = append(r.deleteExceptionalOpeningPeriodsMockData, id)
	return nil
}

func (r *MockRepositoryService) ListExceptionalOpeningPeriods(ctx context.Context, openingTimeID int64) ([]db.ExceptionalPeriod, error) {
	if len(r.listExceptionalOpeningPeriodsMockData) == 0 {
		return []db.ExceptionalPeriod{}, nil
	}

	response := r.listExceptionalOpeningPeriodsMockData[0]
	r.listExceptionalOpeningPeriodsMockData = r.listExceptionalOpeningPeriodsMockData[1:]
	return response.ExceptionalPeriods, response.Error
}

func (r *MockRepositoryService) ListExceptionalClosingPeriods(ctx context.Context, openingTimeID int64) ([]db.ExceptionalPeriod, error) {
	if len(r.listExceptionalClosingPeriodsMockData) == 0 {
		return []db.ExceptionalPeriod{}, nil
	}

	response := r.listExceptionalClosingPeriodsMockData[0]
	r.listExceptionalClosingPeriodsMockData = r.listExceptionalClosingPeriodsMockData[1:]
	return response.ExceptionalPeriods, response.Error
}

func (r *MockRepositoryService) GetCreateExceptionalPeriodMockData() (db.CreateExceptionalPeriodParams, error) {
	if len(r.createExceptionalPeriodMockData) == 0 {
		return db.CreateExceptionalPeriodParams{}, ErrorNotFound()
	}

	response := r.createExceptionalPeriodMockData[0]
	r.createExceptionalPeriodMockData = r.createExceptionalPeriodMockData[1:]
	return response, nil
}

func (r *MockRepositoryService) GetDeleteExceptionalClosingPeriodsMockData() (int64, error) {
	if len(r.deleteExceptionalClosingPeriodsMockData) == 0 {
		return 0, ErrorNotFound()
	}

	response := r.deleteExceptionalClosingPeriodsMockData[0]
	r.deleteExceptionalClosingPeriodsMockData = r.deleteExceptionalClosingPeriodsMockData[1:]
	return response, nil
}

func (r *MockRepositoryService) GetDeleteExceptionalOpeningPeriodsMockData() (int64, error) {
	if len(r.deleteExceptionalOpeningPeriodsMockData) == 0 {
		return 0, ErrorNotFound()
	}

	response := r.deleteExceptionalOpeningPeriodsMockData[0]
	r.deleteExceptionalOpeningPeriodsMockData = r.deleteExceptionalOpeningPeriodsMockData[1:]
	return response, nil
}

func (r *MockRepositoryService) SetListExceptionalOpeningPeriodsMockData(response ExceptionalPeriodsMockData) {
	r.listExceptionalOpeningPeriodsMockData = append(r.listExceptionalOpeningPeriodsMockData, response)
}

func (r *MockRepositoryService) SetListExceptionalClosingPeriodsMockData(response ExceptionalPeriodsMockData) {
	r.listExceptionalClosingPeriodsMockData = append(r.listExceptionalClosingPeriodsMockData, response)
}
