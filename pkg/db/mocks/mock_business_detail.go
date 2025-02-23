package mocks

import (
	"context"

	"github.com/gaia-charge/go-datastore/pkg/db"
)

type BusinessDetailMockData struct {
	BusinessDetail db.BusinessDetail
	Error          error
}

func (r *MockRepositoryService) CreateBusinessDetail(ctx context.Context, arg db.CreateBusinessDetailParams) (db.BusinessDetail, error) {
	r.createBusinessDetailMockData = append(r.createBusinessDetailMockData, arg)
	return db.BusinessDetail{}, nil
}

func (r *MockRepositoryService) DeleteBusinessDetail(ctx context.Context, id int64) error {
	r.deleteBusinessDetailMockData = append(r.deleteBusinessDetailMockData, id)
	return nil
}

func (r *MockRepositoryService) DeleteBusinessDetailLogo(ctx context.Context, id int64) error {
	r.deleteBusinessDetailLogoMockData = append(r.deleteBusinessDetailLogoMockData, id)
	return nil
}

func (r *MockRepositoryService) GetBusinessDetail(ctx context.Context, id int64) (db.BusinessDetail, error) {
	if len(r.getBusinessDetailMockData) == 0 {
		return db.BusinessDetail{}, ErrorNotFound()
	}

	response := r.getBusinessDetailMockData[0]
	r.getBusinessDetailMockData = r.getBusinessDetailMockData[1:]
	return response.BusinessDetail, response.Error
}

func (r *MockRepositoryService) UpdateBusinessDetail(ctx context.Context, arg db.UpdateBusinessDetailParams) (db.BusinessDetail, error) {
	r.updateBusinessDetailMockData = append(r.updateBusinessDetailMockData, arg)
	return db.BusinessDetail{}, nil
}

func (r *MockRepositoryService) GetCreateBusinessDetailMockData() (db.CreateBusinessDetailParams, error) {
	if len(r.createBusinessDetailMockData) == 0 {
		return db.CreateBusinessDetailParams{}, ErrorNotFound()
	}

	response := r.createBusinessDetailMockData[0]
	r.createBusinessDetailMockData = r.createBusinessDetailMockData[1:]
	return response, nil
}

func (r *MockRepositoryService) GetDeleteBusinessDetailMockData() (int64, error) {
	if len(r.deleteBusinessDetailMockData) == 0 {
		return 0, ErrorNotFound()
	}

	response := r.deleteBusinessDetailMockData[0]
	r.deleteBusinessDetailMockData = r.deleteBusinessDetailMockData[1:]
	return response, nil
}

func (r *MockRepositoryService) GetDeleteBusinessDetailLogoMockData() (int64, error) {
	if len(r.deleteBusinessDetailLogoMockData) == 0 {
		return 0, ErrorNotFound()
	}

	response := r.deleteBusinessDetailLogoMockData[0]
	r.deleteBusinessDetailLogoMockData = r.deleteBusinessDetailLogoMockData[1:]
	return response, nil
}

func (r *MockRepositoryService) GetUpdateBusinessDetailMockData() (db.UpdateBusinessDetailParams, error) {
	if len(r.updateBusinessDetailMockData) == 0 {
		return db.UpdateBusinessDetailParams{}, ErrorNotFound()
	}

	response := r.updateBusinessDetailMockData[0]
	r.updateBusinessDetailMockData = r.updateBusinessDetailMockData[1:]
	return response, nil
}

func (r *MockRepositoryService) SetGetBusinessDetailMockData(response BusinessDetailMockData) {
	r.getBusinessDetailMockData = append(r.getBusinessDetailMockData, response)
}
