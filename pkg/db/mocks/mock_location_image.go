package mocks

import (
	"context"

	"github.com/gaia-charge/go-datastore/pkg/db"
)

func (r *MockRepositoryService) DeleteLocationImages(ctx context.Context, id int64) error {
	r.deleteLocationImagesMockData = append(r.deleteLocationImagesMockData, id)
	return nil
}

func (r *MockRepositoryService) ListLocationImages(ctx context.Context, locationID int64) ([]db.Image, error) {
	if len(r.listLocationImagesMockData) == 0 {
		return []db.Image{}, nil
	}

	response := r.listLocationImagesMockData[0]
	r.listLocationImagesMockData = r.listLocationImagesMockData[1:]
	return response.Images, response.Error
}

func (r *MockRepositoryService) SetLocationImage(ctx context.Context, arg db.SetLocationImageParams) error {
	r.setLocationImageMockData = append(r.setLocationImageMockData, arg)
	return nil
}

func (r *MockRepositoryService) GetDeleteLocationImagesMockData() (int64, error) {
	if len(r.deleteLocationImagesMockData) == 0 {
		return 0, ErrorNotFound()
	}

	response := r.deleteLocationImagesMockData[0]
	r.deleteLocationImagesMockData = r.deleteLocationImagesMockData[1:]
	return response, nil
}

func (r *MockRepositoryService) GetSetLocationImageMockData() (db.SetLocationImageParams, error) {
	if len(r.setLocationImageMockData) == 0 {
		return db.SetLocationImageParams{}, ErrorNotFound()
	}

	response := r.setLocationImageMockData[0]
	r.setLocationImageMockData = r.setLocationImageMockData[1:]
	return response, nil
}

func (r *MockRepositoryService) SetListLocationImagesMockData(response ImagesMockData) {
	r.listLocationImagesMockData = append(r.listLocationImagesMockData, response)
}
