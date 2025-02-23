package param

import (
	"github.com/gaia-charge/go-datastore/pkg/db"
)

func NewUpdatePoiByUidParams(poi db.Poi) db.UpdatePoiByUidParams {
	return db.UpdatePoiByUidParams{
		Uid:            poi.Uid,
		Name:           poi.Name,
		Geom:           poi.Geom,
		Description:    poi.Description,
		Address:        poi.Address,
		City:           poi.City,
		PostalCode:     poi.PostalCode,
		TagKey:         poi.TagKey,
		TagValue:       poi.TagValue,
		PaymentOnChain: poi.PaymentOnChain,
		PaymentLn:      poi.PaymentLn,
		PaymentLnTap:   poi.PaymentLnTap,
		PaymentUri:     poi.PaymentUri,
		OpeningTimes:   poi.OpeningTimes,
		Phone:          poi.Phone,
		Website:        poi.Website,
		LastUpdated:    poi.LastUpdated,
	}
}
