package mocks

import (
	"context"

	"github.com/gaia-charge/go-datastore/pkg/db"
)

type ChannelRequestHtlcMockData struct {
	ChannelRequestHtlc db.ChannelRequestHtlc
	Error              error
}

type ChannelRequestHtlcsMockData struct {
	ChannelRequestHtlcs []db.ChannelRequestHtlc
	Error               error
}

func (r *MockRepositoryService) CreateChannelRequestHtlc(ctx context.Context, arg db.CreateChannelRequestHtlcParams) (db.ChannelRequestHtlc, error) {
	r.createChannelRequestHtlcMockData = append(r.createChannelRequestHtlcMockData, arg)
	return db.ChannelRequestHtlc{}, nil
}

func (r *MockRepositoryService) GetChannelRequestHtlc(ctx context.Context, channelRequestID int64) (db.ChannelRequestHtlc, error) {
	if len(r.getChannelRequestHtlcMockData) == 0 {
		return db.ChannelRequestHtlc{}, ErrorNotFound()
	}

	response := r.getChannelRequestHtlcMockData[0]
	r.getChannelRequestHtlcMockData = r.getChannelRequestHtlcMockData[1:]
	return response.ChannelRequestHtlc, response.Error
}

func (r *MockRepositoryService) GetChannelRequestHtlcByCircuitKey(ctx context.Context, arg db.GetChannelRequestHtlcByCircuitKeyParams) (db.ChannelRequestHtlc, error) {
	if len(r.getChannelRequestHtlcByCircuitKeyMockData) == 0 {
		return db.ChannelRequestHtlc{}, ErrorNotFound()
	}

	response := r.getChannelRequestHtlcByCircuitKeyMockData[0]
	r.getChannelRequestHtlcByCircuitKeyMockData = r.getChannelRequestHtlcByCircuitKeyMockData[1:]
	return response.ChannelRequestHtlc, response.Error
}

func (r *MockRepositoryService) ListChannelRequestHtlcs(ctx context.Context, channelRequestID int64) ([]db.ChannelRequestHtlc, error) {
	if len(r.listChannelRequestHtlcsMockData) == 0 {
		return []db.ChannelRequestHtlc{}, nil
	}

	response := r.listChannelRequestHtlcsMockData[0]
	r.listChannelRequestHtlcsMockData = r.listChannelRequestHtlcsMockData[1:]
	return response.ChannelRequestHtlcs, response.Error
}

func (r *MockRepositoryService) UpdateChannelRequestHtlcByCircuitKey(ctx context.Context, arg db.UpdateChannelRequestHtlcByCircuitKeyParams) (db.ChannelRequestHtlc, error) {
	r.updateChannelRequestHtlcByCircuitKeyMockData = append(r.updateChannelRequestHtlcByCircuitKeyMockData, arg)
	return db.ChannelRequestHtlc{}, nil
}

func (r *MockRepositoryService) GetCreateChannelRequestHtlcMockData() (db.CreateChannelRequestHtlcParams, error) {
	if len(r.createChannelRequestHtlcMockData) == 0 {
		return db.CreateChannelRequestHtlcParams{}, ErrorNotFound()
	}

	response := r.createChannelRequestHtlcMockData[0]
	r.createChannelRequestHtlcMockData = r.createChannelRequestHtlcMockData[1:]
	return response, nil
}

func (r *MockRepositoryService) SetGetChannelRequestHtlcMockData(response ChannelRequestHtlcMockData) {
	r.getChannelRequestHtlcMockData = append(r.getChannelRequestHtlcMockData, response)
}

func (r *MockRepositoryService) SetGetChannelRequestHtlcByCircuitKeyMockData(response ChannelRequestHtlcMockData) {
	r.getChannelRequestHtlcByCircuitKeyMockData = append(r.getChannelRequestHtlcByCircuitKeyMockData, response)
}

func (r *MockRepositoryService) SetListChannelRequestHtlcsMockData(response ChannelRequestHtlcsMockData) {
	r.listChannelRequestHtlcsMockData = append(r.listChannelRequestHtlcsMockData, response)
}

func (r *MockRepositoryService) GetUpdateChannelRequestHtlcByCircuitKeyMockData() (db.UpdateChannelRequestHtlcByCircuitKeyParams, error) {
	if len(r.updateChannelRequestHtlcByCircuitKeyMockData) == 0 {
		return db.UpdateChannelRequestHtlcByCircuitKeyParams{}, ErrorNotFound()
	}

	response := r.updateChannelRequestHtlcByCircuitKeyMockData[0]
	r.updateChannelRequestHtlcByCircuitKeyMockData = r.updateChannelRequestHtlcByCircuitKeyMockData[1:]
	return response, nil
}
