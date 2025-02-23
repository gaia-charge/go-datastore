package mocks

import (
	"context"

	"github.com/gaia-charge/go-datastore/pkg/db"
)

type CalibrationValueMockData struct {
	CalibrationValue db.CalibrationValue
	Error            error
}

type CalibrationValuesMockData struct {
	CalibrationValues []db.CalibrationValue
	Error             error
}

func (r *MockRepositoryService) CreateCalibrationValue(ctx context.Context, arg db.CreateCalibrationValueParams) (db.CalibrationValue, error) {
	r.createCalibrationValueMockData = append(r.createCalibrationValueMockData, arg)
	return db.CalibrationValue{}, nil
}

func (r *MockRepositoryService) ListCalibrationValues(ctx context.Context, calibationID int64) ([]db.CalibrationValue, error) {
	if len(r.listCalibrationValuesMockData) == 0 {
		return []db.CalibrationValue{}, nil
	}

	response := r.listCalibrationValuesMockData[0]
	r.listCalibrationValuesMockData = r.listCalibrationValuesMockData[1:]
	return response.CalibrationValues, response.Error
}

func (r *MockRepositoryService) GetCreateCalibrationValueMockData() (db.CreateCalibrationValueParams, error) {
	if len(r.createCalibrationValueMockData) == 0 {
		return db.CreateCalibrationValueParams{}, ErrorNotFound()
	}

	response := r.createCalibrationValueMockData[0]
	r.createCalibrationValueMockData = r.createCalibrationValueMockData[1:]
	return response, nil
}

func (r *MockRepositoryService) SetListCalibrationValuesMockData(response CalibrationValuesMockData) {
	r.listCalibrationValuesMockData = append(r.listCalibrationValuesMockData, response)
}
