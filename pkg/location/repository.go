package location

import (
	"context"

	"github.com/gaia-charge/go-datastore/pkg/db"
)

type LocationRepository interface {
	CreateAdditionalGeoLocation(ctx context.Context, arg db.CreateAdditionalGeoLocationParams) (db.AdditionalGeoLocation, error)
	CreateLocation(ctx context.Context, arg db.CreateLocationParams) (db.Location, error)
	DeleteAdditionalGeoLocations(ctx context.Context, locationID int64) error
	DeleteLocationDirections(ctx context.Context, locationID int64) error
	DeleteLocationImages(ctx context.Context, locationID int64) error
	GetConnector(ctx context.Context, id int64) (db.Connector, error)
	GetEvse(ctx context.Context, id int64) (db.Evse, error)
	GetLocation(ctx context.Context, id int64) (db.Location, error)
	GetLocationByLastUpdated(ctx context.Context, arg db.GetLocationByLastUpdatedParams) (db.Location, error)
	GetLocationByUid(ctx context.Context, uid string) (db.Location, error)
	ListAdditionalGeoLocations(ctx context.Context, locationID int64) ([]db.AdditionalGeoLocation, error)
	ListEvses(ctx context.Context, locationID int64) ([]db.Evse, error)
	ListActiveEvses(ctx context.Context, locationID int64) ([]db.Evse, error)
	ListFacilities(ctx context.Context) ([]db.Facility, error)
	ListLocationsByCountry(ctx context.Context, country string) ([]db.Location, error)
	ListLocationsByGeom(ctx context.Context, arg db.ListLocationsByGeomParams) ([]db.Location, error)
	ListLocationDirections(ctx context.Context, locationID int64) ([]db.DisplayText, error)
	ListLocationFacilities(ctx context.Context, locationID int64) ([]db.Facility, error)
	ListLocationImages(ctx context.Context, locationID int64) ([]db.Image, error)
	ListLocations(ctx context.Context) ([]db.Location, error)
	SetLocationDirection(ctx context.Context, arg db.SetLocationDirectionParams) error
	SetLocationFacility(ctx context.Context, arg db.SetLocationFacilityParams) error
	SetLocationImage(ctx context.Context, arg db.SetLocationImageParams) error
	UnsetLocationFacilities(ctx context.Context, locationID int64) error
	UpdateLocationByUid(ctx context.Context, arg db.UpdateLocationByUidParams) (db.Location, error)
	UpdateLocationLastUpdated(ctx context.Context, arg db.UpdateLocationLastUpdatedParams) error
	UpdateLocationPublished(ctx context.Context, arg db.UpdateLocationPublishedParams) error
	UpdateLocationsPublishedByCredential(ctx context.Context, arg db.UpdateLocationsPublishedByCredentialParams) error
	UpdateLocationsPublishedByPartyAndCountryCode(ctx context.Context, arg db.UpdateLocationsPublishedByPartyAndCountryCodeParams) error
	UpdateLocationsRemovedByCredential(ctx context.Context, arg db.UpdateLocationsRemovedByCredentialParams) error
	UpdateLocationsRemovedByPartyAndCountryCode(ctx context.Context, arg db.UpdateLocationsRemovedByPartyAndCountryCodeParams) error
}

func NewRepository(repositoryService *db.RepositoryService) LocationRepository {
	return LocationRepository(repositoryService)
}
