package param

import (
	"github.com/gaia-charge/go-datastore/pkg/db"
)

func NewUpdateTariffByUidParams(tariff db.Tariff) db.UpdateTariffByUidParams {
	return db.UpdateTariffByUidParams{
		Uid:                 tariff.Uid,
		CountryCode:         tariff.CountryCode,
		PartyID:             tariff.PartyID,
		Currency:            tariff.Currency,
		TariffAltUrl:        tariff.TariffAltUrl,
		EnergyMixID:         tariff.EnergyMixID,
		TariffRestrictionID: tariff.TariffRestrictionID,
		LastUpdated:         tariff.LastUpdated,
	}
}
