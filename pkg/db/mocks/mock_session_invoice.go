package mocks

import (
	"context"

	"github.com/gaia-charge/go-datastore/pkg/db"
)

type SessionInvoiceMockData struct {
	SessionInvoice db.SessionInvoice
	Error          error
}

type SessionInvoicesMockData struct {
	SessionInvoices []db.SessionInvoice
	Error           error
}

func (r *MockRepositoryService) CreateSessionInvoice(ctx context.Context, arg db.CreateSessionInvoiceParams) (db.SessionInvoice, error) {
	r.createSessionInvoiceMockData = append(r.createSessionInvoiceMockData, arg)
	return db.SessionInvoice{}, nil
}

func (r *MockRepositoryService) GetSessionInvoice(ctx context.Context, id int64) (db.SessionInvoice, error) {
	if len(r.getSessionInvoiceMockData) == 0 {
		return db.SessionInvoice{}, ErrorNotFound()
	}

	response := r.getSessionInvoiceMockData[0]
	r.getSessionInvoiceMockData = r.getSessionInvoiceMockData[1:]
	return response.SessionInvoice, response.Error
}

func (r *MockRepositoryService) GetSessionInvoiceByPaymentRequest(ctx context.Context, paymentRequest string) (db.SessionInvoice, error) {
	if len(r.getSessionInvoiceByPaymentRequestMockData) == 0 {
		return db.SessionInvoice{}, ErrorNotFound()
	}

	response := r.getSessionInvoiceByPaymentRequestMockData[0]
	r.getSessionInvoiceByPaymentRequestMockData = r.getSessionInvoiceByPaymentRequestMockData[1:]
	return response.SessionInvoice, response.Error
}

func (r *MockRepositoryService) GetUnsettledSessionInvoiceBySession(ctx context.Context, sessionID int64) (db.SessionInvoice, error) {
	if len(r.getUnsettledSessionInvoiceBySessionMockData) == 0 {
		return db.SessionInvoice{}, ErrorNotFound()
	}

	response := r.getUnsettledSessionInvoiceBySessionMockData[0]
	r.getUnsettledSessionInvoiceBySessionMockData = r.getUnsettledSessionInvoiceBySessionMockData[1:]
	return response.SessionInvoice, response.Error
}

func (r *MockRepositoryService) ListSessionInvoices(ctx context.Context, arg db.ListSessionInvoicesParams) ([]db.SessionInvoice, error) {
	if len(r.listSessionInvoicesMockData) == 0 {
		return []db.SessionInvoice{}, nil
	}

	response := r.listSessionInvoicesMockData[0]
	r.listSessionInvoicesMockData = r.listSessionInvoicesMockData[1:]
	return response.SessionInvoices, response.Error
}

func (r *MockRepositoryService) ListSessionInvoicesByNodeID(ctx context.Context, arg db.ListSessionInvoicesByNodeIDParams) ([]db.SessionInvoice, error) {
	if len(r.listSessionInvoicesByNodeIDMockData) == 0 {
		return []db.SessionInvoice{}, nil
	}

	response := r.listSessionInvoicesByNodeIDMockData[0]
	r.listSessionInvoicesByNodeIDMockData = r.listSessionInvoicesByNodeIDMockData[1:]
	return response.SessionInvoices, response.Error
}

func (r *MockRepositoryService) ListSessionInvoicesBySessionID(ctx context.Context, sessionID int64) ([]db.SessionInvoice, error) {
	if len(r.listSessionInvoicesBySessionIDMockData) == 0 {
		return []db.SessionInvoice{}, nil
	}

	response := r.listSessionInvoicesBySessionIDMockData[0]
	r.listSessionInvoicesBySessionIDMockData = r.listSessionInvoicesBySessionIDMockData[1:]
	return response.SessionInvoices, response.Error
}

func (r *MockRepositoryService) ListSessionInvoicesByUserID(ctx context.Context, arg db.ListSessionInvoicesByUserIDParams) ([]db.SessionInvoice, error) {
	if len(r.listSessionInvoicesByUserIDMockData) == 0 {
		return []db.SessionInvoice{}, nil
	}

	response := r.listSessionInvoicesByUserIDMockData[0]
	r.listSessionInvoicesByUserIDMockData = r.listSessionInvoicesByUserIDMockData[1:]
	return response.SessionInvoices, response.Error
}

func (r *MockRepositoryService) UpdateSessionInvoice(ctx context.Context, arg db.UpdateSessionInvoiceParams) (db.SessionInvoice, error) {
	r.updateSessionInvoiceMockData = append(r.updateSessionInvoiceMockData, arg)
	return db.SessionInvoice{}, nil
}

func (r *MockRepositoryService) GetCreateSessionInvoiceMockData() (db.CreateSessionInvoiceParams, error) {
	if len(r.createSessionInvoiceMockData) == 0 {
		return db.CreateSessionInvoiceParams{}, ErrorNotFound()
	}

	response := r.createSessionInvoiceMockData[0]
	r.createSessionInvoiceMockData = r.createSessionInvoiceMockData[1:]
	return response, nil
}

func (r *MockRepositoryService) GetUpdateSessionInvoiceMockData() (db.UpdateSessionInvoiceParams, error) {
	if len(r.updateSessionInvoiceMockData) == 0 {
		return db.UpdateSessionInvoiceParams{}, ErrorNotFound()
	}

	response := r.updateSessionInvoiceMockData[0]
	r.updateSessionInvoiceMockData = r.updateSessionInvoiceMockData[1:]
	return response, nil
}

func (r *MockRepositoryService) SetGetSessionInvoiceMockData(response SessionInvoiceMockData) {
	r.getSessionInvoiceMockData = append(r.getSessionInvoiceMockData, response)
}

func (r *MockRepositoryService) SetGetSessionInvoiceByPaymentRequestMockData(response SessionInvoiceMockData) {
	r.getSessionInvoiceByPaymentRequestMockData = append(r.getSessionInvoiceByPaymentRequestMockData, response)
}

func (r *MockRepositoryService) SetGetUnsettledSessionInvoiceBySessionMockData(response SessionInvoiceMockData) {
	r.getUnsettledSessionInvoiceBySessionMockData = append(r.getUnsettledSessionInvoiceBySessionMockData, response)
}

func (r *MockRepositoryService) SetListSessionInvoicesMockData(response SessionInvoicesMockData) {
	r.listSessionInvoicesMockData = append(r.listSessionInvoicesMockData, response)
}

func (r *MockRepositoryService) SetListSessionInvoicesByNodeIDMockData(response SessionInvoicesMockData) {
	r.listSessionInvoicesByNodeIDMockData = append(r.listSessionInvoicesByNodeIDMockData, response)
}

func (r *MockRepositoryService) SetListSessionInvoicesBySessionIDMockData(response SessionInvoicesMockData) {
	r.listSessionInvoicesBySessionIDMockData = append(r.listSessionInvoicesBySessionIDMockData, response)
}

func (r *MockRepositoryService) SetListSessionInvoicesByUserIDMockData(response SessionInvoicesMockData) {
	r.listSessionInvoicesByUserIDMockData = append(r.listSessionInvoicesByUserIDMockData, response)
}
