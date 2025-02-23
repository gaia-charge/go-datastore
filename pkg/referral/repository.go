package referral

import (
	"context"

	"github.com/gaia-charge/go-datastore/pkg/db"
)

type ReferralRepository interface {
	CreateReferral(ctx context.Context, arg db.CreateReferralParams) (db.Referral, error)
	GetReferralByIpAddress(ctx context.Context, ipAddress string) (db.Referral, error)
}

func NewRepository(repositoryService *db.RepositoryService) ReferralRepository {
	return ReferralRepository(repositoryService)
}
