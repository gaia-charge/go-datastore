package mocks

import (
	"context"

	"github.com/gaia-charge/go-datastore/pkg/db"
)

type VersionEndpointMockData struct {
	VersionEndpoint db.VersionEndpoint
	Error           error
}

type VersionEndpointsMockData struct {
	VersionEndpoints []db.VersionEndpoint
	Error            error
}

func (r *MockRepositoryService) CreateVersionEndpoint(ctx context.Context, arg db.CreateVersionEndpointParams) (db.VersionEndpoint, error) {
	r.createVersionEndpointMockData = append(r.createVersionEndpointMockData, arg)
	return db.VersionEndpoint{
		VersionID:  arg.VersionID,
		Identifier: arg.Identifier,
		Url:        arg.Url,
	}, nil
}

func (r *MockRepositoryService) DeleteVersionEndpoints(ctx context.Context, versionID int64) error {
	r.deleteVersionEndpointsMockData = append(r.deleteVersionEndpointsMockData, versionID)
	return nil
}

func (r *MockRepositoryService) GetVersionEndpoint(ctx context.Context, id int64) (db.VersionEndpoint, error) {
	if len(r.getVersionEndpointMockData) == 0 {
		return db.VersionEndpoint{}, ErrorNotFound()
	}

	response := r.getVersionEndpointMockData[0]
	r.getVersionEndpointMockData = r.getVersionEndpointMockData[1:]
	return response.VersionEndpoint, response.Error
}

func (r *MockRepositoryService) GetVersionEndpointByIdentity(ctx context.Context, arg db.GetVersionEndpointByIdentityParams) (db.VersionEndpoint, error) {
	if len(r.getVersionEndpointByIdentityMockData) == 0 {
		return db.VersionEndpoint{}, ErrorNotFound()
	}

	response := r.getVersionEndpointByIdentityMockData[0]
	r.getVersionEndpointByIdentityMockData = r.getVersionEndpointByIdentityMockData[1:]
	return response.VersionEndpoint, response.Error
}

func (r *MockRepositoryService) ListVersionEndpoints(ctx context.Context, versionlID int64) ([]db.VersionEndpoint, error) {
	if len(r.listVersionEndpointsMockData) == 0 {
		return []db.VersionEndpoint{}, nil
	}

	response := r.listVersionEndpointsMockData[0]
	r.listVersionEndpointsMockData = r.listVersionEndpointsMockData[1:]
	return response.VersionEndpoints, response.Error
}

func (r *MockRepositoryService) GetCreateVersionEndpointMockData() (db.CreateVersionEndpointParams, error) {
	if len(r.createVersionEndpointMockData) == 0 {
		return db.CreateVersionEndpointParams{}, ErrorNotFound()
	}

	response := r.createVersionEndpointMockData[0]
	r.createVersionEndpointMockData = r.createVersionEndpointMockData[1:]
	return response, nil
}

func (r *MockRepositoryService) GetDeleteVersionEndpointsMockData() (int64, error) {
	if len(r.deleteVersionEndpointsMockData) == 0 {
		return 0, ErrorNotFound()
	}

	response := r.deleteVersionEndpointsMockData[0]
	r.deleteVersionEndpointsMockData = r.deleteVersionEndpointsMockData[1:]
	return response, nil
}

func (r *MockRepositoryService) SetGetVersionEndpointMockData(response VersionEndpointMockData) {
	r.getVersionEndpointMockData = append(r.getVersionEndpointMockData, response)
}

func (r *MockRepositoryService) SetGetVersionEndpointByIdentityMockData(response VersionEndpointMockData) {
	r.getVersionEndpointByIdentityMockData = append(r.getVersionEndpointByIdentityMockData, response)
}

func (r *MockRepositoryService) SetListVersionEndpointsMockData(response VersionEndpointsMockData) {
	r.listVersionEndpointsMockData = append(r.listVersionEndpointsMockData, response)
}
