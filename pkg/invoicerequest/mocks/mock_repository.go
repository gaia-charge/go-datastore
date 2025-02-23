package mocks

import (
	mocks "github.com/gaia-charge/go-datastore/pkg/db/mocks"
	"github.com/gaia-charge/go-datastore/pkg/invoicerequest"
)

func NewRepository(repositoryService *mocks.MockRepositoryService) invoicerequest.InvoiceRequestRepository {
	return invoicerequest.InvoiceRequestRepository(repositoryService)
}
