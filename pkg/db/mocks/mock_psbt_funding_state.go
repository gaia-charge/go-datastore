package mocks

import (
	"context"

	"github.com/gaia-charge/go-datastore/pkg/db"
)

type PsbtFundingStateMockData struct {
	PsbtFundingState db.PsbtFundingState
	Error            error
}

type PsbtFundingStatesMockData struct {
	PsbtFundingStates []db.PsbtFundingState
	Error             error
}

func (r *MockRepositoryService) CreatePsbtFundingState(ctx context.Context, arg db.CreatePsbtFundingStateParams) (db.PsbtFundingState, error) {
	r.createPsbtFundingStateMockData = append(r.createPsbtFundingStateMockData, arg)
	return db.PsbtFundingState{}, nil
}

func (r *MockRepositoryService) GetPsbtFundingState(ctx context.Context, id int64) (db.PsbtFundingState, error) {
	if len(r.getPsbtFundingStateMockData) == 0 {
		return db.PsbtFundingState{}, ErrorNotFound()
	}

	response := r.getPsbtFundingStateMockData[0]
	r.getPsbtFundingStateMockData = r.getPsbtFundingStateMockData[1:]
	return response.PsbtFundingState, response.Error
}

func (r *MockRepositoryService) GetUnfundedPsbtFundingState(ctx context.Context, nodeID int64) (db.PsbtFundingState, error) {
	if len(r.getUnfundedPsbtFundingStateMockData) == 0 {
		return db.PsbtFundingState{}, ErrorNotFound()
	}

	response := r.getUnfundedPsbtFundingStateMockData[0]
	r.getUnfundedPsbtFundingStateMockData = r.getUnfundedPsbtFundingStateMockData[1:]
	return response.PsbtFundingState, response.Error
}

func (r *MockRepositoryService) ListUnfundedPsbtFundingState(ctx context.Context, nodeID int64) ([]db.PsbtFundingState, error) {
	if len(r.listUnfundedPsbtFundingStatesMockData) == 0 {
		return []db.PsbtFundingState{}, nil
	}

	response := r.listUnfundedPsbtFundingStatesMockData[0]
	r.listUnfundedPsbtFundingStatesMockData = r.listUnfundedPsbtFundingStatesMockData[1:]
	return response.PsbtFundingStates, response.Error
}

func (r *MockRepositoryService) UpdatePsbtFundingState(ctx context.Context, arg db.UpdatePsbtFundingStateParams) (db.PsbtFundingState, error) {
	r.updatePsbtFundingStateMockData = append(r.updatePsbtFundingStateMockData, arg)
	return db.PsbtFundingState{}, nil
}

func (r *MockRepositoryService) GetCreatePsbtFundingStateMockData() (db.CreatePsbtFundingStateParams, error) {
	if len(r.createPsbtFundingStateMockData) == 0 {
		return db.CreatePsbtFundingStateParams{}, ErrorNotFound()
	}

	response := r.createPsbtFundingStateMockData[0]
	r.createPsbtFundingStateMockData = r.createPsbtFundingStateMockData[1:]
	return response, nil
}

func (r *MockRepositoryService) GetUpdatePsbtFundingStateMockData() (db.UpdatePsbtFundingStateParams, error) {
	if len(r.updatePsbtFundingStateMockData) == 0 {
		return db.UpdatePsbtFundingStateParams{}, ErrorNotFound()
	}

	response := r.updatePsbtFundingStateMockData[0]
	r.updatePsbtFundingStateMockData = r.updatePsbtFundingStateMockData[1:]
	return response, nil
}

func (r *MockRepositoryService) SetGetPsbtFundingStateMockData(response PsbtFundingStateMockData) {
	r.getPsbtFundingStateMockData = append(r.getPsbtFundingStateMockData, response)
}

func (r *MockRepositoryService) SetGetUnfundedPsbtFundingStateMockData(response PsbtFundingStateMockData) {
	r.getUnfundedPsbtFundingStateMockData = append(r.getUnfundedPsbtFundingStateMockData, response)
}

func (r *MockRepositoryService) SetListPsbtFundingStatesMockData(response PsbtFundingStatesMockData) {
	r.listUnfundedPsbtFundingStatesMockData = append(r.listUnfundedPsbtFundingStatesMockData, response)
}
