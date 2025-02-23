package tariffrestriction

import (
	"context"

	"github.com/gaia-charge/go-datastore/pkg/db"
)

type TariffRestrictionRepository interface {
	CreateTariffRestriction(ctx context.Context, arg db.CreateTariffRestrictionParams) (db.TariffRestriction, error)
	DeleteTariffRestriction(ctx context.Context, tariffID int64) error
	GetTariffRestriction(ctx context.Context, id int64) (db.TariffRestriction, error)
	ListTariffRestrictionWeekdays(ctx context.Context, tariffRestrictionID int64) ([]db.Weekday, error)
	ListWeekdays(ctx context.Context) ([]db.Weekday, error)
	SetTariffRestrictionWeekday(ctx context.Context, arg db.SetTariffRestrictionWeekdayParams) error
	UnsetTariffRestrictionWeekdays(ctx context.Context, tariffRestrictionID int64) error
	UpdateTariffRestriction(ctx context.Context, arg db.UpdateTariffRestrictionParams) (db.TariffRestriction, error)
}

func NewRepository(repositoryService *db.RepositoryService) TariffRestrictionRepository {
	return TariffRestrictionRepository(repositoryService)
}
