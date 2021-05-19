package six910btcplugin

import (
	cl "github.com/Ulbora/BTCPayClient"
)

//MockBTCPayClient MockBTCPayClient
type MockBTCPayClient struct {
	MockClientID        string
	MockTokenResponse   *cl.TokenResponse
	MockPairingCodeURL  string
	MockRateResponse    *cl.RateResponse
	MockInvoiceResponse *cl.InvoiceResponse
}

//New New
func (a *MockBTCPayClient) New() cl.Client {
	return a
}

//GetClientID GetClientID
func (a *MockBTCPayClient) GetClientID() string {
	return a.MockClientID
}

//Token Token
func (a *MockBTCPayClient) Token(req *cl.TokenRequest) *cl.TokenResponse {
	return a.MockTokenResponse
}

//PairClient PairClient
func (a *MockBTCPayClient) PairClient(code string) *cl.TokenResponse {
	return a.MockTokenResponse
}

//GetPairingCodeRequest GetPairingCodeRequest
func (a *MockBTCPayClient) GetPairingCodeRequest(code string) string {
	return a.MockPairingCodeURL
}

//GetRates GetRates
func (a *MockBTCPayClient) GetRates(currencyPairs []string, storeID string) *cl.RateResponse {
	return a.MockRateResponse
}

//CreateInvoice CreateInvoice
func (a *MockBTCPayClient) CreateInvoice(inv *cl.InvoiceReq) *cl.InvoiceResponse {
	return a.MockInvoiceResponse
}

//GetInvoice GetInvoice
func (a *MockBTCPayClient) GetInvoice(invoiceID string) *cl.InvoiceResponse {
	return nil
}

//GetInvoices GetInvoices
func (a *MockBTCPayClient) GetInvoices(args *cl.InvoiceArgs) *cl.InvoiceListResponse {
	return nil
}
