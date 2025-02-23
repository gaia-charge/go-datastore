package geolocation

import (
	"context"

	"github.com/gaia-charge/go-datastore/pkg/db"
)

type GeoLocationRepository interface {
	CreateGeoLocation(ctx context.Context, arg db.CreateGeoLocationParams) (db.GeoLocation, error)
	DeleteGeoLocation(ctx context.Context, id int64) error
	GetGeoLocation(ctx context.Context, id int64) (db.GeoLocation, error)
	UpdateGeoLocation(ctx context.Context, arg db.UpdateGeoLocationParams) (db.GeoLocation, error)
}

func NewRepository(repositoryService *db.RepositoryService) GeoLocationRepository {
	return GeoLocationRepository(repositoryService)
}
