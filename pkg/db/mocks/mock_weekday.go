package mocks

import (
	"context"

	"github.com/gaia-charge/go-datastore/pkg/db"
)

type WeekdaysMockData struct {
	Weekdays []db.Weekday
	Error    error
}

func (r *MockRepositoryService) ListWeekdays(ctx context.Context) ([]db.Weekday, error) {
	if len(r.listWeekdaysMockData) == 0 {
		return []db.Weekday{}, nil
	}

	response := r.listWeekdaysMockData[0]
	r.listWeekdaysMockData = r.listWeekdaysMockData[1:]
	return response.Weekdays, response.Error
}

func (r *MockRepositoryService) SetListWeekdaysMockData(response WeekdaysMockData) {
	r.listWeekdaysMockData = append(r.listWeekdaysMockData, response)
}
