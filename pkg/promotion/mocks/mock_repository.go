package mocks

import (
	mocks "github.com/gaia-charge/go-datastore/pkg/db/mocks"
	"github.com/gaia-charge/go-datastore/pkg/promotion"
)

func NewRepository(repositoryService *mocks.MockRepositoryService) promotion.PromotionRepository {
	return promotion.PromotionRepository(repositoryService)
}
