package mocks

import (
	"context"

	"github.com/gaia-charge/go-datastore/pkg/db"
)

func (r *MockRepositoryService) ListPsbtFundingStateChannelRequests(ctx context.Context, psbtFundingStateID int64) ([]db.ChannelRequest, error) {
	if len(r.listPsbtFundingStateChannelRequestsMockData) == 0 {
		return []db.ChannelRequest{}, nil
	}

	response := r.listPsbtFundingStateChannelRequestsMockData[0]
	r.listPsbtFundingStateChannelRequestsMockData = r.listPsbtFundingStateChannelRequestsMockData[1:]
	return response.ChannelRequests, response.Error
}

func (r *MockRepositoryService) SetPsbtFundingStateChannelRequest(ctx context.Context, arg db.SetPsbtFundingStateChannelRequestParams) error {
	r.setPsbtFundingStateChannelRequestMockData = append(r.setPsbtFundingStateChannelRequestMockData, arg)
	return nil
}

func (r *MockRepositoryService) GetSetPsbtFundingStateChannelRequestMockData() (db.SetPsbtFundingStateChannelRequestParams, error) {
	if len(r.setPsbtFundingStateChannelRequestMockData) == 0 {
		return db.SetPsbtFundingStateChannelRequestParams{}, ErrorNotFound()
	}

	response := r.setPsbtFundingStateChannelRequestMockData[0]
	r.setPsbtFundingStateChannelRequestMockData = r.setPsbtFundingStateChannelRequestMockData[1:]
	return response, nil
}

func (r *MockRepositoryService) SetListPsbtFundingStateChannelRequestsMockData(response ChannelRequestsMockData) {
	r.listPsbtFundingStateChannelRequestsMockData = append(r.listPsbtFundingStateChannelRequestsMockData, response)
}
