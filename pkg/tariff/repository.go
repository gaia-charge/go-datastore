package tariff

import (
	"context"
	"database/sql"

	"github.com/gaia-charge/go-datastore/pkg/db"
)

type TariffRepository interface {
	CreateTariff(ctx context.Context, arg db.CreateTariffParams) (db.Tariff, error)
	DeleteTariffAltTexts(ctx context.Context, tariffID int64) error
	DeleteTariffByUid(ctx context.Context, uid string) error
	GetElementRestriction(ctx context.Context, id int64) (db.ElementRestriction, error)
	GetPriceComponentRounding(ctx context.Context, id int64) (db.PriceComponentRounding, error)
	GetTariff(ctx context.Context, id int64) (db.Tariff, error)
	GetTariffByLastUpdated(ctx context.Context, arg db.GetTariffByLastUpdatedParams) (db.Tariff, error)
	GetTariffByUid(ctx context.Context, uid string) (db.Tariff, error)
	GetTariffLikeUid(ctx context.Context, uid string) (db.Tariff, error)
	GetTariffRestriction(ctx context.Context, id int64) (db.TariffRestriction, error)
	ListElements(ctx context.Context, tariffID int64) ([]db.Element, error)
	ListElementRestrictionWeekdays(ctx context.Context, elementRestrictionID int64) ([]db.Weekday, error)
	ListPriceComponents(ctx context.Context, elementID int64) ([]db.PriceComponent, error)
	ListTariffAltTexts(ctx context.Context, tariffID int64) ([]db.DisplayText, error)
	ListTariffsByCdr(ctx context.Context, cdrID sql.NullInt64) ([]db.Tariff, error)
	ListTariffRestrictionWeekdays(ctx context.Context, tariffRestrictionID int64) ([]db.Weekday, error)
	SetTariffAltText(ctx context.Context, arg db.SetTariffAltTextParams) error
	UpdateTariffByUid(ctx context.Context, arg db.UpdateTariffByUidParams) (db.Tariff, error)
}

func NewRepository(repositoryService *db.RepositoryService) TariffRepository {
	return TariffRepository(repositoryService)
}
