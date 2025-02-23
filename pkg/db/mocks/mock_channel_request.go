package mocks

import (
	"context"

	"github.com/gaia-charge/go-datastore/pkg/db"
)

type ChannelRequestMockData struct {
	ChannelRequest db.ChannelRequest
	Error          error
}

type ChannelRequestsMockData struct {
	ChannelRequests []db.ChannelRequest
	Error           error
}

func (r *MockRepositoryService) CreateChannelRequest(ctx context.Context, arg db.CreateChannelRequestParams) (db.ChannelRequest, error) {
	r.createChannelRequestMockData = append(r.createChannelRequestMockData, arg)
	return db.ChannelRequest{}, nil
}

func (r *MockRepositoryService) DeleteChannelRequest(ctx context.Context, id int64) error {
	r.deleteChannelRequestMockData = append(r.deleteChannelRequestMockData, id)
	return nil
}

func (r *MockRepositoryService) GetChannelRequest(ctx context.Context, id int64) (db.ChannelRequest, error) {
	if len(r.getChannelRequestMockData) == 0 {
		return db.ChannelRequest{}, ErrorNotFound()
	}

	response := r.getChannelRequestMockData[0]
	r.getChannelRequestMockData = r.getChannelRequestMockData[1:]
	return response.ChannelRequest, response.Error
}

func (r *MockRepositoryService) GetChannelRequestByChannelPoint(ctx context.Context, arg db.GetChannelRequestByChannelPointParams) (db.ChannelRequest, error) {
	if len(r.getChannelRequestByChannelPointMockData) == 0 {
		return db.ChannelRequest{}, ErrorNotFound()
	}

	response := r.getChannelRequestByChannelPointMockData[0]
	r.getChannelRequestByChannelPointMockData = r.getChannelRequestByChannelPointMockData[1:]
	return response.ChannelRequest, response.Error
}

func (r *MockRepositoryService) GetChannelRequestByPaymentHash(ctx context.Context, paymentHash []byte) (db.ChannelRequest, error) {
	if len(r.getChannelRequestByPaymentHashMockData) == 0 {
		return db.ChannelRequest{}, ErrorNotFound()
	}

	response := r.getChannelRequestByPaymentHashMockData[0]
	r.getChannelRequestByPaymentHashMockData = r.getChannelRequestByPaymentHashMockData[1:]
	return response.ChannelRequest, response.Error
}

func (r *MockRepositoryService) GetChannelRequestByPendingChanId(ctx context.Context, pendingChanId []byte) (db.ChannelRequest, error) {
	if len(r.getChannelRequestByPendingChanIdMockData) == 0 {
		return db.ChannelRequest{}, ErrorNotFound()
	}

	response := r.getChannelRequestByPendingChanIdMockData[0]
	r.getChannelRequestByPendingChanIdMockData = r.getChannelRequestByPendingChanIdMockData[1:]
	return response.ChannelRequest, response.Error
}

func (r *MockRepositoryService) UpdateChannelRequest(ctx context.Context, arg db.UpdateChannelRequestParams) (db.ChannelRequest, error) {
	r.updateChannelRequestMockData = append(r.updateChannelRequestMockData, arg)
	return db.ChannelRequest{}, nil
}

func (r *MockRepositoryService) UpdatePendingChannelRequestByPubkey(ctx context.Context, arg db.UpdatePendingChannelRequestByPubkeyParams) (db.ChannelRequest, error) {
	r.updatePendingChannelRequestByPubkeyMockData = append(r.updatePendingChannelRequestByPubkeyMockData, arg)
	return db.ChannelRequest{}, nil
}

func (r *MockRepositoryService) GetCreateChannelRequestMockData() (db.CreateChannelRequestParams, error) {
	if len(r.createChannelRequestMockData) == 0 {
		return db.CreateChannelRequestParams{}, ErrorNotFound()
	}

	response := r.createChannelRequestMockData[0]
	r.createChannelRequestMockData = r.createChannelRequestMockData[1:]
	return response, nil
}

func (r *MockRepositoryService) GetDeleteChannelRequestMockData() (int64, error) {
	if len(r.deleteChannelRequestMockData) == 0 {
		return 0, ErrorNotFound()
	}

	response := r.deleteChannelRequestMockData[0]
	r.deleteChannelRequestMockData = r.deleteChannelRequestMockData[1:]
	return response, nil
}

func (r *MockRepositoryService) SetGetChannelRequestMockData(response ChannelRequestMockData) {
	r.getChannelRequestMockData = append(r.getChannelRequestMockData, response)
}

func (r *MockRepositoryService) SetGetChannelRequestByChannelPointMockData(response ChannelRequestMockData) {
	r.getChannelRequestByChannelPointMockData = append(r.getChannelRequestByChannelPointMockData, response)
}

func (r *MockRepositoryService) SetGetChannelRequestByPaymentHashMockData(response ChannelRequestMockData) {
	r.getChannelRequestByPaymentHashMockData = append(r.getChannelRequestByPaymentHashMockData, response)
}

func (r *MockRepositoryService) SetGetChannelRequestByPendingChanIdMockData(response ChannelRequestMockData) {
	r.getChannelRequestByPendingChanIdMockData = append(r.getChannelRequestByPendingChanIdMockData, response)
}

func (r *MockRepositoryService) GetUpdateChannelRequestMockData() (db.UpdateChannelRequestParams, error) {
	if len(r.updateChannelRequestMockData) == 0 {
		return db.UpdateChannelRequestParams{}, ErrorNotFound()
	}

	response := r.updateChannelRequestMockData[0]
	r.updateChannelRequestMockData = r.updateChannelRequestMockData[1:]
	return response, nil
}

func (r *MockRepositoryService) GetUpdatePendingChannelRequestByPubkeyMockData() (db.UpdatePendingChannelRequestByPubkeyParams, error) {
	if len(r.updatePendingChannelRequestByPubkeyMockData) == 0 {
		return db.UpdatePendingChannelRequestByPubkeyParams{}, ErrorNotFound()
	}

	response := r.updatePendingChannelRequestByPubkeyMockData[0]
	r.updatePendingChannelRequestByPubkeyMockData = r.updatePendingChannelRequestByPubkeyMockData[1:]
	return response, nil
}
