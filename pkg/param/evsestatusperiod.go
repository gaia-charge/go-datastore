package param

import (
	"time"

	"github.com/gaia-charge/go-datastore/pkg/db"
)

func NewCreateEvseStatusPeriodParams(evse db.Evse, lastUpdated time.Time) db.CreateEvseStatusPeriodParams {
	return db.CreateEvseStatusPeriodParams{
		EvseID:    evse.ID,
		Status:    evse.Status,
		StartDate: evse.LastUpdated,
		EndDate:   lastUpdated,
	}
}
