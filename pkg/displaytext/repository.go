package displaytext

import (
	"context"

	"github.com/gaia-charge/go-datastore/pkg/db"
)

type DisplayTextRepository interface {
	CreateDisplayText(ctx context.Context, arg db.CreateDisplayTextParams) (db.DisplayText, error)
	DeleteDisplayText(ctx context.Context, id int64) error
	GetDisplayText(ctx context.Context, id int64) (db.DisplayText, error)
}

func NewRepository(repositoryService *db.RepositoryService) DisplayTextRepository {
	return DisplayTextRepository(repositoryService)
}
