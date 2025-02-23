package mocks

import (
	"context"

	"github.com/gaia-charge/go-datastore/pkg/db"
)

type AuthenticationMockData struct {
	Authentication db.Authentication
	Error          error
}

func (r *MockRepositoryService) CreateAuthentication(ctx context.Context, arg db.CreateAuthenticationParams) (db.Authentication, error) {
	r.createAuthenticationMockData = append(r.createAuthenticationMockData, arg)
	return db.Authentication{}, nil
}

func (r *MockRepositoryService) DeleteAuthentication(ctx context.Context, id int64) error {
	r.deleteAuthenticationMockData = append(r.deleteAuthenticationMockData, id)
	return nil
}

func (r *MockRepositoryService) GetAuthenticationByChallenge(ctx context.Context, challenge string) (db.Authentication, error) {
	if len(r.getAuthenticationByChallengeMockData) == 0 {
		return db.Authentication{}, ErrorNotFound()
	}

	response := r.getAuthenticationByChallengeMockData[0]
	r.getAuthenticationByChallengeMockData = r.getAuthenticationByChallengeMockData[1:]
	return response.Authentication, response.Error
}

func (r *MockRepositoryService) GetAuthenticationByCode(ctx context.Context, code string) (db.Authentication, error) {
	if len(r.getAuthenticationByCodeMockData) == 0 {
		return db.Authentication{}, ErrorNotFound()
	}

	response := r.getAuthenticationByCodeMockData[0]
	r.getAuthenticationByCodeMockData = r.getAuthenticationByCodeMockData[1:]
	return response.Authentication, response.Error
}

func (r *MockRepositoryService) UpdateAuthentication(ctx context.Context, arg db.UpdateAuthenticationParams) (db.Authentication, error) {
	r.updateAuthenticationMockData = append(r.updateAuthenticationMockData, arg)
	return db.Authentication{}, nil
}

func (r *MockRepositoryService) GetCreateAuthenticationMockData() (db.CreateAuthenticationParams, error) {
	if len(r.createAuthenticationMockData) == 0 {
		return db.CreateAuthenticationParams{}, ErrorNotFound()
	}

	response := r.createAuthenticationMockData[0]
	r.createAuthenticationMockData = r.createAuthenticationMockData[1:]
	return response, nil
}

func (r *MockRepositoryService) GetDeleteAuthenticationMockData() (int64, error) {
	if len(r.deleteAuthenticationMockData) == 0 {
		return 0, ErrorNotFound()
	}

	response := r.deleteAuthenticationMockData[0]
	r.deleteAuthenticationMockData = r.deleteAuthenticationMockData[1:]
	return response, nil
}

func (r *MockRepositoryService) SetGetAuthenticationByCodeMockData(response AuthenticationMockData) {
	r.getAuthenticationByCodeMockData = append(r.getAuthenticationByCodeMockData, response)
}

func (r *MockRepositoryService) SetGetAuthenticationByChallengeMockData(response AuthenticationMockData) {
	r.getAuthenticationByChallengeMockData = append(r.getAuthenticationByChallengeMockData, response)
}

func (r *MockRepositoryService) GetUpdateAuthenticationMockData() (db.UpdateAuthenticationParams, error) {
	if len(r.updateAuthenticationMockData) == 0 {
		return db.UpdateAuthenticationParams{}, ErrorNotFound()
	}

	response := r.updateAuthenticationMockData[0]
	r.updateAuthenticationMockData = r.updateAuthenticationMockData[1:]
	return response, nil
}
