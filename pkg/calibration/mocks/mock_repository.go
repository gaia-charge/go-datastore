package mocks

import (
	"github.com/gaia-charge/go-datastore/pkg/calibration"
	mocks "github.com/gaia-charge/go-datastore/pkg/db/mocks"
)

func NewRepository(repositoryService *mocks.MockRepositoryService) calibration.CalibrationRepository {
	return calibration.CalibrationRepository(repositoryService)
}
