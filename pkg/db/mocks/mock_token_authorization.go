package mocks

import (
	"context"

	"github.com/gaia-charge/go-datastore/pkg/db"
)

type TokenAuthorizationMockData struct {
	TokenAuthorization db.TokenAuthorization
	Error              error
}

type TokenAuthorizationsMockData struct {
	TokenAuthorizations []db.TokenAuthorization
	Error               error
}

func (r *MockRepositoryService) CreateTokenAuthorization(ctx context.Context, arg db.CreateTokenAuthorizationParams) (db.TokenAuthorization, error) {
	r.createTokenAuthorizationMockData = append(r.createTokenAuthorizationMockData, arg)
	return db.TokenAuthorization{}, nil
}

func (r *MockRepositoryService) GetTokenAuthorizationByAuthorizationID(ctx context.Context, authorizationID string) (db.TokenAuthorization, error) {
	if len(r.getTokenAuthorizationByAuthorizationIDMockData) == 0 {
		return db.TokenAuthorization{}, ErrorNotFound()
	}

	response := r.getTokenAuthorizationByAuthorizationIDMockData[0]
	r.getTokenAuthorizationByAuthorizationIDMockData = r.getTokenAuthorizationByAuthorizationIDMockData[1:]
	return response.TokenAuthorization, response.Error
}

func (r *MockRepositoryService) GetLastTokenAuthorizationByTokenID(ctx context.Context, tokenID int64) (db.TokenAuthorization, error) {
	if len(r.getLastTokenAuthorizationByTokenIDMockData) == 0 {
		return db.TokenAuthorization{}, ErrorNotFound()
	}

	response := r.getLastTokenAuthorizationByTokenIDMockData[0]
	r.getLastTokenAuthorizationByTokenIDMockData = r.getLastTokenAuthorizationByTokenIDMockData[1:]
	return response.TokenAuthorization, response.Error
}

func (r *MockRepositoryService) UpdateTokenAuthorizationByAuthorizationID(ctx context.Context, arg db.UpdateTokenAuthorizationByAuthorizationIDParams) (db.TokenAuthorization, error) {
	r.updateTokenAuthorizationByAuthorizationIDMockData = append(r.updateTokenAuthorizationByAuthorizationIDMockData, arg)
	return db.TokenAuthorization{}, nil
}

func (r *MockRepositoryService) GetCreateTokenAuthorizationMockData() (db.CreateTokenAuthorizationParams, error) {
	if len(r.createTokenAuthorizationMockData) == 0 {
		return db.CreateTokenAuthorizationParams{}, ErrorNotFound()
	}

	response := r.createTokenAuthorizationMockData[0]
	r.createTokenAuthorizationMockData = r.createTokenAuthorizationMockData[1:]
	return response, nil
}

func (r *MockRepositoryService) GetUpdateTokenAuthorizationByAuthorizationIDMockData() (db.UpdateTokenAuthorizationByAuthorizationIDParams, error) {
	if len(r.updateTokenAuthorizationByAuthorizationIDMockData) == 0 {
		return db.UpdateTokenAuthorizationByAuthorizationIDParams{}, ErrorNotFound()
	}

	response := r.updateTokenAuthorizationByAuthorizationIDMockData[0]
	r.updateTokenAuthorizationByAuthorizationIDMockData = r.updateTokenAuthorizationByAuthorizationIDMockData[1:]
	return response, nil
}

func (r *MockRepositoryService) SetGetTokenAuthorizationByAuthorizationIDMockData(response TokenAuthorizationMockData) {
	r.getTokenAuthorizationByAuthorizationIDMockData = append(r.getTokenAuthorizationByAuthorizationIDMockData, response)
}

func (r *MockRepositoryService) SetGetLastTokenAuthorizationByTokenIDMockData(response TokenAuthorizationMockData) {
	r.getLastTokenAuthorizationByTokenIDMockData = append(r.getLastTokenAuthorizationByTokenIDMockData, response)
}
