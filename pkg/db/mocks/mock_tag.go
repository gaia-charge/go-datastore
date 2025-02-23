package mocks

import (
	"context"

	"github.com/gaia-charge/go-datastore/pkg/db"
)

type TagMockData struct {
	Tag   db.Tag
	Error error
}

type TagsMockData struct {
	Tags  []db.Tag
	Error error
}

func (r *MockRepositoryService) CreateTag(ctx context.Context, arg db.CreateTagParams) (db.Tag, error) {
	r.createTagMockData = append(r.createTagMockData, arg)
	return db.Tag{}, nil
}

func (r *MockRepositoryService) GetTagByKeyValue(ctx context.Context, arg db.GetTagByKeyValueParams) (db.Tag, error) {
	if len(r.getTagByKeyValueMockData) == 0 {
		return db.Tag{}, ErrorNotFound()
	}

	response := r.getTagByKeyValueMockData[0]
	r.getTagByKeyValueMockData = r.getTagByKeyValueMockData[1:]
	return response.Tag, response.Error
}

func (r *MockRepositoryService) GetCreateTagMockData() (db.CreateTagParams, error) {
	if len(r.createTagMockData) == 0 {
		return db.CreateTagParams{}, ErrorNotFound()
	}

	response := r.createTagMockData[0]
	r.createTagMockData = r.createTagMockData[1:]
	return response, nil
}

func (r *MockRepositoryService) SetGetTagByKeyValueMockData(response TagMockData) {
	r.getTagByKeyValueMockData = append(r.getTagByKeyValueMockData, response)
}
