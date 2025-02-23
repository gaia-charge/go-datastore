package emailsubscription

import (
	"context"

	"github.com/gaia-charge/go-datastore/pkg/db"
)

type EmailSubscriptionRepository interface {
	CreateEmailSubscription(ctx context.Context, arg db.CreateEmailSubscriptionParams) (db.EmailSubscription, error)
	GetEmailSubscriptionByEmail(ctx context.Context, email string) (db.EmailSubscription, error)
	UpdateEmailSubscription(ctx context.Context, arg db.UpdateEmailSubscriptionParams) (db.EmailSubscription, error)
}

func NewRepository(repositoryService *db.RepositoryService) EmailSubscriptionRepository {
	return EmailSubscriptionRepository(repositoryService)
}
