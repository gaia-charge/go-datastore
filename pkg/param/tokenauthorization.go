package param

import (
	"github.com/gaia-charge/go-datastore/pkg/db"
	"github.com/google/uuid"
)

func NewCreateTokenAuthorizationParams(tokenID int64) db.CreateTokenAuthorizationParams {
	return db.CreateTokenAuthorizationParams{
		TokenID:         tokenID,
		Authorized:      true,
		AuthorizationID: uuid.NewString(),
	}
}

func NewUpdateTokenAuthorizationParams(tokenAuthorization db.TokenAuthorization) db.UpdateTokenAuthorizationByAuthorizationIDParams {
	return db.UpdateTokenAuthorizationByAuthorizationIDParams{
		AuthorizationID: tokenAuthorization.AuthorizationID,
		Authorized:      tokenAuthorization.Authorized,
		CountryCode:     tokenAuthorization.CountryCode,
		PartyID:         tokenAuthorization.PartyID,
	}
}
