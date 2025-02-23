package mocks

import (
	"context"

	"github.com/gaia-charge/go-datastore/pkg/db"
)

type VersionMockData struct {
	Version db.Version
	Error   error
}

type VersionsMockData struct {
	Versions []db.Version
	Error    error
}

func (r *MockRepositoryService) CreateVersion(ctx context.Context, arg db.CreateVersionParams) (db.Version, error) {
	r.createVersionMockData = append(r.createVersionMockData, arg)
	return db.Version{
		CredentialID: arg.CredentialID,
		Version:      arg.Version,
		Url:          arg.Url,
	}, nil
}

func (r *MockRepositoryService) DeleteVersions(ctx context.Context, credentialID int64) error {
	r.deleteVersionsMockData = append(r.deleteVersionsMockData, credentialID)
	return nil
}

func (r *MockRepositoryService) GetVersion(ctx context.Context, id int64) (db.Version, error) {
	if len(r.getVersionMockData) == 0 {
		return db.Version{}, ErrorNotFound()
	}

	response := r.getVersionMockData[0]
	r.getVersionMockData = r.getVersionMockData[1:]
	return response.Version, response.Error
}

func (r *MockRepositoryService) ListVersions(ctx context.Context, credentialID int64) ([]db.Version, error) {
	if len(r.listVersionsMockData) == 0 {
		return []db.Version{}, nil
	}

	response := r.listVersionsMockData[0]
	r.listVersionsMockData = r.listVersionsMockData[1:]
	return response.Versions, response.Error
}

func (r *MockRepositoryService) GetCreateVersionMockData() (db.CreateVersionParams, error) {
	if len(r.createVersionMockData) == 0 {
		return db.CreateVersionParams{}, ErrorNotFound()
	}

	response := r.createVersionMockData[0]
	r.createVersionMockData = r.createVersionMockData[1:]
	return response, nil
}

func (r *MockRepositoryService) GetDeleteVersionsMockData() (int64, error) {
	if len(r.deleteVersionsMockData) == 0 {
		return 0, ErrorNotFound()
	}

	response := r.deleteVersionsMockData[0]
	r.deleteVersionsMockData = r.deleteVersionsMockData[1:]
	return response, nil
}

func (r *MockRepositoryService) SetGetVersionMockData(response VersionMockData) {
	r.getVersionMockData = append(r.getVersionMockData, response)
}

func (r *MockRepositoryService) SetListVersionsMockData(response VersionsMockData) {
	r.listVersionsMockData = append(r.listVersionsMockData, response)
}
