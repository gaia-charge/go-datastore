package credential

import (
	"context"
	"database/sql"

	"github.com/gaia-charge/go-datastore/pkg/db"
)

type CredentialRepository interface {
	CreateCredential(ctx context.Context, arg db.CreateCredentialParams) (db.Credential, error)
	GetCredential(ctx context.Context, id int64) (db.Credential, error)
	GetCredentialByPartyAndCountryCode(ctx context.Context, arg db.GetCredentialByPartyAndCountryCodeParams) (db.Credential, error)
	GetCredentialByServerToken(ctx context.Context, serverToken sql.NullString) (db.Credential, error)
	ListCredentials(ctx context.Context) ([]db.Credential, error)
	UpdateCredential(ctx context.Context, arg db.UpdateCredentialParams) (db.Credential, error)
}

func NewRepository(repositoryService *db.RepositoryService) CredentialRepository {
	return CredentialRepository(repositoryService)
}
