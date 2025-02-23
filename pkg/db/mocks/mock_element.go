package mocks

import (
	"context"

	"github.com/gaia-charge/go-datastore/pkg/db"
)

type ElementMockData struct {
	Element db.Element
	Error   error
}

type ElementsMockData struct {
	Elements []db.Element
	Error    error
}

func (r *MockRepositoryService) CreateElement(ctx context.Context, arg db.CreateElementParams) (db.Element, error) {
	r.createElementMockData = append(r.createElementMockData, arg)
	return db.Element{}, nil
}

func (r *MockRepositoryService) DeleteElements(ctx context.Context, tariffID int64) error {
	r.deleteElementsMockData = append(r.deleteElementsMockData, tariffID)
	return nil
}

func (r *MockRepositoryService) ListElements(ctx context.Context, tariffID int64) ([]db.Element, error) {
	if len(r.listElementsMockData) == 0 {
		return []db.Element{}, nil
	}

	response := r.listElementsMockData[0]
	r.listElementsMockData = r.listElementsMockData[1:]
	return response.Elements, response.Error
}

func (r *MockRepositoryService) GetCreateElementMockData() (db.CreateElementParams, error) {
	if len(r.createElementMockData) == 0 {
		return db.CreateElementParams{}, ErrorNotFound()
	}

	response := r.createElementMockData[0]
	r.createElementMockData = r.createElementMockData[1:]
	return response, nil
}

func (r *MockRepositoryService) GetDeleteElementsMockData() (int64, error) {
	if len(r.deleteElementsMockData) == 0 {
		return 0, ErrorNotFound()
	}

	response := r.deleteElementsMockData[0]
	r.deleteElementsMockData = r.deleteElementsMockData[1:]
	return response, nil
}

func (r *MockRepositoryService) SetListElementsMockData(response ElementsMockData) {
	r.listElementsMockData = append(r.listElementsMockData, response)
}
