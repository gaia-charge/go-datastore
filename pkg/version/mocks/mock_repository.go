package mocks

import (
	mocks "github.com/gaia-charge/go-datastore/pkg/db/mocks"
	"github.com/gaia-charge/go-datastore/pkg/version"
)

func NewRepository(repositoryService *mocks.MockRepositoryService) version.VersionRepository {
	return version.VersionRepository(repositoryService)
}
