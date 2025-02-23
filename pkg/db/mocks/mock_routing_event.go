package mocks

import (
	"context"

	"github.com/gaia-charge/go-datastore/pkg/db"
)

type RoutingEventMockData struct {
	RoutingEvent db.RoutingEvent
	Error        error
}

func (r *MockRepositoryService) CreateRoutingEvent(ctx context.Context, arg db.CreateRoutingEventParams) (db.RoutingEvent, error) {
	r.createRoutingEventMockData = append(r.createRoutingEventMockData, arg)
	return db.RoutingEvent{}, nil
}

func (r *MockRepositoryService) UpdateRoutingEvent(ctx context.Context, arg db.UpdateRoutingEventParams) (db.RoutingEvent, error) {
	r.updateRoutingEventMockData = append(r.updateRoutingEventMockData, arg)
	return db.RoutingEvent{}, nil
}

func (r *MockRepositoryService) GetCreateRoutingEventMockData() (db.CreateRoutingEventParams, error) {
	if len(r.createRoutingEventMockData) == 0 {
		return db.CreateRoutingEventParams{}, ErrorNotFound()
	}

	response := r.createRoutingEventMockData[0]
	r.createRoutingEventMockData = r.createRoutingEventMockData[1:]
	return response, nil
}

func (r *MockRepositoryService) GetUpdateRoutingEventMockData() (db.UpdateRoutingEventParams, error) {
	if len(r.updateRoutingEventMockData) == 0 {
		return db.UpdateRoutingEventParams{}, ErrorNotFound()
	}

	response := r.updateRoutingEventMockData[0]
	r.updateRoutingEventMockData = r.updateRoutingEventMockData[1:]
	return response, nil
}
