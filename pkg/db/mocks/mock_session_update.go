package mocks

import (
	"context"

	"github.com/gaia-charge/go-datastore/pkg/db"
)

type SessionUpdateMockData struct {
	SessionUpdate db.SessionUpdate
	Error         error
}

type SessionUpdatesMockData struct {
	SessionUpdates []db.SessionUpdate
	Error          error
}

func (r *MockRepositoryService) CreateSessionUpdate(ctx context.Context, arg db.CreateSessionUpdateParams) (db.SessionUpdate, error) {
	r.createSessionUpdateMockData = append(r.createSessionUpdateMockData, arg)
	return db.SessionUpdate{}, nil
}

func (r *MockRepositoryService) ListSessionUpdatesBySessionID(ctx context.Context, sessionID int64) ([]db.SessionUpdate, error) {
	if len(r.listSessionUpdatesBySessionIDMockData) == 0 {
		return []db.SessionUpdate{}, nil
	}

	response := r.listSessionUpdatesBySessionIDMockData[0]
	r.listSessionUpdatesBySessionIDMockData = r.listSessionUpdatesBySessionIDMockData[1:]
	return response.SessionUpdates, response.Error
}

func (r *MockRepositoryService) GetCreateSessionUpdateMockData() (db.CreateSessionUpdateParams, error) {
	if len(r.createSessionUpdateMockData) == 0 {
		return db.CreateSessionUpdateParams{}, ErrorNotFound()
	}

	response := r.createSessionUpdateMockData[0]
	r.createSessionUpdateMockData = r.createSessionUpdateMockData[1:]
	return response, nil
}

func (r *MockRepositoryService) SetListSessionUpdatesBySessionIDMockData(response SessionUpdatesMockData) {
	r.listSessionUpdatesBySessionIDMockData = append(r.listSessionUpdatesBySessionIDMockData, response)
}
