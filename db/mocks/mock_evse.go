package mocks

import (
	"context"

	"github.com/satimoto/go-datastore/db"
)

type EvseResponse struct {
	Evse  db.Evse
	Error error
}

type EvsesResponse struct {
	Evses []db.Evse
	Error error
}

func (r *MockRepositoryService) GetEvseByUid(ctx context.Context, uid string) (db.Evse, error) {
	if len(r.getEvseByUidResponse) == 0 {
		return db.Evse{}, nil
	}

	response := r.getEvseByUidResponse[0]
	r.getEvseByUidResponse = r.getEvseByUidResponse[1:]
	return response.Evse, response.Error

}

func (r *MockRepositoryService) ListEvses(ctx context.Context, locationID int64) ([]db.Evse, error) {
	if len(r.listEvsesResponse) == 0 {
		return []db.Evse{}, nil
	}

	response := r.listEvsesResponse[0]
	r.listEvsesResponse = r.listEvsesResponse[1:]
	return response.Evses, response.Error
}

func (r *MockRepositoryService) SetGetEvseByUidResponse(response EvseResponse) {
	r.getEvseByUidResponse = append(r.getEvseByUidResponse, response)
}

func (r *MockRepositoryService) SetListEvsesResponse(response EvsesResponse) {
	r.listEvsesResponse = append(r.listEvsesResponse, response)
}
