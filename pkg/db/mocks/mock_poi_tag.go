package mocks

import (
	"context"

	"github.com/gaia-charge/go-datastore/pkg/db"
)

func (r *MockRepositoryService) ListPoiTags(ctx context.Context, poiID int64) ([]db.Tag, error) {
	if len(r.listPoiTagsMockData) == 0 {
		return []db.Tag{}, nil
	}

	response := r.listPoiTagsMockData[0]
	r.listPoiTagsMockData = r.listPoiTagsMockData[1:]
	return response.Tags, response.Error
}

func (r *MockRepositoryService) SetPoiTag(ctx context.Context, arg db.SetPoiTagParams) error {
	r.setPoiTagMockData = append(r.setPoiTagMockData, arg)
	return nil
}

func (r *MockRepositoryService) UnsetPoiTags(ctx context.Context, poiID int64) error {
	r.unsetPoiTagsMockData = append(r.unsetPoiTagsMockData, poiID)
	return nil
}

func (r *MockRepositoryService) GetSetPoiTagMockData() (db.SetPoiTagParams, error) {
	if len(r.setPoiTagMockData) == 0 {
		return db.SetPoiTagParams{}, ErrorNotFound()
	}

	response := r.setPoiTagMockData[0]
	r.setPoiTagMockData = r.setPoiTagMockData[1:]
	return response, nil
}

func (r *MockRepositoryService) SetListPoiTagsMockData(response TagsMockData) {
	r.listPoiTagsMockData = append(r.listPoiTagsMockData, response)
}

func (r *MockRepositoryService) GetUnsetPoiTagsMockData() (int64, error) {
	if len(r.unsetPoiTagsMockData) == 0 {
		return 0, ErrorNotFound()
	}

	response := r.unsetPoiTagsMockData[0]
	r.unsetPoiTagsMockData = r.unsetPoiTagsMockData[1:]
	return response, nil
}
