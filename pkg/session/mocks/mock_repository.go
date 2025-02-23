package mocks

import (
	mocks "github.com/gaia-charge/go-datastore/pkg/db/mocks"
	"github.com/gaia-charge/go-datastore/pkg/session"
)

func NewRepository(repositoryService *mocks.MockRepositoryService) session.SessionRepository {
	return session.SessionRepository(repositoryService)
}
