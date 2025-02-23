package token

import (
	"context"

	"github.com/gaia-charge/go-datastore/pkg/db"
)

type TokenRepository interface {
	CountTokens(ctx context.Context, arg db.CountTokensParams) (int64, error)
	CreateToken(ctx context.Context, arg db.CreateTokenParams) (db.Token, error)
	DeleteTokenByUid(ctx context.Context, uid string) error
	GetToken(ctx context.Context, id int64) (db.Token, error)
	GetTokenByAuthID(ctx context.Context, authID string) (db.Token, error)
	GetTokenByUid(ctx context.Context, uid string) (db.Token, error)
	GetTokenByUserID(ctx context.Context, arg db.GetTokenByUserIDParams) (db.Token, error)
	ListCredentials(ctx context.Context) ([]db.Credential, error)
	ListTokens(ctx context.Context, arg db.ListTokensParams) ([]db.Token, error)
	ListRfidTokensByUserID(ctx context.Context, userID int64) ([]db.Token, error)
	ListTokensByUserID(ctx context.Context, userID int64) ([]db.Token, error)
	UpdateTokenByUid(ctx context.Context, arg db.UpdateTokenByUidParams) (db.Token, error)
}

func NewRepository(repositoryService *db.RepositoryService) TokenRepository {
	return TokenRepository(repositoryService)
}
