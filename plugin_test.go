package six910btcplugin

import (
	"fmt"
	"testing"

	cl "github.com/Ulbora/BTCPayClient"
	"github.com/btcsuite/btcd/btcec"
)

const (
	testBaseURL = "https://testnet.demo.btcpayserver.org"
)

func TestPayPlugin_SecureRandom(t *testing.T) {
	var pi PayPlugin

	hn := pi.secureRandom(32)
	fmt.Println("hex num: ", hn)
	if hn == "" {
		t.Fail()
	}
	// t.Fail()
}

func TestPayPlugin_Connect(t *testing.T) {
	var ppi PayPlugin

	p := ppi.New()

	var mc MockBTCPayClient
	mc.MockClientID = "eeeddd"
	var tknr cl.TokenResponse
	var tkn cl.TokenData
	tkn.Token = "1123aaa"
	tkn.ParingCode = "pa111"
	tknr.Data = []cl.TokenData{tkn}
	mc.MockTokenResponse = &tknr
	mc.MockPairingCodeURL = "http://test.com/pair/123"
	p.SetClient(mc.New())

	btc := p.NewPairConnect(testBaseURL)
	fmt.Println("btc: ", *btc)
	if btc.PrivateKey == "" || p.GetToken() == "" {
		t.Fail()
	}

	// t.Fail()

}

func TestPayPlugin_NewClient(t *testing.T) {
	var pkh = "74f522c6704e39d102db6ae98dfb286e11c678a76fa32c93cc50244e436d936f"
	var pubk = "022709b094ef015db670278fb7bcbf90d35c5e3ed11094ce6ccd43f9458fc91c22"

	var btc BTCPay
	btc.PublicKey = pubk
	btc.PrivateKey = pkh
	btc.Host = testBaseURL
	btc.ClientID = "Tf4UFY3a3XyZX8PV4wj7zDaehQam4SkeTzq"
	btc.Token = "Dm3xKKMLbraE7BbQTxgdyLNavQENvCPW7Mgau3xzRXH4"

	var ppi PayPlugin
	c := ppi.NewClient(&btc)
	if c == nil || c.GetToken() == "" {
		t.Fail()
	}
	// t.Fail()

}

func TestPayPlugin_getClient(t *testing.T) {
	var ppi PayPlugin

	p := ppi.New()
	fmt.Println("p:", p)

	pkh := ppi.secureRandom(32)
	var cryt cl.Cryptography
	cc := cryt.New()

	kp := cc.LoadKeyPair(pkh, btcec.S256())
	ppi.getClient("aaa", kp)
	if ppi.Client == nil {
		t.Fail()
	}
}

func TestPayPlugin_getNewClient(t *testing.T) {
	var ppi PayPlugin

	p := ppi.New()
	fmt.Println("p:", p)

	pkh := ppi.secureRandom(32)
	var cryt cl.Cryptography
	cc := cryt.New()

	kp := cc.LoadKeyPair(pkh, btcec.S256())
	ppi.getNewClient("aaa", kp, "222")
	if ppi.Client == nil || p.GetToken() == "" {
		t.Fail()
	}
}

func TestPayPlugin_CreateInvoice(t *testing.T) {
	var pkh = "74f522c6704e39d102db6ae98dfb286e11c678a76fa32c93cc50244e436d936f"
	var pubk = "022709b094ef015db670278fb7bcbf90d35c5e3ed11094ce6ccd43f9458fc91c22"

	var btc BTCPay
	btc.PublicKey = pubk
	btc.PrivateKey = pkh
	btc.Host = testBaseURL
	btc.ClientID = "Tf4UFY3a3XyZX8PV4wj7zDaehQam4SkeTzq"
	btc.Token = "Dm3xKKMLbraE7BbQTxgdyLNavQENvCPW7Mgau3xzRXH4"

	var ppi PayPlugin

	var mc MockBTCPayClient
	mc.MockClientID = "eeeddd"
	var tknr cl.TokenResponse
	var tkn cl.TokenData
	tkn.Token = "1123aaa"
	tkn.ParingCode = "pa111"
	tknr.Data = []cl.TokenData{tkn}
	mc.MockTokenResponse = &tknr
	mc.MockPairingCodeURL = "http://test.com/pair/123"

	var invres cl.InvoiceResponse
	invres.Data.URL = "http://test"
	mc.MockInvoiceResponse = &invres

	ppi.SetClient(mc.New())

	c := ppi.NewClient(&btc)

	var req cl.InvoiceReq
	req.Price = 100.00
	req.Token = c.GetToken() //btc.Token
	req.Currency = "USD"
	req.TransactionSpeed = "medium"
	req.Buyer.Name = "bob willson"
	req.Buyer.Email = "bob@bob.com"
	res := c.CreateInvoice(&req)
	fmt.Println("inv res url: ", res.Data.URL)
	fmt.Println("inv req: ", req)
	if res.Data.URL != "http://test" {
		t.Fail()
	}

	// t.Fail()

}

func TestPayPlugin_SetLogLevel(t *testing.T) {
	var ppi PayPlugin
	var mc MockBTCPayClient
	ppi.SetClient(mc.New())
	p := ppi.New()
	p.SetLogLevel(3)
}

func TestPayPlugin_IsPluginLoaded(t *testing.T) {
	var ppi PayPlugin

	p := ppi.New()

	var mc MockBTCPayClient
	mc.MockClientID = "eeeddd"
	var tknr cl.TokenResponse
	var tkn cl.TokenData
	tkn.Token = "1123aaa"
	tkn.ParingCode = "pa111"
	tknr.Data = []cl.TokenData{tkn}
	mc.MockTokenResponse = &tknr
	mc.MockPairingCodeURL = "http://test.com/pair/123"
	p.SetClient(&mc)
	l := p.IsPluginLoaded()
	fmt.Println("loaded: ", l)
	if !l {
		t.Fail()
	}
	// t.Fail()
}
