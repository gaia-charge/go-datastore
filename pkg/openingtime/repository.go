package openingtime

import (
	"context"

	"github.com/satimoto/go-datastore/pkg/db"
)

type OpeningTimeRepository interface {
	GetOpeningTime(ctx context.Context, id int64) (db.OpeningTime, error)
	ListExceptionalOpeningPeriods(ctx context.Context, openingTimeID int64) ([]db.ExceptionalPeriod, error)
	ListExceptionalClosingPeriods(ctx context.Context, openingTimeID int64) ([]db.ExceptionalPeriod, error)
	ListRegularHours(ctx context.Context, openingTimeID int64) ([]db.RegularHour, error)
}

func NewRepository(repositoryService *db.RepositoryService) OpeningTimeRepository {
	return OpeningTimeRepository(repositoryService)
}
