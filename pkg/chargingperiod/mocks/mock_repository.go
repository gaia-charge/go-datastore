package mocks

import (
	"github.com/gaia-charge/go-datastore/pkg/chargingperiod"
	mocks "github.com/gaia-charge/go-datastore/pkg/db/mocks"
)

func NewRepository(repositoryService *mocks.MockRepositoryService) chargingperiod.ChargingPeriodRepository {
	return chargingperiod.ChargingPeriodRepository(repositoryService)
}
