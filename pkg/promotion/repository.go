package promotion

import (
	"context"

	"github.com/gaia-charge/go-datastore/pkg/db"
)

type PromotionRepository interface {
	CreatePromotion(ctx context.Context, arg db.CreatePromotionParams) (db.Promotion, error)
	GetPromotion(ctx context.Context, id int64) (db.Promotion, error)
	GetPromotionByCode(ctx context.Context, code string) (db.Promotion, error)
}

func NewRepository(repositoryService *db.RepositoryService) PromotionRepository {
	return PromotionRepository(repositoryService)
}
