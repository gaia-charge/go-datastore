package mocks

import (
	"context"

	"github.com/gaia-charge/go-datastore/pkg/db"
)

type ElementRestrictionMockData struct {
	ElementRestriction db.ElementRestriction
	Error              error
}

func (r *MockRepositoryService) CreateElementRestriction(ctx context.Context, arg db.CreateElementRestrictionParams) (db.ElementRestriction, error) {
	r.createElementRestrictionMockData = append(r.createElementRestrictionMockData, arg)
	return db.ElementRestriction{}, nil
}

func (r *MockRepositoryService) DeleteElementRestrictions(ctx context.Context, tariffID int64) error {
	r.deleteElementRestrictionsMockData = append(r.deleteElementRestrictionsMockData, tariffID)
	return nil
}

func (r *MockRepositoryService) GetElementRestriction(ctx context.Context, id int64) (db.ElementRestriction, error) {
	if len(r.getElementRestrictionMockData) == 0 {
		return db.ElementRestriction{}, ErrorNotFound()
	}

	response := r.getElementRestrictionMockData[0]
	r.getElementRestrictionMockData = r.getElementRestrictionMockData[1:]
	return response.ElementRestriction, response.Error
}

func (r *MockRepositoryService) UpdateElementRestriction(ctx context.Context, arg db.UpdateElementRestrictionParams) (db.ElementRestriction, error) {
	r.updateElementRestrictionMockData = append(r.updateElementRestrictionMockData, arg)
	return db.ElementRestriction{}, nil
}

func (r *MockRepositoryService) GetCreateElementRestrictionMockData() (db.CreateElementRestrictionParams, error) {
	if len(r.createElementRestrictionMockData) == 0 {
		return db.CreateElementRestrictionParams{}, ErrorNotFound()
	}

	response := r.createElementRestrictionMockData[0]
	r.createElementRestrictionMockData = r.createElementRestrictionMockData[1:]
	return response, nil
}

func (r *MockRepositoryService) GetDeleteElementRestrictionsMockData() (int64, error) {
	if len(r.deleteElementRestrictionsMockData) == 0 {
		return 0, ErrorNotFound()
	}

	response := r.deleteElementRestrictionsMockData[0]
	r.deleteElementRestrictionsMockData = r.deleteElementRestrictionsMockData[1:]
	return response, nil
}

func (r *MockRepositoryService) GetUpdateElementRestrictionMockData() (db.UpdateElementRestrictionParams, error) {
	if len(r.updateElementRestrictionMockData) == 0 {
		return db.UpdateElementRestrictionParams{}, ErrorNotFound()
	}

	response := r.updateElementRestrictionMockData[0]
	r.updateElementRestrictionMockData = r.updateElementRestrictionMockData[1:]
	return response, nil
}

func (r *MockRepositoryService) SetGetElementRestrictionMockData(response ElementRestrictionMockData) {
	r.getElementRestrictionMockData = append(r.getElementRestrictionMockData, response)
}
