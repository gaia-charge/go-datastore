package evse

import (
	"context"
	"database/sql"

	"github.com/gaia-charge/go-datastore/pkg/db"
)

type EvseRepository interface {
	CreateEvse(ctx context.Context, arg db.CreateEvseParams) (db.Evse, error)
	CreateEvseStatusPeriod(ctx context.Context, arg db.CreateEvseStatusPeriodParams) (db.EvseStatusPeriod, error)
	CreateStatusSchedule(ctx context.Context, arg db.CreateStatusScheduleParams) (db.StatusSchedule, error)
	DeleteConnectors(ctx context.Context, evseID int64) error
	DeleteEvseDirections(ctx context.Context, evseID int64) error
	DeleteEvseImages(ctx context.Context, evseID int64) error
	DeleteStatusSchedules(ctx context.Context, evseID int64) error
	GetEvse(ctx context.Context, id int64) (db.Evse, error)
	GetEvseByEvseID(ctx context.Context, evseID sql.NullString) (db.Evse, error)
	GetEvseByIdentifier(ctx context.Context, identifier sql.NullString) (db.Evse, error)
	GetEvseByUid(ctx context.Context, uid string) (db.Evse, error)
	GetGeoLocation(ctx context.Context, id int64) (db.GeoLocation, error)
	ListCapabilities(ctx context.Context) ([]db.Capability, error)
	ListConnectors(ctx context.Context, evseID int64) ([]db.Connector, error)
	ListEvses(ctx context.Context, locationID int64) ([]db.Evse, error)
	ListEvsesLikeEvseID(ctx context.Context, evseID sql.NullString) ([]db.Evse, error)
	ListEvseStatusPeriods(ctx context.Context, evseID int64) ([]db.EvseStatusPeriod, error)
	ListActiveEvses(ctx context.Context, locationID int64) ([]db.Evse, error)
	ListEvseCapabilities(ctx context.Context, evseID int64) ([]db.Capability, error)
	ListEvseDirections(ctx context.Context, evseID int64) ([]db.DisplayText, error)
	ListEvseImages(ctx context.Context, evseID int64) ([]db.Image, error)
	ListEvseParkingRestrictions(ctx context.Context, evseID int64) ([]db.ParkingRestriction, error)
	ListParkingRestrictions(ctx context.Context) ([]db.ParkingRestriction, error)
	ListStatusSchedules(ctx context.Context, evseID int64) ([]db.StatusSchedule, error)
	SetEvseCapability(ctx context.Context, arg db.SetEvseCapabilityParams) error
	SetEvseDirection(ctx context.Context, arg db.SetEvseDirectionParams) error
	SetEvseImage(ctx context.Context, arg db.SetEvseImageParams) error
	SetEvseParkingRestriction(ctx context.Context, arg db.SetEvseParkingRestrictionParams) error
	UnsetEvseCapabilities(ctx context.Context, evseID int64) error
	UnsetEvseParkingRestrictions(ctx context.Context, evseID int64) error
	UpdateEvseByUid(ctx context.Context, arg db.UpdateEvseByUidParams) (db.Evse, error)
	UpdateEvseLastUpdated(ctx context.Context, arg db.UpdateEvseLastUpdatedParams) error
	UpdateLocationLastUpdated(ctx context.Context, arg db.UpdateLocationLastUpdatedParams) error
}

func NewRepository(repositoryService *db.RepositoryService) EvseRepository {
	return EvseRepository(repositoryService)
}
