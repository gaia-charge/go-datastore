package mocks

import (
	"context"

	"github.com/gaia-charge/go-datastore/pkg/db"
)

func (r *MockRepositoryService) SetTariffRestrictionWeekday(ctx context.Context, arg db.SetTariffRestrictionWeekdayParams) error {
	r.setTariffRestrictionWeekdayMockData = append(r.setTariffRestrictionWeekdayMockData, arg)
	return nil
}

func (r *MockRepositoryService) ListTariffRestrictionWeekdays(ctx context.Context, tariffRestrictionID int64) ([]db.Weekday, error) {
	if len(r.listTariffRestrictionWeekdaysMockData) == 0 {
		return []db.Weekday{}, nil
	}

	response := r.listTariffRestrictionWeekdaysMockData[0]
	r.listTariffRestrictionWeekdaysMockData = r.listTariffRestrictionWeekdaysMockData[1:]
	return response.Weekdays, response.Error
}

func (r *MockRepositoryService) UnsetTariffRestrictionWeekdays(ctx context.Context, tariffRestrictionID int64) error {
	r.unsetTariffRestrictionWeekdaysMockData = append(r.unsetTariffRestrictionWeekdaysMockData, tariffRestrictionID)
	return nil
}

func (r *MockRepositoryService) GetSetTariffRestrictionWeekdayMockData() (db.SetTariffRestrictionWeekdayParams, error) {
	if len(r.setTariffRestrictionWeekdayMockData) == 0 {
		return db.SetTariffRestrictionWeekdayParams{}, ErrorNotFound()
	}

	response := r.setTariffRestrictionWeekdayMockData[0]
	r.setTariffRestrictionWeekdayMockData = r.setTariffRestrictionWeekdayMockData[1:]
	return response, nil
}

func (r *MockRepositoryService) SetListTariffRestrictionWeekdaysMockData(response WeekdaysMockData) {
	r.listTariffRestrictionWeekdaysMockData = append(r.listTariffRestrictionWeekdaysMockData, response)
}

func (r *MockRepositoryService) GetUnsetTariffRestrictionWeekdaysMockData() (int64, error) {
	if len(r.unsetTariffRestrictionWeekdaysMockData) == 0 {
		return 0, ErrorNotFound()
	}

	response := r.unsetTariffRestrictionWeekdaysMockData[0]
	r.unsetTariffRestrictionWeekdaysMockData = r.unsetTariffRestrictionWeekdaysMockData[1:]
	return response, nil
}
