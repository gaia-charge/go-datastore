package param

import "github.com/gaia-charge/go-datastore/pkg/db"

func NewUpdateInvoiceRequestParams(invoiceRequest db.InvoiceRequest) db.UpdateInvoiceRequestParams {
	return db.UpdateInvoiceRequestParams{
		ID:             invoiceRequest.ID,
		PriceFiat:      invoiceRequest.PriceFiat,
		PriceMsat:      invoiceRequest.PriceMsat,
		CommissionFiat: invoiceRequest.CommissionFiat,
		CommissionMsat: invoiceRequest.CommissionMsat,
		TaxFiat:        invoiceRequest.TaxFiat,
		TaxMsat:        invoiceRequest.TaxMsat,
		TotalFiat:      invoiceRequest.TotalFiat,
		TotalMsat:      invoiceRequest.TotalMsat,
		IsSettled:      invoiceRequest.IsSettled,
		PaymentRequest: invoiceRequest.PaymentRequest,
	}
}
