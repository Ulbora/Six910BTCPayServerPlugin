package six910btcplugin

import (
	//px "github.com/Ulbora/GoProxy"
	"crypto/ecdsa"
	"crypto/rand"
	"encoding/hex"

	lg "github.com/Ulbora/Level_Logger"

	cl "github.com/Ulbora/BTCPayClient"
	"github.com/btcsuite/btcd/btcec"
)

//Plugin Plugin
type Plugin interface {
	SetClient(c cl.Client)
	NewPairConnect(host string) *BTCPay
	CreateInvoice(inv *cl.InvoiceReq) *cl.InvoiceResponse
}

//PayPlugin PayPlugin
type PayPlugin struct {
	Client cl.Client
	log    *lg.Logger
}

//BTCPay BTCPay
type BTCPay struct {
	ClientID   string
	PublicKey  string
	PrivateKey string
	Token      string
	Host       string
}

//New New
func (p *PayPlugin) New() Plugin {
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	p.log = &l
	return p
}

//SetClient SetClient
func (p *PayPlugin) SetClient(c cl.Client) {
	p.Client = c
}

//NewClient NewClient
func (p *PayPlugin) NewClient(btc *BTCPay) Plugin {
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	p.log = &l

	var cryt cl.Cryptography
	cc := cryt.New()

	kp := cc.LoadKeyPair(btc.PrivateKey, btcec.S256())
	pub := cc.GetPublicKey(kp)
	p.log.Debug("public key: ", pub)

	var ptc cl.BTCPayClient

	var head cl.Headers
	ptc.SetHeader(head)

	p.getNewClient(btc.Host, kp, btc.Token)

	return p
}

//NewPairConnect NewPairConnect
func (p *PayPlugin) NewPairConnect(host string) *BTCPay {
	var rtn BTCPay
	pkh := p.secureRandom(32)
	var cryt cl.Cryptography
	cc := cryt.New()

	kp := cc.LoadKeyPair(pkh, btcec.S256())
	pub := cc.GetPublicKey(kp)
	p.log.Debug("private key: ", pkh)
	p.log.Debug("public key: ", pub)

	p.log.Debug("client: ", p.Client)

	p.getClient(host, kp)

	var tkr cl.TokenRequest
	tkr.ID = p.Client.GetClientID() // cc.GetSinFromKey(kp)
	tkr.Label = "Six910 access"
	tkr.Facade = "merchant"

	resp := p.Client.Token(&tkr)

	var pairingURL string
	if len(resp.Data) > 0 {
		pairingURL = p.Client.GetPairingCodeRequest(resp.Data[0].ParingCode)
	}
	p.log.Debug("ClientID: ", p.Client.GetClientID())
	p.log.Debug("token: ", resp.Data[0].Token)
	p.log.Debug("pairing url: ", pairingURL)
	rtn.ClientID = p.Client.GetClientID()
	rtn.Host = host
	rtn.PrivateKey = pkh
	rtn.PublicKey = pub
	rtn.Token = resp.Data[0].Token

	return &rtn

}

func (p *PayPlugin) getClient(host string, kp *ecdsa.PrivateKey) {
	if p.Client == nil {
		var ptc cl.BTCPayClient
		var head cl.Headers
		ptc.SetHeader(head)
		p.Client = ptc.New(host, kp, "")
	}

}
func (p *PayPlugin) getNewClient(host string, kp *ecdsa.PrivateKey, token string) {
	if p.Client == nil {
		var ptc cl.BTCPayClient
		var head cl.Headers
		ptc.SetHeader(head)
		p.Client = ptc.New(host, kp, token)
	}
}

//CreateInvoice CreateInvoice
func (p *PayPlugin) CreateInvoice(inv *cl.InvoiceReq) *cl.InvoiceResponse {
	return p.Client.CreateInvoice(inv)
}

//SecureRandom SecureRandom
func (p *PayPlugin) secureRandom(size int32) string {
	var rtn string
	bytes := make([]byte, size)
	_, err := rand.Read(bytes)
	if err == nil {
		rtn = hex.EncodeToString(bytes)
	}
	return rtn
}

// go mod init github.com/Ulbora/Six910BTCPayServerPlugin
