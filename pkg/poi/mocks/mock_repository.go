package mocks

import (
	mocks "github.com/gaia-charge/go-datastore/pkg/db/mocks"
	"github.com/gaia-charge/go-datastore/pkg/poi"
)

func NewRepository(repositoryService *mocks.MockRepositoryService) poi.PoiRepository {
	return poi.PoiRepository(repositoryService)
}
