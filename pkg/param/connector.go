package param

import "github.com/gaia-charge/go-datastore/pkg/db"

func NewUpdateConnectorByEvseParams(connector db.Connector) db.UpdateConnectorByEvseParams {
	return db.UpdateConnectorByEvseParams{
		EvseID:             connector.EvseID,
		Uid:                connector.Uid,
		Identifier:         connector.Identifier,
		Standard:           connector.Standard,
		Format:             connector.Format,
		PowerType:          connector.PowerType,
		Voltage:            connector.Voltage,
		Amperage:           connector.Amperage,
		Wattage:            connector.Wattage,
		TariffID:           connector.TariffID,
		TermsAndConditions: connector.TermsAndConditions,
		IsPublished:        connector.IsPublished,
		IsRemoved:          connector.IsRemoved,
		LastUpdated:        connector.LastUpdated,
	}
}
