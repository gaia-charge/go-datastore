package mocks

import (
	"context"

	"github.com/gaia-charge/go-datastore/pkg/db"
)

type DisplayTextMockData struct {
	DisplayText db.DisplayText
	Error       error
}

type DisplayTextsMockData struct {
	DisplayTexts []db.DisplayText
	Error        error
}

func (r *MockRepositoryService) CreateDisplayText(ctx context.Context, arg db.CreateDisplayTextParams) (db.DisplayText, error) {
	r.createDisplayTextMockData = append(r.createDisplayTextMockData, arg)
	return db.DisplayText{}, nil
}

func (r *MockRepositoryService) DeleteDisplayText(ctx context.Context, id int64) error {
	r.deleteDisplayTextMockData = append(r.deleteDisplayTextMockData, id)
	return nil
}

func (r *MockRepositoryService) GetDisplayText(ctx context.Context, id int64) (db.DisplayText, error) {
	if len(r.getDisplayTextMockData) == 0 {
		return db.DisplayText{}, ErrorNotFound()
	}

	response := r.getDisplayTextMockData[0]
	r.getDisplayTextMockData = r.getDisplayTextMockData[1:]
	return response.DisplayText, response.Error
}

func (r *MockRepositoryService) GetCreateDisplayTextMockData() (db.CreateDisplayTextParams, error) {
	if len(r.createDisplayTextMockData) == 0 {
		return db.CreateDisplayTextParams{}, ErrorNotFound()
	}

	response := r.createDisplayTextMockData[0]
	r.createDisplayTextMockData = r.createDisplayTextMockData[1:]
	return response, nil
}

func (r *MockRepositoryService) GetDeleteDisplayTextMockData() (int64, error) {
	if len(r.deleteDisplayTextMockData) == 0 {
		return 0, ErrorNotFound()
	}

	response := r.deleteDisplayTextMockData[0]
	r.deleteDisplayTextMockData = r.deleteDisplayTextMockData[1:]
	return response, nil
}

func (r *MockRepositoryService) SetGetDisplayTextMockData(response DisplayTextMockData) {
	r.getDisplayTextMockData = append(r.getDisplayTextMockData, response)
}
