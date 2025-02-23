package calibration

import (
	"context"

	"github.com/gaia-charge/go-datastore/pkg/db"
)

type CalibrationRepository interface {
	CreateCalibration(ctx context.Context, arg db.CreateCalibrationParams) (db.Calibration, error)
	CreateCalibrationValue(ctx context.Context, arg db.CreateCalibrationValueParams) (db.CalibrationValue, error)
	GetCalibration(ctx context.Context, id int64) (db.Calibration, error)
	ListCalibrationValues(ctx context.Context, calibrationID int64) ([]db.CalibrationValue, error)
}

func NewRepository(repositoryService *db.RepositoryService) CalibrationRepository {
	return CalibrationRepository(repositoryService)
}
