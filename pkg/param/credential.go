package param

import "github.com/gaia-charge/go-datastore/pkg/db"

func NewUpdateCredentialParams(credential db.Credential) db.UpdateCredentialParams {
	return db.UpdateCredentialParams{
		ID:          credential.ID,
		ClientToken: credential.ClientToken,
		ServerToken: credential.ServerToken,
		Url:         credential.Url,
		CountryCode: credential.CountryCode,
		PartyID:     credential.PartyID,
		IsAvailable: credential.IsAvailable,
		IsHub:       credential.IsHub,
		VersionID:   credential.VersionID,
		LastUpdated: credential.LastUpdated,
	}
}
