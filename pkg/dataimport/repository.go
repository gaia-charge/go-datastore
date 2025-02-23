package dataimport

import (
	"context"

	"github.com/gaia-charge/go-datastore/pkg/db"
)

type DataImportRepository interface {
	GetHtbTariffByName(ctx context.Context, name string) (db.HtbTariff, error)
	ListHtbTariffs(ctx context.Context) ([]db.HtbTariff, error)
}

func NewRepository(repositoryService *db.RepositoryService) DataImportRepository {
	return DataImportRepository(repositoryService)
}
