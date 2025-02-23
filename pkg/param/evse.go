package param

import (
	"github.com/gaia-charge/go-datastore/pkg/db"
)

func NewUpdateEvseByUidParams(evse db.Evse) db.UpdateEvseByUidParams {
	return db.UpdateEvseByUidParams{
		Uid:               evse.Uid,
		EvseID:            evse.EvseID,
		Identifier:        evse.Identifier,
		Status:            evse.Status,
		FloorLevel:        evse.FloorLevel,
		Geom:              evse.Geom,
		GeoLocationID:     evse.GeoLocationID,
		IsRemoteCapable:   evse.IsRemoteCapable,
		IsRfidCapable:     evse.IsRfidCapable,
		PhysicalReference: evse.PhysicalReference,
		LastUpdated:       evse.LastUpdated,
	}
}
