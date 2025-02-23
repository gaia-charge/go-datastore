package businessdetail

import (
	"context"

	"github.com/gaia-charge/go-datastore/pkg/db"
)

type BusinessDetailRepository interface {
	CreateBusinessDetail(ctx context.Context, arg db.CreateBusinessDetailParams) (db.BusinessDetail, error)
	DeleteBusinessDetail(ctx context.Context, id int64) error
	DeleteBusinessDetailLogo(ctx context.Context, id int64) error
	GetBusinessDetail(ctx context.Context, id int64) (db.BusinessDetail, error)
	UpdateBusinessDetail(ctx context.Context, arg db.UpdateBusinessDetailParams) (db.BusinessDetail, error)
}

func NewRepository(repositoryService *db.RepositoryService) BusinessDetailRepository {
	return BusinessDetailRepository(repositoryService)
}
