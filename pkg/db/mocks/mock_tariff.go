package mocks

import (
	"context"
	"database/sql"

	"github.com/gaia-charge/go-datastore/pkg/db"
)

type TariffMockData struct {
	Tariff db.Tariff
	Error  error
}

type TariffsMockData struct {
	Tariffs []db.Tariff
	Error   error
}

func (r *MockRepositoryService) CreateTariff(ctx context.Context, arg db.CreateTariffParams) (db.Tariff, error) {
	r.createTariffMockData = append(r.createTariffMockData, arg)
	return db.Tariff{
		Uid:          arg.Uid,
		Currency:     arg.Currency,
		TariffAltUrl: arg.TariffAltUrl,
		EnergyMixID:  arg.EnergyMixID,
		LastUpdated:  arg.LastUpdated,
		CdrID:        arg.CdrID,
	}, nil
}

func (r *MockRepositoryService) DeleteTariffByUid(ctx context.Context, uid string) error {
	r.deleteTariffByUidMockData = append(r.deleteTariffByUidMockData, uid)
	return nil
}

func (r *MockRepositoryService) GetTariff(ctx context.Context, id int64) (db.Tariff, error) {
	if len(r.getTariffMockData) == 0 {
		return db.Tariff{}, ErrorNotFound()
	}

	response := r.getTariffMockData[0]
	r.getTariffMockData = r.getTariffMockData[1:]
	return response.Tariff, response.Error
}

func (r *MockRepositoryService) GetTariffByLastUpdated(ctx context.Context, arg db.GetTariffByLastUpdatedParams) (db.Tariff, error) {
	if len(r.getTariffByLastUpdatedMockData) == 0 {
		return db.Tariff{}, ErrorNotFound()
	}

	response := r.getTariffByLastUpdatedMockData[0]
	r.getTariffByLastUpdatedMockData = r.getTariffByLastUpdatedMockData[1:]
	return response.Tariff, response.Error
}

func (r *MockRepositoryService) GetTariffByUid(ctx context.Context, uid string) (db.Tariff, error) {
	if len(r.getTariffByUidMockData) == 0 {
		return db.Tariff{}, ErrorNotFound()
	}

	response := r.getTariffByUidMockData[0]
	r.getTariffByUidMockData = r.getTariffByUidMockData[1:]
	return response.Tariff, response.Error
}

func (r *MockRepositoryService) GetTariffLikeUid(ctx context.Context, uid string) (db.Tariff, error) {
	if len(r.getTariffLikeUidMockData) == 0 {
		return db.Tariff{}, ErrorNotFound()
	}

	response := r.getTariffLikeUidMockData[0]
	r.getTariffLikeUidMockData = r.getTariffLikeUidMockData[1:]
	return response.Tariff, response.Error
}

func (r *MockRepositoryService) ListTariffsByCdr(ctx context.Context, cdrID sql.NullInt64) ([]db.Tariff, error) {
	if len(r.listTariffsByCdrMockData) == 0 {
		return []db.Tariff{}, nil
	}

	response := r.listTariffsByCdrMockData[0]
	r.listTariffsByCdrMockData = r.listTariffsByCdrMockData[1:]
	return response.Tariffs, response.Error
}

func (r *MockRepositoryService) UpdateTariffByUid(ctx context.Context, arg db.UpdateTariffByUidParams) (db.Tariff, error) {
	r.updateTariffByUidMockData = append(r.updateTariffByUidMockData, arg)
	return db.Tariff{}, nil
}

func (r *MockRepositoryService) GetCreateTariffMockData() (db.CreateTariffParams, error) {
	if len(r.createTariffMockData) == 0 {
		return db.CreateTariffParams{}, ErrorNotFound()
	}

	response := r.createTariffMockData[0]
	r.createTariffMockData = r.createTariffMockData[1:]
	return response, nil
}

func (r *MockRepositoryService) GetDeleteTariffByUidMockData() (string, error) {
	if len(r.deleteTariffByUidMockData) == 0 {
		return "", ErrorNotFound()
	}

	response := r.deleteTariffByUidMockData[0]
	r.deleteTariffByUidMockData = r.deleteTariffByUidMockData[1:]
	return response, nil
}

func (r *MockRepositoryService) GetUpdateTariffByUidMockData() (db.UpdateTariffByUidParams, error) {
	if len(r.updateTariffByUidMockData) == 0 {
		return db.UpdateTariffByUidParams{}, ErrorNotFound()
	}

	response := r.updateTariffByUidMockData[0]
	r.updateTariffByUidMockData = r.updateTariffByUidMockData[1:]
	return response, nil
}

func (r *MockRepositoryService) SetGetTariffMockData(response TariffMockData) {
	r.getTariffMockData = append(r.getTariffMockData, response)
}

func (r *MockRepositoryService) SetGetTariffByUidMockData(response TariffMockData) {
	r.getTariffByUidMockData = append(r.getTariffByUidMockData, response)
}

func (r *MockRepositoryService) SetGetTariffLikeUidMockData(response TariffMockData) {
	r.getTariffLikeUidMockData = append(r.getTariffLikeUidMockData, response)
}

func (r *MockRepositoryService) SetGetTariffByLastUpdatedMockData(response TariffMockData) {
	r.getTariffByUidMockData = append(r.getTariffByLastUpdatedMockData, response)
}

func (r *MockRepositoryService) SetListTariffsByCdrMockData(response TariffsMockData) {
	r.listTariffsByCdrMockData = append(r.listTariffsByCdrMockData, response)
}
