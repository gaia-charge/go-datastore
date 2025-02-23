package tokenauthorization

import (
	"context"

	"github.com/gaia-charge/go-datastore/pkg/db"
)

type TokenAuthorizationRepository interface {
	CreateTokenAuthorization(ctx context.Context, arg db.CreateTokenAuthorizationParams) (db.TokenAuthorization, error)
	GetTokenAuthorizationByAuthorizationID(ctx context.Context, authorizationID string) (db.TokenAuthorization, error)
	GetLastTokenAuthorizationByTokenID(ctx context.Context, tokenID int64) (db.TokenAuthorization, error)
	ListTokenAuthorizationConnectors(ctx context.Context, tokenAuthorizationID int64) ([]db.Connector, error)
	ListTokenAuthorizationEvses(ctx context.Context, tokenAuthorizationID int64) ([]db.Evse, error)
	SetTokenAuthorizationConnector(ctx context.Context, arg db.SetTokenAuthorizationConnectorParams) error
	SetTokenAuthorizationEvse(ctx context.Context, arg db.SetTokenAuthorizationEvseParams) error
	UpdateTokenAuthorizationByAuthorizationID(ctx context.Context, arg db.UpdateTokenAuthorizationByAuthorizationIDParams) (db.TokenAuthorization, error)
}

func NewRepository(repositoryService *db.RepositoryService) TokenAuthorizationRepository {
	return TokenAuthorizationRepository(repositoryService)
}
