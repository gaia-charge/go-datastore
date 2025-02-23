package param

import (
	"github.com/gaia-charge/go-datastore/pkg/db"
)

func NewCreateSessionUpdateParams(session db.Session) db.CreateSessionUpdateParams {
	return db.CreateSessionUpdateParams{
		SessionID:   session.ID,
		UserID:      session.UserID,
		Kwh:         session.Kwh,
		TotalCost:   session.TotalCost,
		Status:      session.Status,
		LastUpdated: session.LastUpdated,
	}
}
