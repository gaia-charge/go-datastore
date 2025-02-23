package image

import (
	"context"

	"github.com/gaia-charge/go-datastore/pkg/db"
)

type ImageRepository interface {
	CreateImage(ctx context.Context, arg db.CreateImageParams) (db.Image, error)
	DeleteImage(ctx context.Context, id int64) error
	GetImage(ctx context.Context, id int64) (db.Image, error)
	UpdateImage(ctx context.Context, arg db.UpdateImageParams) (db.Image, error)
}

func NewRepository(repositoryService *db.RepositoryService) ImageRepository {
	return ImageRepository(repositoryService)
}
