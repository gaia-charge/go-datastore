package mocks

import (
	"context"

	"github.com/gaia-charge/go-datastore/pkg/db"
)

func (r *MockRepositoryService) SetEvseCapability(ctx context.Context, arg db.SetEvseCapabilityParams) error {
	r.setEvseCapabilityMockData = append(r.setEvseCapabilityMockData, arg)
	return nil
}

func (r *MockRepositoryService) ListEvseCapabilities(ctx context.Context, evseID int64) ([]db.Capability, error) {
	if len(r.listEvseCapabilitiesMockData) == 0 {
		return []db.Capability{}, nil
	}

	response := r.listEvseCapabilitiesMockData[0]
	r.listEvseCapabilitiesMockData = r.listEvseCapabilitiesMockData[1:]
	return response.Capabilities, response.Error
}

func (r *MockRepositoryService) UnsetEvseCapabilities(ctx context.Context, evseID int64) error {
	r.unsetEvseCapabilitiesMockData = append(r.unsetEvseCapabilitiesMockData, evseID)
	return nil
}

func (r *MockRepositoryService) GetSetEvseCapabilityMockData() (db.SetEvseCapabilityParams, error) {
	if len(r.setEvseCapabilityMockData) == 0 {
		return db.SetEvseCapabilityParams{}, ErrorNotFound()
	}

	response := r.setEvseCapabilityMockData[0]
	r.setEvseCapabilityMockData = r.setEvseCapabilityMockData[1:]
	return response, nil
}

func (r *MockRepositoryService) SetListEvseCapabilitiesMockData(response CapabilitiesMockData) {
	r.listEvseCapabilitiesMockData = append(r.listEvseCapabilitiesMockData, response)
}

func (r *MockRepositoryService) GetUnsetEvseCapabilitiesMockData() (int64, error) {
	if len(r.unsetEvseCapabilitiesMockData) == 0 {
		return 0, ErrorNotFound()
	}

	response := r.unsetEvseCapabilitiesMockData[0]
	r.unsetEvseCapabilitiesMockData = r.unsetEvseCapabilitiesMockData[1:]
	return response, nil
}
