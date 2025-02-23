package session

import (
	"context"
	"database/sql"

	"github.com/gaia-charge/go-datastore/pkg/db"
)

type SessionRepository interface {
	CreateSession(ctx context.Context, arg db.CreateSessionParams) (db.Session, error)
	CreateSessionInvoice(ctx context.Context, arg db.CreateSessionInvoiceParams) (db.SessionInvoice, error)
	CreateSessionUpdate(ctx context.Context, arg db.CreateSessionUpdateParams) (db.SessionUpdate, error)
	DeleteSessionChargingPeriods(ctx context.Context, sessionID int64) error
	GetSessionByAuthorizationID(ctx context.Context, authorizationID string) (db.Session, error)
	GetSessionByLastUpdated(ctx context.Context, arg db.GetSessionByLastUpdatedParams) (db.Session, error)
	GetSession(ctx context.Context, id int64) (db.Session, error)
	GetSessionByUid(ctx context.Context, uid string) (db.Session, error)
	GetSessionInvoice(ctx context.Context, id int64) (db.SessionInvoice, error)
	GetSessionInvoiceByPaymentRequest(ctx context.Context, paymentRequest string) (db.SessionInvoice, error)
	GetUnsettledSessionInvoiceBySession(ctx context.Context, sessionID int64) (db.SessionInvoice, error)
	GetUserBySessionID(ctx context.Context, id int64) (db.User, error)
	ListChargingPeriodDimensions(ctx context.Context, chargingPeriodID int64) ([]db.ChargingPeriodDimension, error)
	ListInvoicedSessionsByUserID(ctx context.Context, userID int64) ([]db.Session, error)
	ListInProgressSessionsByNodeID(ctx context.Context, nodeID sql.NullInt64) ([]db.Session, error)
	ListInProgressSessionsByUserID(ctx context.Context, userID int64) ([]db.Session, error)
	ListSessionChargingPeriods(ctx context.Context, sessionID int64) ([]db.ChargingPeriod, error)
	ListSessionInvoices(ctx context.Context, arg db.ListSessionInvoicesParams) ([]db.SessionInvoice, error)
	ListSessionInvoicesByNodeID(ctx context.Context, arg db.ListSessionInvoicesByNodeIDParams) ([]db.SessionInvoice, error)
	ListSessionInvoicesBySessionID(ctx context.Context, sessionID int64) ([]db.SessionInvoice, error)
	ListSessionInvoicesByUserID(ctx context.Context, arg db.ListSessionInvoicesByUserIDParams) ([]db.SessionInvoice, error)
	ListSessionUpdatesBySessionID(ctx context.Context, sessionID int64) ([]db.SessionUpdate, error)
	SetSessionChargingPeriod(ctx context.Context, arg db.SetSessionChargingPeriodParams) error
	UpdateSessionByUid(ctx context.Context, arg db.UpdateSessionByUidParams) (db.Session, error)
	UpdateSessionInvoice(ctx context.Context, arg db.UpdateSessionInvoiceParams) (db.SessionInvoice, error)
	UpdateSessionIsFlaggedByUid(ctx context.Context, arg db.UpdateSessionIsFlaggedByUidParams) error
}

func NewRepository(repositoryService *db.RepositoryService) SessionRepository {
	return SessionRepository(repositoryService)
}
