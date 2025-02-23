package mocks

import (
	"github.com/gaia-charge/go-datastore/pkg/credential"
	mocks "github.com/gaia-charge/go-datastore/pkg/db/mocks"
)

func NewRepository(repositoryService *mocks.MockRepositoryService) credential.CredentialRepository {
	return credential.CredentialRepository(repositoryService)
}
