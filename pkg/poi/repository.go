package poi

import (
	"context"

	"github.com/gaia-charge/go-datastore/pkg/db"
)

type PoiRepository interface {
	CreatePoi(ctx context.Context, arg db.CreatePoiParams) (db.Poi, error)
	CreateTag(ctx context.Context, arg db.CreateTagParams) (db.Tag, error)
	DeletePoiByUid(ctx context.Context, uid string) error
	GetPoiByLastUpdated(ctx context.Context) (db.Poi, error)
	GetPoiByUid(ctx context.Context, uid string) (db.Poi, error)
	GetTagByKeyValue(ctx context.Context, arg db.GetTagByKeyValueParams) (db.Tag, error)
	ListPoisByGeom(ctx context.Context, arg db.ListPoisByGeomParams) ([]db.Poi, error)
	ListPoiTags(ctx context.Context, poiID int64) ([]db.Tag, error)
	SetPoiTag(ctx context.Context, arg db.SetPoiTagParams) error
	UpdatePoiByUid(ctx context.Context, arg db.UpdatePoiByUidParams) (db.Poi, error)
	UnsetPoiTags(ctx context.Context, poiID int64) error
}

func NewRepository(repositoryService *db.RepositoryService) PoiRepository {
	return PoiRepository(repositoryService)
}
