package param

import (
	"github.com/gaia-charge/go-datastore/pkg/db"
)

func NewUpdateTokenByUidParams(token db.Token) db.UpdateTokenByUidParams {
	return db.UpdateTokenByUidParams{
		Uid:          token.Uid,
		Type:         token.Type,
		AuthID:       token.AuthID,
		VisualNumber: token.VisualNumber,
		Issuer:       token.Issuer,
		Allowed:      token.Allowed,
		Valid:        token.Valid,
		Whitelist:    token.Whitelist,
		Language:     token.Language,
		LastUpdated:  token.LastUpdated,
	}
}
