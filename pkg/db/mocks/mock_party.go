package mocks

import (
	"context"

	"github.com/gaia-charge/go-datastore/pkg/db"
)

type PartyMockData struct {
	Party db.Party
	Error error
}

type PartiesMockData struct {
	Parties []db.Party
	Error   error
}

func (r *MockRepositoryService) CreateParty(ctx context.Context, arg db.CreatePartyParams) (db.Party, error) {
	r.createPartyMockData = append(r.createPartyMockData, arg)
	return db.Party{}, nil
}

func (r *MockRepositoryService) GetParty(ctx context.Context, id int64) (db.Party, error) {
	if len(r.getPartyMockData) == 0 {
		return db.Party{}, ErrorNotFound()
	}

	response := r.getPartyMockData[0]
	r.getPartyMockData = r.getPartyMockData[1:]
	return response.Party, response.Error
}

func (r *MockRepositoryService) GetPartyByCredential(ctx context.Context, arg db.GetPartyByCredentialParams) (db.Party, error) {
	if len(r.getPartyByCredentialMockData) == 0 {
		return db.Party{}, ErrorNotFound()
	}

	response := r.getPartyByCredentialMockData[0]
	r.getPartyByCredentialMockData = r.getPartyByCredentialMockData[1:]
	return response.Party, response.Error
}

func (r *MockRepositoryService) ListPartiesByCredentialID(ctx context.Context, credentialID int64) ([]db.Party, error) {
	if len(r.listPartiesByCredentialIDMockData) == 0 {
		return []db.Party{}, nil
	}

	response := r.listPartiesByCredentialIDMockData[0]
	r.listPartiesByCredentialIDMockData = r.listPartiesByCredentialIDMockData[1:]
	return response.Parties, response.Error
}

func (r *MockRepositoryService) UpdateParty(ctx context.Context, arg db.UpdatePartyParams) (db.Party, error) {
	r.updatePartyMockData = append(r.updatePartyMockData, arg)
	return db.Party{}, nil
}

func (r *MockRepositoryService) UpdatePartyByCredential(ctx context.Context, arg db.UpdatePartyByCredentialParams) (db.Party, error) {
	r.updatePartyByCredentialMockData = append(r.updatePartyByCredentialMockData, arg)
	return db.Party{}, nil
}

func (r *MockRepositoryService) GetCreatePartyMockData() (db.CreatePartyParams, error) {
	if len(r.createPartyMockData) == 0 {
		return db.CreatePartyParams{}, ErrorNotFound()
	}

	response := r.createPartyMockData[0]
	r.createPartyMockData = r.createPartyMockData[1:]
	return response, nil
}

func (r *MockRepositoryService) SetGetPartyMockData(response PartyMockData) {
	r.getPartyMockData = append(r.getPartyMockData, response)
}

func (r *MockRepositoryService) SetGetPartyByCredentialMockData(response PartyMockData) {
	r.getPartyMockData = append(r.getPartyByCredentialMockData, response)
}

func (r *MockRepositoryService) SetListPartiesByCredentialIDMockData(response PartiesMockData) {
	r.listPartiesByCredentialIDMockData = append(r.listPartiesByCredentialIDMockData, response)
}

func (r *MockRepositoryService) GetUpdatePartyMockData() (db.UpdatePartyParams, error) {
	if len(r.updatePartyMockData) == 0 {
		return db.UpdatePartyParams{}, ErrorNotFound()
	}

	response := r.updatePartyMockData[0]
	r.updatePartyMockData = r.updatePartyMockData[1:]
	return response, nil
}

func (r *MockRepositoryService) GetUpdatePartyByCredentialMockData() (db.UpdatePartyByCredentialParams, error) {
	if len(r.updatePartyByCredentialMockData) == 0 {
		return db.UpdatePartyByCredentialParams{}, ErrorNotFound()
	}

	response := r.updatePartyByCredentialMockData[0]
	r.updatePartyByCredentialMockData = r.updatePartyByCredentialMockData[1:]
	return response, nil
}
