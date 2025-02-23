package openingtime

import (
	"context"

	"github.com/gaia-charge/go-datastore/pkg/db"
)

type OpeningTimeRepository interface {
	CreateExceptionalPeriod(ctx context.Context, arg db.CreateExceptionalPeriodParams) (db.ExceptionalPeriod, error)
	CreateOpeningTime(ctx context.Context, twentyfourseven bool) (db.OpeningTime, error)
	CreateRegularHour(ctx context.Context, arg db.CreateRegularHourParams) (db.RegularHour, error)
	DeleteExceptionalClosingPeriods(ctx context.Context, openingTimeID int64) error
	DeleteExceptionalOpeningPeriods(ctx context.Context, openingTimeID int64) error
	DeleteOpeningTime(ctx context.Context, id int64) error
	DeleteRegularHours(ctx context.Context, openingTimeID int64) error
	GetOpeningTime(ctx context.Context, id int64) (db.OpeningTime, error)
	ListExceptionalOpeningPeriods(ctx context.Context, openingTimeID int64) ([]db.ExceptionalPeriod, error)
	ListExceptionalClosingPeriods(ctx context.Context, openingTimeID int64) ([]db.ExceptionalPeriod, error)
	ListRegularHours(ctx context.Context, openingTimeID int64) ([]db.RegularHour, error)
	UpdateOpeningTime(ctx context.Context, arg db.UpdateOpeningTimeParams) (db.OpeningTime, error)
}

func NewRepository(repositoryService *db.RepositoryService) OpeningTimeRepository {
	return OpeningTimeRepository(repositoryService)
}
