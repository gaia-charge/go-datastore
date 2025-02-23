package connector

import (
	"context"
	"database/sql"

	"github.com/gaia-charge/go-datastore/pkg/db"
)

type ConnectorRepository interface {
	CreateConnector(ctx context.Context, arg db.CreateConnectorParams) (db.Connector, error)
	GetConnector(ctx context.Context, id int64) (db.Connector, error)
	GetConnectorByIdentifier(ctx context.Context, identifier sql.NullString) (db.Connector, error)
	GetConnectorByEvse(ctx context.Context, arg db.GetConnectorByEvseParams) (db.Connector, error)
	GetGeoLocation(ctx context.Context, id int64) (db.GeoLocation, error)
	GetEvse(ctx context.Context, id int64) (db.Evse, error)
	ListConnectors(ctx context.Context, evseID int64) ([]db.Connector, error)
	ListConnectorsByEvseID(ctx context.Context, evseID sql.NullString) ([]db.Connector, error)
	ListConnectorsByPartyAndCountryCode(ctx context.Context, arg db.ListConnectorsByPartyAndCountryCodeParams) ([]db.Connector, error)
	UpdateConnectorByEvse(ctx context.Context, arg db.UpdateConnectorByEvseParams) (db.Connector, error)
	UpdateEvseLastUpdated(ctx context.Context, arg db.UpdateEvseLastUpdatedParams) error
	UpdateLocationLastUpdated(ctx context.Context, arg db.UpdateLocationLastUpdatedParams) error
}

func NewRepository(repositoryService *db.RepositoryService) ConnectorRepository {
	return ConnectorRepository(repositoryService)
}
