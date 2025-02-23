package param

import (
	"time"

	"github.com/gaia-charge/go-datastore/pkg/db"
)

func NewUpdateSessionByUidParams(session db.Session) db.UpdateSessionByUidParams {
	return db.UpdateSessionByUidParams{
		Uid:                session.Uid,
		AuthorizationID:    session.AuthorizationID,
		StartDatetime:      session.StartDatetime,
		EndDatetime:        session.EndDatetime,
		Kwh:                session.Kwh,
		AuthMethod:         session.AuthMethod,
		MeterID:            session.MeterID,
		Currency:           session.Currency,
		TotalCost:          session.TotalCost,
		Status:             session.Status,
		InvoiceRequestID:   session.InvoiceRequestID,
		IsConfirmedStarted: session.IsConfirmedStarted,
		IsConfirmedStopped: session.IsConfirmedStopped,
		IsFlagged:          session.IsFlagged,
		LastUpdated:        session.LastUpdated,
	}
}

func NewCreateSessionInvoiceParams(session db.Session) db.CreateSessionInvoiceParams {
	return db.CreateSessionInvoiceParams{
		SessionID:   session.ID,
		Currency:    session.Currency,
		IsSettled:   false,
		IsExpired:   false,
		LastUpdated: time.Now(),
	}
}

func NewUpdateSessionInvoiceParams(sessionInvoice db.SessionInvoice) db.UpdateSessionInvoiceParams {
	return db.UpdateSessionInvoiceParams{
		ID:               sessionInvoice.ID,
		Currency:         sessionInvoice.Currency,
		CurrencyRate:     sessionInvoice.CurrencyRate,
		CurrencyRateMsat: sessionInvoice.CurrencyRateMsat,
		PriceFiat:        sessionInvoice.PriceFiat,
		PriceMsat:        sessionInvoice.PriceMsat,
		CommissionFiat:   sessionInvoice.CommissionFiat,
		CommissionMsat:   sessionInvoice.CommissionMsat,
		TaxFiat:          sessionInvoice.TaxFiat,
		TaxMsat:          sessionInvoice.TaxMsat,
		TotalFiat:        sessionInvoice.TotalFiat,
		TotalMsat:        sessionInvoice.TotalMsat,
		PaymentRequest:   sessionInvoice.PaymentRequest,
		Signature:        sessionInvoice.Signature,
		IsSettled:        sessionInvoice.IsSettled,
		IsExpired:        sessionInvoice.IsExpired,
		EstimatedEnergy:  sessionInvoice.EstimatedEnergy,
		EstimatedTime:    sessionInvoice.EstimatedTime,
		MeteredEnergy:    sessionInvoice.MeteredEnergy,
		MeteredTime:      sessionInvoice.MeteredTime,
		LastUpdated:      time.Now(),
	}
}
