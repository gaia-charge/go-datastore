package mocks

import (
	"context"

	"github.com/gaia-charge/go-datastore/pkg/db"
)

type PromotionMockData struct {
	Promotion db.Promotion
	Error     error
}

type PromotionsMockData struct {
	Promotions []db.Promotion
	Error      error
}

func (r *MockRepositoryService) CreatePromotion(ctx context.Context, arg db.CreatePromotionParams) (db.Promotion, error) {
	r.createPromotionMockData = append(r.createPromotionMockData, arg)
	return db.Promotion{}, nil
}

func (r *MockRepositoryService) GetPromotion(ctx context.Context, id int64) (db.Promotion, error) {
	if len(r.getPromotionMockData) == 0 {
		return db.Promotion{}, ErrorNotFound()
	}

	response := r.getPromotionMockData[0]
	r.getPromotionMockData = r.getPromotionMockData[1:]
	return response.Promotion, response.Error
}

func (r *MockRepositoryService) GetPromotionByCode(ctx context.Context, code string) (db.Promotion, error) {
	if len(r.getPromotionByCodeMockData) == 0 {
		return db.Promotion{}, ErrorNotFound()
	}

	response := r.getPromotionByCodeMockData[0]
	r.getPromotionByCodeMockData = r.getPromotionByCodeMockData[1:]
	return response.Promotion, response.Error
}

func (r *MockRepositoryService) GetCreatePromotionMockData() (db.CreatePromotionParams, error) {
	if len(r.createPromotionMockData) == 0 {
		return db.CreatePromotionParams{}, ErrorNotFound()
	}

	response := r.createPromotionMockData[0]
	r.createPromotionMockData = r.createPromotionMockData[1:]
	return response, nil
}

func (r *MockRepositoryService) SetGetPromotionMockData(response PromotionMockData) {
	r.getPromotionMockData = append(r.getPromotionMockData, response)
}

func (r *MockRepositoryService) SetGetPromotionByCodeMockData(response PromotionMockData) {
	r.getPromotionByCodeMockData = append(r.getPromotionByCodeMockData, response)
}
