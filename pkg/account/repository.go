package account

import (
	"context"

	"github.com/gaia-charge/go-datastore/pkg/db"
)

type AccountRepository interface {
	CreateCountryAccount(ctx context.Context, arg db.CreateCountryAccountParams) (db.CountryAccount, error)
	CreateCurrency(ctx context.Context, arg db.CreateCurrencyParams) (db.Currency, error)
	GetCountryAccountByCountry(ctx context.Context, country string) (db.CountryAccount, error)
	GetCurrencyByCode(ctx context.Context, code string) (db.Currency, error)
	ListCountryAccounts(ctx context.Context) ([]db.CountryAccount, error)
}

func NewRepository(repositoryService *db.RepositoryService) AccountRepository {
	return AccountRepository(repositoryService)
}
