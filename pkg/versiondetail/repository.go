package versiondetail

import (
	"context"

	"github.com/gaia-charge/go-datastore/pkg/db"
)

type VersionDetailRepository interface {
	CreateVersionEndpoint(ctx context.Context, arg db.CreateVersionEndpointParams) (db.VersionEndpoint, error)
	DeleteVersionEndpoints(ctx context.Context, versionID int64) error
	GetVersionEndpoint(ctx context.Context, id int64) (db.VersionEndpoint, error)
	GetVersionEndpointByIdentity(ctx context.Context, arg db.GetVersionEndpointByIdentityParams) (db.VersionEndpoint, error)
	ListVersionEndpoints(ctx context.Context, versionID int64) ([]db.VersionEndpoint, error)
}

func NewRepository(repositoryService *db.RepositoryService) VersionDetailRepository {
	return VersionDetailRepository(repositoryService)
}
