package mocks

import (
	"context"

	"github.com/gaia-charge/go-datastore/pkg/db"
)

func (r *MockRepositoryService) DeleteTariffAltTexts(ctx context.Context, tariffID int64) error {
	r.deleteTariffAltTextsMockData = append(r.deleteTariffAltTextsMockData, tariffID)
	return nil
}

func (r *MockRepositoryService) ListTariffAltTexts(ctx context.Context, tariffID int64) ([]db.DisplayText, error) {
	if len(r.listTariffAltTextsMockData) == 0 {
		return []db.DisplayText{}, nil
	}

	response := r.listTariffAltTextsMockData[0]
	r.listTariffAltTextsMockData = r.listTariffAltTextsMockData[1:]
	return response.DisplayTexts, response.Error
}

func (r *MockRepositoryService) SetTariffAltText(ctx context.Context, arg db.SetTariffAltTextParams) error {
	r.setTariffAltTextMockData = append(r.setTariffAltTextMockData, arg)
	return nil
}

func (r *MockRepositoryService) GetDeleteTariffAltTextsMockData() (int64, error) {
	if len(r.deleteTariffAltTextsMockData) == 0 {
		return 0, ErrorNotFound()
	}

	response := r.deleteTariffAltTextsMockData[0]
	r.deleteTariffAltTextsMockData = r.deleteTariffAltTextsMockData[1:]
	return response, nil
}

func (r *MockRepositoryService) GetSetTariffAltTextMockData() (db.SetTariffAltTextParams, error) {
	if len(r.setTariffAltTextMockData) == 0 {
		return db.SetTariffAltTextParams{}, ErrorNotFound()
	}

	response := r.setTariffAltTextMockData[0]
	r.setTariffAltTextMockData = r.setTariffAltTextMockData[1:]
	return response, nil
}

func (r *MockRepositoryService) SetListTariffAltTextsMockData(response DisplayTextsMockData) {
	r.listTariffAltTextsMockData = append(r.listTariffAltTextsMockData, response)
}
