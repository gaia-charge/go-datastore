package mocks

import (
	"context"
	"database/sql"

	"github.com/gaia-charge/go-datastore/pkg/db"
)

type CredentialMockData struct {
	Credential db.Credential
	Error      error
}

type CredentialsMockData struct {
	Credentials []db.Credential
	Error       error
}

func (r *MockRepositoryService) CreateCredential(ctx context.Context, arg db.CreateCredentialParams) (db.Credential, error) {
	r.createCredentialMockData = append(r.createCredentialMockData, arg)
	return db.Credential{}, nil
}

func (r *MockRepositoryService) GetCredential(ctx context.Context, id int64) (db.Credential, error) {
	if len(r.getCredentialMockData) == 0 {
		return db.Credential{}, ErrorNotFound()
	}

	response := r.getCredentialMockData[0]
	r.getCredentialMockData = r.getCredentialMockData[1:]
	return response.Credential, response.Error
}

func (r *MockRepositoryService) GetCredentialByPartyAndCountryCode(ctx context.Context, arg db.GetCredentialByPartyAndCountryCodeParams) (db.Credential, error) {
	if len(r.getCredentialByPartyAndCountryCodeMockData) == 0 {
		return db.Credential{}, ErrorNotFound()
	}

	response := r.getCredentialByPartyAndCountryCodeMockData[0]
	r.getCredentialByPartyAndCountryCodeMockData = r.getCredentialByPartyAndCountryCodeMockData[1:]
	return response.Credential, response.Error
}

func (r *MockRepositoryService) GetCredentialByServerToken(ctx context.Context, serverToken sql.NullString) (db.Credential, error) {
	if len(r.getCredentialByServerTokenMockData) == 0 {
		return db.Credential{}, ErrorNotFound()
	}

	response := r.getCredentialByServerTokenMockData[0]
	r.getCredentialByServerTokenMockData = r.getCredentialByServerTokenMockData[1:]
	return response.Credential, response.Error
}

func (r *MockRepositoryService) ListCredentials(ctx context.Context) ([]db.Credential, error) {
	if len(r.listCredentialsMockData) == 0 {
		return []db.Credential{}, nil
	}

	response := r.listCredentialsMockData[0]
	r.listCredentialsMockData = r.listCredentialsMockData[1:]
	return response.Credentials, response.Error
}

func (r *MockRepositoryService) UpdateCredential(ctx context.Context, arg db.UpdateCredentialParams) (db.Credential, error) {
	r.updateCredentialMockData = append(r.updateCredentialMockData, arg)
	return db.Credential{
		ID: arg.ID,
	}, nil
}

func (r *MockRepositoryService) GetCreateCredentialMockData() (db.CreateCredentialParams, error) {
	if len(r.createCredentialMockData) == 0 {
		return db.CreateCredentialParams{}, ErrorNotFound()
	}

	response := r.createCredentialMockData[0]
	r.createCredentialMockData = r.createCredentialMockData[1:]
	return response, nil
}

func (r *MockRepositoryService) SetGetCredentialMockData(response CredentialMockData) {
	r.getCredentialMockData = append(r.getCredentialMockData, response)
}

func (r *MockRepositoryService) SetGetCredentialByPartyAndCountryCodeMockData(response CredentialMockData) {
	r.getCredentialByPartyAndCountryCodeMockData = append(r.getCredentialByPartyAndCountryCodeMockData, response)
}

func (r *MockRepositoryService) SetGetCredentialByServerTokenMockData(response CredentialMockData) {
	r.getCredentialByServerTokenMockData = append(r.getCredentialByServerTokenMockData, response)
}

func (r *MockRepositoryService) SetListCredentialsMockData(response CredentialsMockData) {
	r.listCredentialsMockData = append(r.listCredentialsMockData, response)
}

func (r *MockRepositoryService) GetUpdateCredentialMockData() (db.UpdateCredentialParams, error) {
	if len(r.updateCredentialMockData) == 0 {
		return db.UpdateCredentialParams{}, ErrorNotFound()
	}

	response := r.updateCredentialMockData[0]
	r.updateCredentialMockData = r.updateCredentialMockData[1:]
	return response, nil
}
