package mocks

import (
	"context"
	"database/sql"

	"github.com/gaia-charge/go-datastore/pkg/db"
)

type ConnectorMockData struct {
	Connector db.Connector
	Error     error
}

type ConnectorsMockData struct {
	Connectors []db.Connector
	Error      error
}

func (r *MockRepositoryService) CreateConnector(ctx context.Context, arg db.CreateConnectorParams) (db.Connector, error) {
	r.createConnectorMockData = append(r.createConnectorMockData, arg)
	return db.Connector{}, nil
}

func (r *MockRepositoryService) DeleteConnectors(ctx context.Context, evseID int64) error {
	r.deleteConnectorsMockData = append(r.deleteConnectorsMockData, evseID)
	return nil
}

func (r *MockRepositoryService) GetConnector(ctx context.Context, id int64) (db.Connector, error) {
	if len(r.getConnectorMockData) == 0 {
		return db.Connector{}, ErrorNotFound()
	}

	response := r.getConnectorMockData[0]
	r.getConnectorMockData = r.getConnectorMockData[1:]
	return response.Connector, response.Error
}

func (r *MockRepositoryService) GetConnectorByIdentifier(ctx context.Context, identifier sql.NullString) (db.Connector, error) {
	if len(r.getConnectorByIdentifierMockData) == 0 {
		return db.Connector{}, ErrorNotFound()
	}

	response := r.getConnectorByIdentifierMockData[0]
	r.getConnectorByIdentifierMockData = r.getConnectorByIdentifierMockData[1:]
	return response.Connector, response.Error
}

func (r *MockRepositoryService) GetConnectorByEvse(ctx context.Context, arg db.GetConnectorByEvseParams) (db.Connector, error) {
	if len(r.getConnectorByEvseMockData) == 0 {
		return db.Connector{}, ErrorNotFound()
	}

	response := r.getConnectorByEvseMockData[0]
	r.getConnectorByEvseMockData = r.getConnectorByEvseMockData[1:]
	return response.Connector, response.Error
}

func (r *MockRepositoryService) ListConnectors(ctx context.Context, id int64) ([]db.Connector, error) {
	if len(r.listConnectorsMockData) == 0 {
		return []db.Connector{}, nil
	}

	response := r.listConnectorsMockData[0]
	r.listConnectorsMockData = r.listConnectorsMockData[1:]
	return response.Connectors, response.Error
}

func (r *MockRepositoryService) ListConnectorsByEvseID(ctx context.Context, ebseID sql.NullString) ([]db.Connector, error) {
	if len(r.listConnectorsByEvseIDMockData) == 0 {
		return []db.Connector{}, nil
	}

	response := r.listConnectorsByEvseIDMockData[0]
	r.listConnectorsByEvseIDMockData = r.listConnectorsByEvseIDMockData[1:]
	return response.Connectors, response.Error
}

func (r *MockRepositoryService) ListConnectorsByPartyAndCountryCode(ctx context.Context, arg db.ListConnectorsByPartyAndCountryCodeParams) ([]db.Connector, error) {
	if len(r.listConnectorsByPartyAndCountryCodeMockData) == 0 {
		return []db.Connector{}, nil
	}

	response := r.listConnectorsByPartyAndCountryCodeMockData[0]
	r.listConnectorsByPartyAndCountryCodeMockData = r.listConnectorsByPartyAndCountryCodeMockData[1:]
	return response.Connectors, response.Error
}

func (r *MockRepositoryService) UpdateConnectorByEvse(ctx context.Context, arg db.UpdateConnectorByEvseParams) (db.Connector, error) {
	r.updateConnectorByEvseMockData = append(r.updateConnectorByEvseMockData, arg)
	return db.Connector{
		EvseID:             arg.EvseID,
		Uid:                arg.Uid,
		Identifier:         arg.Identifier,
		Standard:           arg.Standard,
		Format:             arg.Format,
		PowerType:          arg.PowerType,
		Voltage:            arg.Voltage,
		Amperage:           arg.Amperage,
		Wattage:            arg.Wattage,
		TariffID:           arg.TariffID,
		TermsAndConditions: arg.TermsAndConditions,
		LastUpdated:        arg.LastUpdated,
	}, nil
}

func (r *MockRepositoryService) GetCreateConnectorMockData() (db.CreateConnectorParams, error) {
	if len(r.createConnectorMockData) == 0 {
		return db.CreateConnectorParams{}, ErrorNotFound()
	}

	response := r.createConnectorMockData[0]
	r.createConnectorMockData = r.createConnectorMockData[1:]
	return response, nil
}

func (r *MockRepositoryService) GetDeleteConnectorsMockData() (int64, error) {
	if len(r.deleteConnectorsMockData) == 0 {
		return 0, ErrorNotFound()
	}

	response := r.deleteConnectorsMockData[0]
	r.deleteConnectorsMockData = r.deleteConnectorsMockData[1:]
	return response, nil
}

func (r *MockRepositoryService) GetUpdateConnectorByEvseMockData() (db.UpdateConnectorByEvseParams, error) {
	if len(r.updateConnectorByEvseMockData) == 0 {
		return db.UpdateConnectorByEvseParams{}, ErrorNotFound()
	}

	response := r.updateConnectorByEvseMockData[0]
	r.updateConnectorByEvseMockData = r.updateConnectorByEvseMockData[1:]
	return response, nil
}

func (r *MockRepositoryService) SetGetConnectorMockData(response ConnectorMockData) {
	r.getConnectorMockData = append(r.getConnectorMockData, response)
}

func (r *MockRepositoryService) SetGetConnectorByIdentifierMockData(response ConnectorMockData) {
	r.getConnectorByIdentifierMockData = append(r.getConnectorByIdentifierMockData, response)
}
func (r *MockRepositoryService) SetGetConnectorByEvseMockData(response ConnectorMockData) {
	r.getConnectorByEvseMockData = append(r.getConnectorByEvseMockData, response)
}

func (r *MockRepositoryService) SetListConnectorsMockData(response ConnectorsMockData) {
	r.listConnectorsMockData = append(r.listConnectorsMockData, response)
}

func (r *MockRepositoryService) SetListConnectorsByEvseIDMockData(response ConnectorsMockData) {
	r.listConnectorsByEvseIDMockData = append(r.listConnectorsByEvseIDMockData, response)
}

func (r *MockRepositoryService) SetListConnectorsByPartyAndCountryCodeMockData(response ConnectorsMockData) {
	r.listConnectorsByPartyAndCountryCodeMockData = append(r.listConnectorsByPartyAndCountryCodeMockData, response)
}
