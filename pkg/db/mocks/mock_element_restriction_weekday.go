package mocks

import (
	"context"

	"github.com/gaia-charge/go-datastore/pkg/db"
)

func (r *MockRepositoryService) SetElementRestrictionWeekday(ctx context.Context, arg db.SetElementRestrictionWeekdayParams) error {
	r.setElementRestrictionWeekdayMockData = append(r.setElementRestrictionWeekdayMockData, arg)
	return nil
}

func (r *MockRepositoryService) ListElementRestrictionWeekdays(ctx context.Context, elementRestrictionID int64) ([]db.Weekday, error) {
	if len(r.listElementRestrictionWeekdaysMockData) == 0 {
		return []db.Weekday{}, nil
	}

	response := r.listElementRestrictionWeekdaysMockData[0]
	r.listElementRestrictionWeekdaysMockData = r.listElementRestrictionWeekdaysMockData[1:]
	return response.Weekdays, response.Error
}

func (r *MockRepositoryService) UnsetElementRestrictionWeekdays(ctx context.Context, elementRestrictionID int64) error {
	r.unsetElementRestrictionWeekdaysMockData = append(r.unsetElementRestrictionWeekdaysMockData, elementRestrictionID)
	return nil
}

func (r *MockRepositoryService) GetSetElementRestrictionWeekdayMockData() (db.SetElementRestrictionWeekdayParams, error) {
	if len(r.setElementRestrictionWeekdayMockData) == 0 {
		return db.SetElementRestrictionWeekdayParams{}, ErrorNotFound()
	}

	response := r.setElementRestrictionWeekdayMockData[0]
	r.setElementRestrictionWeekdayMockData = r.setElementRestrictionWeekdayMockData[1:]
	return response, nil
}

func (r *MockRepositoryService) SetListElementRestrictionWeekdaysMockData(response WeekdaysMockData) {
	r.listElementRestrictionWeekdaysMockData = append(r.listElementRestrictionWeekdaysMockData, response)
}

func (r *MockRepositoryService) GetUnsetElementRestrictionWeekdaysMockData() (int64, error) {
	if len(r.unsetElementRestrictionWeekdaysMockData) == 0 {
		return 0, ErrorNotFound()
	}

	response := r.unsetElementRestrictionWeekdaysMockData[0]
	r.unsetElementRestrictionWeekdaysMockData = r.unsetElementRestrictionWeekdaysMockData[1:]
	return response, nil
}
