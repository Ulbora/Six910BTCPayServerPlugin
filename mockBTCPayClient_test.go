package six910btcplugin

import (
	"testing"

	cl "github.com/Ulbora/BTCPayClient"
)

func TestMockBTCPayClient_GetRates(t *testing.T) {
	var mc MockBTCPayClient
	var rt cl.RateResponse
	mc.MockRateResponse = &rt

	var cp = []string{"USD"}

	res := mc.GetRates(cp, "111")

	if res == nil {
		t.Fail()
	}

}

func TestMockBTCPayClient_PairClient(t *testing.T) {
	var mc MockBTCPayClient
	var rt cl.TokenResponse
	mc.MockTokenResponse = &rt

	res := mc.PairClient("111")

	if res == nil {
		t.Fail()
	}

}

func TestMockBTCPayClient_GetInvoice(t *testing.T) {
	var mc MockBTCPayClient

	res := mc.GetInvoice("111")

	if res != nil {
		t.Fail()
	}

}

func TestMockBTCPayClient_GetInvoices(t *testing.T) {
	var mc MockBTCPayClient

	res := mc.GetInvoices(nil)

	if res != nil {
		t.Fail()
	}

}
