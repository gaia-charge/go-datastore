package version

import (
	"context"

	"github.com/gaia-charge/go-datastore/pkg/db"
)

type VersionRepository interface {
	CreateVersion(ctx context.Context, arg db.CreateVersionParams) (db.Version, error)
	DeleteVersions(ctx context.Context, credentialID int64) error
	GetVersion(ctx context.Context, id int64) (db.Version, error)
	ListVersions(ctx context.Context, credentialID int64) ([]db.Version, error)
}

func NewRepository(repositoryService *db.RepositoryService) VersionRepository {
	return VersionRepository(repositoryService)
}
