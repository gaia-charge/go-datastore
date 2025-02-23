package mocks

import (
	"context"

	"github.com/gaia-charge/go-datastore/pkg/db"
)

type CapabilitiesMockData struct {
	Capabilities []db.Capability
	Error        error
}

func (r *MockRepositoryService) ListCapabilities(ctx context.Context) ([]db.Capability, error) {
	if len(r.listCapabilitiesMockData) == 0 {
		return []db.Capability{}, nil
	}

	response := r.listCapabilitiesMockData[0]
	r.listCapabilitiesMockData = r.listCapabilitiesMockData[1:]
	return response.Capabilities, response.Error
}

func (r *MockRepositoryService) SetListCapabilitiesMockData(response CapabilitiesMockData) {
	r.listCapabilitiesMockData = append(r.listCapabilitiesMockData, response)
}
