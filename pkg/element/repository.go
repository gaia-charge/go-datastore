package element

import (
	"context"

	"github.com/gaia-charge/go-datastore/pkg/db"
)

type ElementRepository interface {
	CreateElement(ctx context.Context, arg db.CreateElementParams) (db.Element, error)
	DeleteElements(ctx context.Context, tariffID int64) error
	ListElements(ctx context.Context, tariffID int64) ([]db.Element, error)
}

func NewRepository(repositoryService *db.RepositoryService) ElementRepository {
	return ElementRepository(repositoryService)
}
