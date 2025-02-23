package mocks

import (
	"context"

	"github.com/gaia-charge/go-datastore/pkg/db"
)

type CalibrationMockData struct {
	Calibration db.Calibration
	Error       error
}

type CalibrationsMockData struct {
	Calibrations []db.Calibration
	Error        error
}

func (r *MockRepositoryService) CreateCalibration(ctx context.Context, arg db.CreateCalibrationParams) (db.Calibration, error) {
	r.createCalibrationMockData = append(r.createCalibrationMockData, arg)
	return db.Calibration{}, nil
}

func (r *MockRepositoryService) GetCalibration(ctx context.Context, id int64) (db.Calibration, error) {
	if len(r.getCalibrationMockData) == 0 {
		return db.Calibration{}, ErrorNotFound()
	}

	response := r.getCalibrationMockData[0]
	r.getCalibrationMockData = r.getCalibrationMockData[1:]
	return response.Calibration, response.Error
}

func (r *MockRepositoryService) GetCreateCalibrationMockData() (db.CreateCalibrationParams, error) {
	if len(r.createCalibrationMockData) == 0 {
		return db.CreateCalibrationParams{}, ErrorNotFound()
	}

	response := r.createCalibrationMockData[0]
	r.createCalibrationMockData = r.createCalibrationMockData[1:]
	return response, nil
}

func (r *MockRepositoryService) SetGetCalibrationMockData(response CalibrationMockData) {
	r.getCalibrationMockData = append(r.getCalibrationMockData, response)
}
