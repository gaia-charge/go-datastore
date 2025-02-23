package mocks

import (
	"context"

	"github.com/gaia-charge/go-datastore/pkg/db"
)

func (r *MockRepositoryService) ListTokenAuthorizationConnectors(ctx context.Context, tokenAuthorizationID int64) ([]db.Connector, error) {
	if len(r.listTokenAuthorizationConnectorsMockData) == 0 {
		return []db.Connector{}, nil
	}

	response := r.listTokenAuthorizationConnectorsMockData[0]
	r.listTokenAuthorizationConnectorsMockData = r.listTokenAuthorizationConnectorsMockData[1:]
	return response.Connectors, response.Error
}

func (r *MockRepositoryService) SetTokenAuthorizationConnector(ctx context.Context, arg db.SetTokenAuthorizationConnectorParams) error {
	r.setTokenAuthorizationConnectorMockData = append(r.setTokenAuthorizationConnectorMockData, arg)
	return nil
}

func (r *MockRepositoryService) GetSetTokenAuthorizationConnectorMockData() (db.SetTokenAuthorizationConnectorParams, error) {
	if len(r.setTokenAuthorizationConnectorMockData) == 0 {
		return db.SetTokenAuthorizationConnectorParams{}, ErrorNotFound()
	}

	response := r.setTokenAuthorizationConnectorMockData[0]
	r.setTokenAuthorizationConnectorMockData = r.setTokenAuthorizationConnectorMockData[1:]
	return response, nil
}

func (r *MockRepositoryService) SetListTokenAuthorizationConnectorsMockData(response ConnectorsMockData) {
	r.listTokenAuthorizationConnectorsMockData = append(r.listTokenAuthorizationConnectorsMockData, response)
}
