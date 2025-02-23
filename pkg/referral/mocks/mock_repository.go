package mocks

import (
	mocks "github.com/gaia-charge/go-datastore/pkg/db/mocks"
	"github.com/gaia-charge/go-datastore/pkg/referral"
)

func NewRepository(repositoryService *mocks.MockRepositoryService) referral.ReferralRepository {
	return referral.ReferralRepository(repositoryService)
}
